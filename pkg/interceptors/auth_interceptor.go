package interceptors

import (
	"context"
	authLib "github.com/neiasit/auth-library/pkg/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
	"strings"
)

func NewAuthInterceptor(provider authLib.AuthProvider, logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info("Intercepting request", "method", info.FullMethod)

		// Extract metadata from context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Error("Missing metadata", "method", info.FullMethod)
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		// Get the authorization header
		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			logger.Error("Missing authorization token", "method", info.FullMethod)
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization token")
		}

		// Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		if tokenString == authHeader[0] {
			logger.Error("Malformed authorization token", "method", info.FullMethod)
			return nil, status.Errorf(codes.Unauthenticated, "malformed authorization token")
		}

		// Authorize the token and the route
		userDetails, err := provider.Authorize(ctx, info.FullMethod, tokenString)
		if err != nil {
			logger.Error("Authorization failed", "method", info.FullMethod, "error", err)
			return nil, status.Errorf(codes.PermissionDenied, "authorization failed: %v", err)
		}

		// Log successful authorization
		logger.Info("Authorization succeeded", "method", info.FullMethod, "user", userDetails.Username)

		// Add user details to context
		ctx = context.WithValue(ctx, authLib.UserDetailsKey, userDetails)

		// Call the handler
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Error("Handler error", "method", info.FullMethod, "error", err)
			return nil, err
		}

		logger.Info("Request handled successfully", "method", info.FullMethod)
		return resp, nil
	}
}
