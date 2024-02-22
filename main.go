package main

import (
	"context"
	"log"
	"net"

	"github.com/erabxes/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

// create a type that implements the InvoiceServer interface
type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	// create an instance of the grpc server
	serverRegistrar := grpc.NewServer()
	// create an instance of the type that implements the InvoiceServer
	service := &myInvoiceServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	// Now to launch the server
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to launch server: %v", err)
	}

}
