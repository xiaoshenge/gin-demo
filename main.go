package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// logFile := fmt.Sprintf("logs/%s", time.Now().Format("2006-01-02"))
	// f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ip", func(c *gin.Context) {
		var localIP string
		host, _ := os.Hostname()
		addrs, _ := net.LookupIP(host)
		for _, addr := range addrs {
			if ipv4 := addr.To4(); ipv4 != nil {
				if localIP == "" {
					localIP = fmt.Sprintf("%s", ipv4)
				} else {
					localIP = fmt.Sprintf("%s,%s", localIP, ipv4)
				}
			}
		}

		c.JSON(200, gin.H{
			"srvIp": localIP,
		})
	})

	server := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()


	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2, syscall.SIGKILL)

	<-quit

	log.Println("receive interrupt signal")

	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cannel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server Shutdown:", err)
	}

	log.Println("server exiting")
}
