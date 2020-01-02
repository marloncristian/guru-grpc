package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	pb "github.com/marloncristian/guru-grpc/server/api/v1/services"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

func main() {

	log.SetOutput(os.Stdout)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//socket configurations
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	log.Print("initialing server...")

	//initializes the grpc server
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCustomerServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

//server structure: business logic for handling income requests
type server struct {
}

//simple add method
func (*server) Add(ctx context.Context, request *pb.CustomerAddRequest) (*pb.CustomerAddResponse, error) {
	response := &pb.CustomerAddResponse{
		CustomerId: 1710000,
	}
	return response, nil
}
