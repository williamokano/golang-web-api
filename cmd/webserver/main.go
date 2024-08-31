package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/williamokano/golang-web-api/internal/database"
	"github.com/williamokano/golang-web-api/internal/web"
)

func main() {
	// Start database
	db, err := database.CreateDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	err = database.MigrateDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	server := web.NewServer(db).SetupServer()
	go func() {
		// start listening to connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shut down the server
	// with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)

	// add listener for sigint and sigterm as sigkill cannot be caught
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	timeout := 5 * time.Second
	if val, ok := os.LookupEnv("SERVER_TIMEOUT"); ok {
		if newTimeout, err := strconv.Atoi(val); err != nil {
			timeout = time.Duration(newTimeout) * time.Second
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func() {
		db.Close()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
	}()

	// catch ctx.Done() with 5 seconds timeout
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("timeout of %.0f seconds.", timeout.Seconds()))
	}
	log.Println("Server exiting")
}
