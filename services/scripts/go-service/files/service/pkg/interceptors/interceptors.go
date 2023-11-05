package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Get the HTTP headers from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing context metadata")
	}

	// Create a new context with the Authorization header as metadata
	newCtx := metadata.NewOutgoingContext(ctx, md)

	// Call the gRPC handler with the new context
	return handler(newCtx, req)
}