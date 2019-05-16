package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExamplePing() {
	r := setRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	// Output: {"message":"pong"}
}
