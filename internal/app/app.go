package app

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/lucasd-coder/user-manger-service/config"
	"github.com/lucasd-coder/user-manger-service/internal/domain/user/service"
	"github.com/lucasd-coder/user-manger-service/pkg/logger"
	"github.com/lucasd-coder/user-manger-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	logger := logger.NewLog(cfg)

	log := logger.GetGRPCLogger()

	lis, err := net.Listen("tcp", "localhost:"+cfg.Port)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			logger.GetGRPCUnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			logger.GetGRPCStreamServerInterceptor(),
		),
	)
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	pb.RegisterUserServiceServer(grpcServer, &service.UserService{})

	reflection.Register(grpcServer)

	log.Infof("Started listening... address[:%s]", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
