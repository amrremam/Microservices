package main

import (
	"fmt"
	"net"
	"os"

	"github.com/amrremam/Microservices.Go/protos/server"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/building-microservices-youtube/currency/data"
	pb "github.com/nicholasjackson/building-microservices-youtube/currency/protos/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewCurrency(rates, log)

	// register the currency server
	pb.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", 5052))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(server)
}