package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arunmalligere/to-do-service/app/to-do/handlers"
)

func main()  {
	log := log.New(os.Stdout, "TO-DO::", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error  {
	// =========================================================================
	// Configuration
	type Web struct {
			AppHost string
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
		}
	cfg := struct {
		test string
		Web Web
	} {
		test: "test",
		Web: Web {
			AppHost: "0.0.0.0:3000",
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}


	// =========================================================================
	// Startup logs 
	log.Println("main: Started : Initializing ...")
	defer log.Println("main: Completed")

	// =========================================================================
	// Start the web service

	log.Println("Service started")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server {
		Addr: cfg.Web.AppHost,
		Handler: handlers.API(shutdown, log),
		ReadTimeout: cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- api.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown
	select {
	case err := <- serverErrors:
		log.Println("Fatal error")
		return err
	case sig := <-shutdown:
		log.Println("Shutdown", sig)
		return errors.New("Shutdown")
	}

	return nil
}