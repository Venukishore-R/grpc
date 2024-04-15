package main

import (
	"context"
	"github.com/Venukishore-R/grpc/v2/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoiceCreateServer
}

func (s myInvoiceServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf: []byte("test1"),
		Doc: []byte("test2"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "addr:8089")
	if err != nil {
		log.Fatalf("cannot able to create listener: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}

	invoicer.RegisterInvoiceCreateServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve %v", err)
	}
}
