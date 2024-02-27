package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	config "github.com/polshe-v/microservices_chat_server/internal/config"
	env "github.com/polshe-v/microservices_chat_server/internal/config/env"
	"github.com/polshe-v/microservices_chat_server/internal/converter"
	chatRepository "github.com/polshe-v/microservices_chat_server/internal/repository/chat"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	chatService "github.com/polshe-v/microservices_chat_server/internal/service/chat"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "Path to config file")
}

const delim = "---"

type server struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chat := req.GetChat()
	log.Printf("\n%s\nUsernames: %v\n%s", delim, strings.Join(chat.GetUsernames(), ", "), delim)

	id, err := s.chatService.Create(ctx, converter.ToChatFromDesc(chat))
	if err != nil {
		return nil, err
	}

	log.Printf("Created chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("\n%s\nID: %d\n%s", delim, req.GetId(), delim)

	err := s.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	message := req.GetMessage()
	log.Printf("\n%s\nFrom: %s\nText: %s\nTimestamp: %v\n%s", delim, message.GetFrom(), message.GetText(), message.GetTimestamp(), delim)

	return &empty.Empty{}, nil
}

func main() {
	// Parse the command-line flags from os.Args[1:].
	flag.Parse()
	ctx := context.Background()

	// Read config file.
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGrpcConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	// Open IP and port for server.
	lis, err := net.Listen(grpcConfig.Transport(), grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pgConfig, err := env.NewPgConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Create repository layer.
	chatRepo := chatRepository.NewRepository(pool)
	chatSrv := chatService.NewService(chatRepo)

	// Create gRPC *Server which has no service registered and has not started to accept requests yet.
	s := grpc.NewServer()

	// Upon the client's request, the server will automatically provide information on the supported methods.
	reflection.Register(s)

	// Register service with corresponded interface.
	desc.RegisterChatV1Server(s, &server{chatService: chatSrv})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
