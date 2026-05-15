package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) Login(ctx context.Context, loginRequest *iam_v1.LoginRequest) (*iam_v1.LoginResponse, error) {
	if loginRequest.Login == "" || loginRequest.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password required")
	}

	userReq := entity.User{
		Login:    loginRequest.Login,
		Password: loginRequest.Password,
	}
	sessionUUID, err := a.sessionService.Login(ctx, userReq)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	return &iam_v1.LoginResponse{
		SessionUuid: sessionUUID,
	}, nil
}
