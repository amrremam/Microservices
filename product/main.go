package main

import (
	"context"
	"github.com/amrremam/Microservices.Go/product/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)


func main() {
	l := log.New(os.Stdout, "product", log.LstdFlags)
	hi := handlers.NewHello(l)
	ph := handlers.NewProducts(l)
	gb := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/hi", hi)
	sm.Handle("/bye", gb)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
