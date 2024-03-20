package app

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	descAccess "github.com/polshe-v/microservices_auth/internal/pkg/access_v1"
	"github.com/polshe-v/microservices_chat_server/internal/api/chat"
	rpcAuth "github.com/polshe-v/microservices_chat_server/internal/client/rpc/auth"
	"github.com/polshe-v/microservices_chat_server/internal/config"
	"github.com/polshe-v/microservices_chat_server/internal/config/env"
	"github.com/polshe-v/microservices_chat_server/internal/interceptor"
	"github.com/polshe-v/microservices_chat_server/internal/repository"
	chatRepository "github.com/polshe-v/microservices_chat_server/internal/repository/chat"
	logRepository "github.com/polshe-v/microservices_chat_server/internal/repository/log"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	chatService "github.com/polshe-v/microservices_chat_server/internal/service/chat"
	"github.com/polshe-v/microservices_common/pkg/closer"
	"github.com/polshe-v/microservices_common/pkg/db"
	"github.com/polshe-v/microservices_common/pkg/db/pg"
	"github.com/polshe-v/microservices_common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GrpcConfig
	authConfig config.AuthConfig

	authClient        *rpcAuth.Client
	dbClient          db.Client
	txManager         db.TxManager
	interceptorClient *interceptor.Client

	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository
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

func (s *serviceProvider) AuthConfig() config.AuthConfig {
	if s.authConfig == nil {
		cfg, err := env.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get authentication service config: %v", err)
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) AuthClient() *rpcAuth.Client {
	if s.authClient == nil {
		cfg := s.AuthConfig()
		creds, err := credentials.NewClientTLSFromFile(cfg.CertPath(), "")
		if err != nil {
			log.Fatalf("failed to get credentials of authentication service: %v", err)
		}

		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(creds),
		)
		if err != nil {
			log.Fatalf("failed to connect to authentication service: %v", err)
		}

		s.authClient = rpcAuth.NewAuthClient(descAccess.NewAccessV1Client(conn))
	}

	return s.authClient
}

func (s *serviceProvider) InterceptorClient() *interceptor.Client {
	if s.interceptorClient == nil {
		s.interceptorClient = &interceptor.Client{
			Client: s.AuthClient(),
		}
	}
	return s.interceptorClient
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

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}
	return s.chatRepository
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DBClient(ctx))
	}
	return s.logRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx), s.LogRepository(ctx), s.TxManager(ctx))
	}
	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}
	return s.chatImpl
}
