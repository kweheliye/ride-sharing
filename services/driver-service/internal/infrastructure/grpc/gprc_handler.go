package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	service2 "ride-sharing/services/driver-service/internal/service"
	pb "ride-sharing/shared/proto/driver"
)

type driverGrpcHandler struct {
	pb.UnimplementedDriverServiceServer

	service *service2.Service
}

func NewGrpcHandler(s *grpc.Server, service *service2.Service) {
	handler := &driverGrpcHandler{
		service: service,
	}

	pb.RegisterDriverServiceServer(s, handler)
}

func (h *driverGrpcHandler) RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDriver not implemented")
}

func (h *driverGrpcHandler) UnregisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterDriver not implemented")
}
