package grpc

import (
	"context"
	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	SessionUUIDMetadatakey = "session-uuid"
)

type contextKey string

const (
	userContextKey        contextKey = "user"
	sessionUUIDContextKey contextKey = "session-uuid"
)

type IAMClient = iam_v1.IAMServiceClient

type AuthInterceptor struct {
	iamClient IAMClient
}

func NewAuthInterceptor(iamClient IAMClient) *AuthInterceptor {
	return &AuthInterceptor{iamClient: iamClient}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		authCtx, err := i.authenticate(ctx)
		if err != nil {
			return nil, err
		}

		return handler(authCtx, req)
	}
}

func (i *AuthInterceptor) authenticate(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	sessionUUIDs := md.Get(SessionUUIDMetadatakey)
	if len(sessionUUIDs) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "session is not provided")
	}

	sessionUUID := sessionUUIDs[0]

	if sessionUUID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "empty session-uuid")
	}

	whoamiRes, err := i.iamClient.Whoami(ctx, &iam_v1.WhoamiRequest{
		SessionUuid: sessionUUID,
	})

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "ivalid session: %v", err)
	}

	authCtx := context.WithValue(ctx, userContextKey, whoamiRes.UserUuid)
	authCtx = context.WithValue(authCtx, sessionUUIDContextKey, sessionUUID)
	return authCtx, nil
}

func GetUserContextKey() contextKey {
	return userContextKey
}

func GetSessionUUIDFromContext(ctx context.Context) (string, bool) {
	sessionUUID, ok := ctx.Value(sessionUUIDContextKey).(string)
	return sessionUUID, ok
}

func AddSessionUUIDToContext(ctx context.Context, sessionUUID string) context.Context {
	return context.WithValue(ctx, sessionUUIDContextKey, sessionUUID)
}

func ForwardSessionUUIDToGRPC(ctx context.Context) context.Context {
	sessionUUID, ok := GetSessionUUIDFromContext(ctx)
	if !ok || sessionUUID == "" {
		return ctx
	}

	return metadata.AppendToOutgoingContext(ctx, SessionUUIDMetadatakey, sessionUUID)
}
