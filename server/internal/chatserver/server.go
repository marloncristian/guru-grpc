package chatserver

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/marloncristian/guru-grpc/server/rpc/chat"
)

// Server customer server
type Server struct {
	message chan string
}

// Send sends message to the server
func (s *Server) Send(ctx context.Context, message *wrappers.StringValue) (*empty.Empty, error) {
	if message != nil {
		log.Printf("Send requested: message=%v", message.Value)
		s.message <- message.Value
	} else {
		log.Print("Send requested: message=<empty>")
	}

	return &empty.Empty{}, nil
}

// Subscribe is streaming method to get echo messages from the server
func (s *Server) Subscribe(e *empty.Empty, stream pb.ChatService_SubscribeServer) error {
	log.Print("Subscribe requested")
	for {
		m := <-s.message
		n := wrappers.StringValue{Value: fmt.Sprintf("I have received from you: %s. Thanks!", m)}
		if err := stream.Send(&n); err != nil {
			s.message <- m
			log.Printf("Stream connection failed: %v", err)
			return nil
		}
		log.Printf("Message sent: %+v", n.Value)
	}
}

// NewChatServer novo chat server
func NewChatServer() *Server {
	return &Server{
		message: make(chan string, 1000),
	}
}
