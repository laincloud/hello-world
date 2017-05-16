package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// LogFlag 控制日志的前缀
const LogFlag = log.LstdFlags | log.Lmicroseconds | log.Lshortfile

var (
	errLogger  *log.Logger
	infoLogger *log.Logger
)

func init() {
	errLogger = log.New(os.Stderr, "ERROR ", LogFlag)
	infoLogger = log.New(os.Stdout, "INFO ", LogFlag)
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)

	// Register API endpoint
	http.HandleFunc("/hello", hello)

	server := http.Server{
		Addr: ":8080",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errLogger.Printf("server.ListenAndServe() failed, error: %s.", err)
		}
	}()
	infoLogger.Printf("server.ListenAndServe()..., Addr: %s.", ":8080")

	<-quit
	infoLogger.Printf("Shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		errLogger.Printf("server.Shutdown() failed, error: %s.", err)
	}

	infoLogger.Printf("Shut down.")
}

func hello(w http.ResponseWriter, r *http.Request) {
	infoLogger.Printf("Receive a hello request, URL: %s, Method: %s, RemoteAddr: %s.", r.URL.String(), r.Method, r.RemoteAddr)

	if _, err := fmt.Fprint(w, "Hello, world."); err != nil {
		errLogger.Printf("fmt.Fprint() failed, error: %s.", err)
	}

	infoLogger.Printf("Response a hello request, URL: %s, Method: %s, RemoteAddr: %s.", r.URL.String(), r.Method, r.RemoteAddr)
}
