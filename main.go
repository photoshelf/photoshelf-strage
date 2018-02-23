package main

import (
	"fmt"
	"github.com/photoshelf/photoshelf-storage/application"
	"github.com/photoshelf/photoshelf-storage/infrastructure/server"
	"log"
	"net/http"
	"os"
)

func main() {
	conf, err := application.Configure(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	httpAddress := fmt.Sprintf(":%d", conf.Server.Port)
	grpcAddress := fmt.Sprintf(":%d", conf.Gateway.Port)

	gw, err := server.NewGateway(grpcAddress)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	go func() {
		if err := http.ListenAndServe(httpAddress, gw); err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
	}()

	s := server.NewServer()
	defer s.GracefulStop()

	if err := http.ListenAndServe(httpAddress, s); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
