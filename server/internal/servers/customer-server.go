package servers

import (
	"context"
	pb "github.com/marloncristian/guru-grpc/server/rpc/customer"
)

// CustomerServer customer server
type CustomerServer struct {
}

//Add simple add method
func (*CustomerServer) Add(ctx context.Context, request *pb.CustomerAddRequest) (*pb.CustomerAddResponse, error) {
	response := &pb.CustomerAddResponse{
		CustomerId: 1710000,
	}
	return response, nil
}
