FROM golang:1.11.4-alpine3.8 as build
RUN apk add --update --no-cache ca-certificates git
RUN mkdir -p /go/src/github.com/xiaoshenge/gin-demo
WORKDIR /go/src/github.com/xiaoshenge/gin-demo
COPY go.mod .
COPY go.sum .
ENV GO111MODULE=on 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /go/bin/gin-demo
FROM scratch
COPY --from=build /go/bin/gin-demo /go/bin/gin-demo
EXPOSE 9090
ENTRYPOINT [ "/go/bin/gin-demo" ]