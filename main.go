package main
import (
	"fmt"
	"net"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ip", func (c *gin.Context)  {
		var localIP string
		host, _ := os.Hostname()
		addrs, _ := net.LookupIP(host)
		for _, addr := range addrs {
			if ipv4 := addr.To4(); ipv4 != nil {
				if localIP == "" {
					localIP = fmt.Sprintf("%s", ipv4)
				} else{
					localIP = fmt.Sprintf("%s,%s",localIP,ipv4)
				}
			}   
		}

		c.JSON(200, gin.H{
			"srvIp": localIP,
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}