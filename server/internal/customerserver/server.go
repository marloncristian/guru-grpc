package customerserver

import (
	"context"
	pb "github.com/marloncristian/guru-grpc/server/rpc/customer"
	"math/rand"
)

// Server customer server
type Server struct {
}

//Add simple add method
func (*Server) Add(ctx context.Context, request *pb.CustomerAddRequest) (*pb.CustomerAddResponse, error) {
	//business layer
	response := &pb.CustomerAddResponse{
		CustomerId: rand.Int63n(100),
	}
	return response, nil
}
