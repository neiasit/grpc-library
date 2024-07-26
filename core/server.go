package core

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"time"
)

var grpcServerTag = slog.String("server", "grpc_server")

func NewGrpcServer(logger *slog.Logger, unaryInterceptors []grpc.UnaryServerInterceptor) *grpc.Server {
	logger.Info("Initializing gRPC server")
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(server)
	return server
}

func RunGrpcServer(lc fx.Lifecycle, srv *grpc.Server, logger *slog.Logger, cfg *Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting server", grpcServerTag, "address", cfg.Address())
			listener, err := net.Listen("tcp", cfg.Address())
			if err != nil {
				logger.Error("cannot start server", "error", err.Error(), grpcServerTag)
				return err
			}
			go func() {
				err := srv.Serve(listener)
				if err != nil {
					logger.Error("cannot start server", "error", err.Error(), grpcServerTag)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutting down", grpcServerTag)
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			srv.GracefulStop()
			return nil
		},
	})
}
