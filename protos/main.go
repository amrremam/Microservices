package main

import (
	"net"
	"net/http"
	"os"

	protos "github.com/amrremam/Microservices.Go/protos/currency"
	"github.com/amrremam/Microservices.Go/protos/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)


func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	l, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Error("Unable to serve", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}

