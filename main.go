package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		// start listening to connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server
	// with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)

	// add listener for sigint and sigterm as sigkill cannot be catch
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")
	timeout := 10 * time.Second
	if val, ok := os.LookupEnv("SERVER_TIMEOUT"); ok {
		if newTimeout, err := strconv.Atoi(val); err != nil {
			timeout = time.Duration(newTimeout) * time.Second
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catch ctx.Done() with 5 seconds timeout
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("timeout of %.0f seconds.", timeout.Seconds()))
	}
	log.Println("Server exiting")
}
