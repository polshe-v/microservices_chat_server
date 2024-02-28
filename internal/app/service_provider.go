package app

import (
	"context"
	"log"

	"github.com/polshe-v/microservices_chat_server/internal/api/chat"
	"github.com/polshe-v/microservices_chat_server/internal/client/db"
	"github.com/polshe-v/microservices_chat_server/internal/client/db/pg"
	"github.com/polshe-v/microservices_chat_server/internal/closer"
	"github.com/polshe-v/microservices_chat_server/internal/config"
	"github.com/polshe-v/microservices_chat_server/internal/config/env"
	"github.com/polshe-v/microservices_chat_server/internal/repository"
	chatRepository "github.com/polshe-v/microservices_chat_server/internal/repository/chat"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	chatService "github.com/polshe-v/microservices_chat_server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GrpcConfig

	dbClient db.Client

	chatRepository repository.ChatRepository
	chatService    service.ChatService
	chatImpl       *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GrpcConfig() config.GrpcConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		c, err := pg.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = c.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		closer.Add(c.Close)

		s.dbClient = c
	}

	return s.dbClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}
	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx))
	}
	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}
	return s.chatImpl
}
