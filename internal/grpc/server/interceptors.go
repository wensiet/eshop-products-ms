package productsAPI

import (
	"context"
	"eshop-products-ms/internal/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func CorsLikeInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}
	logger := config.GetLogger()
	logger.Debug(fmt.Sprintf("metadata: %v", md))

	allowedHost := "http://gateway:8000"

	hostValues := md.Get("host")

	logger.Debug(fmt.Sprintf("host: %s", hostValues))

	if len(hostValues) == 0 || hostValues[0] != allowedHost {
		err := status.Error(codes.PermissionDenied, "invalid host")
		return nil, err
	}

	return handler(ctx, req)
}
