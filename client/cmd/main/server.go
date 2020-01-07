package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/marloncristian/guru-grpc/client/rpc/chat"
	"github.com/marloncristian/guru-grpc/client/rpc/customer"
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
	customerclient := customer.NewCustomerServiceClient(conn)
	chatclient := chat.NewChatServiceClient(conn)

	//calls the add method and handles response
	request := &customer.CustomerAddRequest{
		Id:   21,
		Name: "Marlon Cristian Pereira",
	}
	resp, err := customerclient.Add(context.Background(), request)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	} else {
		log.Print(resp.CustomerId)
	}

	//sends a chat message throught chat channel
	_, err = chatclient.Send(context.Background(), &wrappers.StringValue{Value: "ping"})
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	//listens for any server message
	client, err := chatclient.Subscribe(context.Background(), &empty.Empty{})
	lchan := make(chan string)
	go func() {
		r, err := client.Recv()
		if err != nil {
			log.Printf("err")
		} else {
			log.Printf(fmt.Sprintf("Response : %v", r.Value))

			//closes the waiting channel as soon as it receives a response
			close(lchan)
		}
	}()
	<-lchan
}
