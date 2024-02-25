package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	config "github.com/polshe-v/microservices_chat_server/internal/config"
	env "github.com/polshe-v/microservices_chat_server/internal/config/env"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

var configPath string
var errQueryBuild = errors.New("failed to build query")

func init() {
	flag.StringVar(&configPath, "config", ".env", "Path to config file")
}

const delim = "---"

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("\n%s\nUsernames: %v\n%s", delim, strings.Join(req.GetUsernames(), ", "), delim)

	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("usernames").
		Values(req.GetUsernames()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("%v", err)
		return nil, errQueryBuild
	}

	var id int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Printf("%v", err)
		return nil, errors.New("failed to create chat")
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("\n%s\nID: %d\n%s", delim, req.GetId(), delim)

	builderDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("%v", err)
		return nil, errQueryBuild
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("%v", err)
		return nil, errors.New("failed to delete chat")
	}
	log.Printf("result: %v", res)

	return &empty.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("\n%s\nFrom: %s\nText: %s\nTimestamp: %v\n%s", delim, req.GetFrom(), req.GetText(), req.GetTimestamp(), delim)

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

	// Create gRPC *Server which has no service registered and has not started to accept requests yet.
	s := grpc.NewServer()

	// Upon the client's request, the server will automatically provide information on the supported methods.
	reflection.Register(s)

	// Register service with corresponded interface.
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
