package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/polshe-v/microservices_chat_server/internal/api/chat"
	"github.com/polshe-v/microservices_chat_server/internal/closer"
	config "github.com/polshe-v/microservices_chat_server/internal/config"
	"github.com/polshe-v/microservices_chat_server/internal/config/env"
	"github.com/polshe-v/microservices_chat_server/internal/repository"
	chatRepository "github.com/polshe-v/microservices_chat_server/internal/repository/chat"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	chatService "github.com/polshe-v/microservices_chat_server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GrpcConfig

	pgPool *pgxpool.Pool

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

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		// Create database connections pool.
		pool, err := pgxpool.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.PgPool(ctx))
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
