package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/photoshelf/photoshelf-storage/infrastructure/container"
	"github.com/photoshelf/photoshelf-storage/presentation/controller"
	"github.com/photoshelf/photoshelf-storage/presentation/protobuf"
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	s := grpc.NewServer()

	ctrl := controller.NewPhotoController()
	container.Get(ctrl)

	protobuf.RegisterPhotoServiceServer(s, ctrl)
	return s
}

func NewGateway(endpoint string) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	if err := protobuf.RegisterPhotoServiceHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{}); err != nil {
		return nil, err
	}

	return mux, nil
}
