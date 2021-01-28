package main

import (
	pb "currency/protos"
	"currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	c := server.NewCurrency(log)

	pb.RegisterCurrencyServer(gs, c)

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Listen error")
		os.Exit(1)
	}

	reflection.Register(gs)

	log.Info("Listening 9092")
	gs.Serve(lis)
}
