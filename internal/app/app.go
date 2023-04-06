package app

import (
	"ShamanEstBanan-GophKeeper-server/internal/server"
	"context"
	"crypto/tls"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"ShamanEstBanan-GophKeeper-server/internal/config"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	"ShamanEstBanan-GophKeeper-server/internal/service"
	"ShamanEstBanan-GophKeeper-server/internal/storage"
	"ShamanEstBanan-GophKeeper-server/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type App struct {
	logger *zap.Logger
	server *grpc.Server
}

func New() (*App, error) {
	cfg := config.New()
	ctx := context.Background()
	l, err := logger.New(cfg.Debug)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// init storage
	pool, err := pgxpool.New(ctx, cfg.PostgresDSN)
	if err != nil {
		l.Error(err.Error())
		return nil, err
	}
	st := storage.New(pool, l)
	service := service.New(l, st)
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var s *grpc.Server
	if cfg.Debug == true {
		s = grpc.NewServer()
	} else {
		s = grpc.NewServer(grpc.Creds(tlsCredentials))
	}

	// регистрируем сервис
	pb.RegisterKeeperServiceServer(s, &server.KeeperService{
		Service:                          service,
		UnimplementedKeeperServiceServer: pb.UnimplementedKeeperServiceServer{},
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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(
		"internal/app/certs/server-cert.pem",
		"internal/app/certs/server-key.pem",
	)
	if err != nil {
		return nil, err
	}
	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

// listen OS signals to gracefully shutdown server
func listen(ctx context.Context, server *grpc.Server, listener net.Listener) error {
	srv := server

	eg, ctx := errgroup.WithContext(ctx)

	log.Println("server started 1", zap.String("addr:", ""))
	eg.Go(func() error {
		<-ctx.Done()
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.GracefulStop()

		return nil
	})

	return eg.Wait()
}
