package v1

import (
	"context"

	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
)

func (a *Api) Whoami(ctx context.Context, request *iam_v1.WhoamiRequest) (*iam_v1.WhoamiResponse, error) {

	user, err := a.iamService.Whoami(ctx, request.SessionUuid)
	if err != nil {
		return nil, err
	}
	return &iam_v1.WhoamiResponse{
		UserUuid: user.Uuid,
		Login:    user.Login,
		Email:    user.Email,
	}, nil
}
