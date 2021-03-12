package main

import (
	"log"
	"net"
	"os"

	"github.com/shatskiym/go_services/protos/currency"
	"github.com/shatskiym/go_services/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := log.New(os.Stdout, "micro-api", log.LstdFlags)

	gs := grpc.NewServer()

	ch := server.NewCurrency(log)

	currency.RegisterCurrencyServer(gs, ch)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	gs.Serve(l)

}
