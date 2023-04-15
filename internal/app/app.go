package app

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"os/signal"
	"syscall"

	"ShamanEstBanan-GophKeeper-server/internal/server/authServer"

	"ShamanEstBanan-GophKeeper-server/internal/config"
	"ShamanEstBanan-GophKeeper-server/internal/db/migrate"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	"ShamanEstBanan-GophKeeper-server/internal/server"
	"ShamanEstBanan-GophKeeper-server/internal/service"
	"ShamanEstBanan-GophKeeper-server/internal/storage"
	"ShamanEstBanan-GophKeeper-server/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type App struct {
	logger *zap.Logger
	server *grpc.Server
}

const (
	migrationsPath = "internal/db/migrate/migrations/"
)

func New() (*App, error) {
	cfg := config.New()
	ctx := context.Background()
	l, err := logger.New(cfg.Debug)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Init migrations
	if err := migrate.Run(cfg.PostgresDSN, migrate.WithPath(migrationsPath)); err != nil {
		log.Fatalf("failed executing migrate DB: %v\n", err) //nolint: revive
	}

	// init storage
	pool, err := pgxpool.New(ctx, cfg.PostgresDSN)
	if err != nil {
		l.Error(err.Error())
		return nil, err
	}
	st := storage.New(pool, l)
	newService := service.New(l, st)
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var s *grpc.Server
	if cfg.Debug == true {
		s = grpc.NewServer(grpc.UnaryInterceptor(authServer.AuthInterceptor))
	} else {
		s = grpc.NewServer(
			grpc.UnaryInterceptor(authServer.AuthInterceptor),
			grpc.Creds(tlsCredentials),
		)
	}

	// регистрируем сервис
	pb.RegisterKeeperServiceServer(s, &server.KeeperService{
		Service:                          newService,
		UnimplementedKeeperServiceServer: pb.UnimplementedKeeperServiceServer{},
	})
	pb.RegisterAuthServiceServer(s, &authServer.AuthServer{
		Service:                        newService,
		UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
	})
	a := &App{
		logger: l,
		server: s,
	}
	return a, nil
}

func (a App) Run() error {
	_, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	listener, err := net.Listen("tcp", ":3080")
	if err = a.server.Serve(listener); err != nil {
		log.Fatal(err)
	}
	defer a.server.GracefulStop()
	return err
}

const (
	serverCert = "internal/app/certs/server-cert.pem"
	serverKey  = "internal/app/certs/server-key.pem"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, err
	}
	// Create the credentials and return it
	cfg := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(cfg), nil
}
