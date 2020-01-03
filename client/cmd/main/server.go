package main

import (
	"context"
	"flag"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	pb "github.com/marloncristian/guru-grpc/client/rpc/customer"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {

	log.SetOutput(os.Stdout)

	//grpc dial options
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	//dial to server
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	//interface for server request interop
	client := pb.NewCustomerServiceClient(conn)
	request := &pb.CustomerAddRequest{
		Id:   21,
		Name: "Marlon Cristian Pereira",
	}

	//calls the add method and handles response
	resp, err := client.Add(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	} else {
		log.Print(resp.CustomerId)
	}
}
