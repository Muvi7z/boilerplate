package http

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/middleware/grpc"
	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
	"net/http"
)

const SessionUUIDHeader = "X-Session-UUID"

type IAMClient = iam_v1.IAMServiceClient

type AuthMiddleware struct {
	iamClient IAMClient
}

func NewAuthMiddleware(iamClient IAMClient) *AuthMiddleware {
	return &AuthMiddleware{iamClient: iamClient}
}

func (a *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionUUID := r.Header.Get(SessionUUIDHeader)
		if sessionUUID == "" {
			http.Error(w, "Session UUID not found in request header", http.StatusUnauthorized)
			return
		}

		whoamiRes, err := a.iamClient.Whoami(r.Context(), &iam_v1.WhoamiRequest{
			SessionUuid: sessionUUID,
		})
		if err != nil {
			http.Error(w, "Authorization failed", http.StatusInternalServerError)
		}

		ctx := r.Context()
		ctx = grpc.AddSessionUUIDToContext(ctx, sessionUUID)
		ctx = context.WithValue(ctx, grpc.GetUserContextKey(), whoamiRes.UserUuid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
