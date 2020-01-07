package customerserver

import (
	"context"
	"log"
	"math/rand"

	pb "github.com/marloncristian/guru-grpc/server/rpc/customer"
)

// Server customer server
type Server struct {
}

//Add simple add method
func (*Server) Add(ctx context.Context, request *pb.CustomerAddRequest) (*pb.CustomerAddResponse, error) {

	log.Printf("Received : %v", request.Name)
	//business layer
	response := &pb.CustomerAddResponse{
		CustomerId: rand.Int63n(100),
	}
	return response, nil
}

// NewCustomerServer novo servidor
func NewCustomerServer() *Server {
	return &Server{}
}
