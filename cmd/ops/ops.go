package main

import (
	"flag"
	"log"

	"github.com/bersennaidoo/ichatops/application/grpc/petstore/client"
	jaegerClient "github.com/bersennaidoo/ichatops/application/ops/jaeger/client"
	"github.com/bersennaidoo/ichatops/application/ops/prom"
	"github.com/bersennaidoo/ichatops/application/ops/server"
)

var (
	addr         = flag.String("addr", "0.0.0.0:7000", "The address to run the service on.")
	jaegerAddr   = flag.String("jaegerAddr", "127.0.0.1:16685", "The address of the Jaeger query service.")
	promAddr     = flag.String("promAddr", "127.0.0.1:9000", "The address of the Prometheus service.")
	petstoreAddr = flag.String("petstoreAddr", "127.0.0.1:6742", "The address of the Petstore.")
)

func main() {
	flag.Parse()

	j, err := jaegerClient.New(*jaegerAddr)
	if err != nil {
		log.Fatalf("could not connect to Jaeger: %s", err)
	}

	p, err := prom.New("http://" + *promAddr)
	if err != nil {
		log.Fatalf("could not connect to Prometheus: %s", err)
	}

	ps, err := client.New(*petstoreAddr)
	if err != nil {
		log.Fatalf("could not connecto the Petstore: %s", err)
	}
	clients := server.Clients{
		Jaeger:   j,
		Prom:     p,
		Petstore: ps,
	}
	serv, err := server.New(*addr, clients)
	if err != nil {
		panic(err)
	}
	log.Println("serving starting on: ", *addr)
	serv.Start()
}
