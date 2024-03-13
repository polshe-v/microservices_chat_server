package app

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/polshe-v/microservices_chat_server/internal/config"
	"github.com/polshe-v/microservices_chat_server/internal/interceptor"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
	"github.com/polshe-v/microservices_common/pkg/closer"
)

// App structure contains main application structures.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "Path to config file")
}

// NewApp creates new App object.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Run executes the application.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGrpcServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	// Parse the command-line flags from os.Args[1:].
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	cfg := a.serviceProvider.GrpcConfig()
	creds, err := credentials.NewServerTLSFromFile(cfg.CertPath(), cfg.KeyPath())
	if err != nil {
		return err
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptor.PolicyInterceptor),
	)

	interceptor.AuthService = &interceptor.AuthServiceParams{
		AuthAddress:  cfg.AuthAddress(),
		AuthCertPath: cfg.AuthCertPath(),
	}

	// Upon the client's request, the server will automatically provide information on the supported methods.
	reflection.Register(a.grpcServer)

	// Register service with corresponded interface.
	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))

	return nil
}

func (a *App) runGrpcServer() error {
	// Open IP and port for server.
	lis, err := net.Listen(a.serviceProvider.GrpcConfig().Transport(), a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		return err
	}

	log.Printf("gRPC server running on %v", a.serviceProvider.GrpcConfig().Address())

	err = a.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
