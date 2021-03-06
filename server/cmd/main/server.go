package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	"github.com/marloncristian/guru-grpc/server/internal/customerserver"
	"github.com/marloncristian/guru-grpc/server/rpc/customer"

	"github.com/marloncristian/guru-grpc/server/internal/chatserver"
	"github.com/marloncristian/guru-grpc/server/rpc/chat"
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

	//register customer server
	customer.RegisterCustomerServiceServer(grpcServer, customerserver.NewCustomerServer())

	//register chat server
	chat.RegisterChatServiceServer(grpcServer, chatserver.NewChatServer())

	grpcServer.Serve(lis)
}
