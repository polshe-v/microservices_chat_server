package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"net"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

const (
	grpcTransport = "tcp"
	grpcIP        = "127.0.0.1"
	grpcPort      = 50001
	delim         = "---"
)

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("\n%s\nUsernames: %v\n%s", delim, strings.Join(req.GetUsernames(), ", "), delim)

	// Generate random ID.
	id, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		log.Fatalf("failed to generate ID: %v", err)
	}

	return &desc.CreateResponse{
		Id: id.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("\n%s\nID: %d\n%s", delim, req.GetId(), delim)

	return &empty.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("\n%s\nFrom: %s\nText: %s\nTimestamp: %v\n%s", delim, req.GetFrom(), req.GetText(), req.GetTimestamp(), delim)

	return &empty.Empty{}, nil
}

func main() {
	// Open IP and port for server.
	lis, err := net.Listen(grpcTransport, fmt.Sprintf("%s:%d", grpcIP, grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create gRPC *Server which has no service registered and has not started to accept requests yet.
	s := grpc.NewServer()

	// Upon the client's request, the server will automatically provide information on the supported methods.
	reflection.Register(s)

	// Register service with corresponded interface.
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
