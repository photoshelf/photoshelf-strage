package main

import (
	"fmt"
	"github.com/photoshelf/photoshelf-storage/application"
	"github.com/photoshelf/photoshelf-storage/infrastructure/server"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	conf, err := application.Configure(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	grpcAddress := fmt.Sprintf(":%d", conf.Server.Port)
	httpAddress := fmt.Sprintf(":%d", conf.Gateway.Port)

	gw, err := server.NewGateway(grpcAddress)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	go func() {
		log.Printf("Gateway Server start with port%s\n", httpAddress)
		if err := http.ListenAndServe(httpAddress, gw); err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
	}()

	s := server.NewServer()
	defer s.GracefulStop()

	log.Printf("gRPC Server start with port%s\n", grpcAddress)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
