package customerserver

import (
	"context"
	pb "github.com/marloncristian/guru-grpc/server/rpc/customer"
)

// Server customer server
type Server struct {
}

//Add simple add method
func (*Server) Add(ctx context.Context, request *pb.CustomerAddRequest) (*pb.CustomerAddResponse, error) {
	//business layer
	response := &pb.CustomerAddResponse{
		CustomerId: 1710000,
	}
	return response, nil
}
