package v1

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
)

func (a *Api) Register(ctx context.Context, request *iam_v1.RegisterRequest) (*iam_v1.RegisterResponse, error) {
	var notificationMethods []entity.NotificationMethod

	for _, notificationMethod := range request.NotificationMethods {
		notificationMethods = append(notificationMethods, entity.NotificationMethod{
			ProviderName: notificationMethod.ProviderName,
			Target:       notificationMethod.Target,
		})
	}

	userReq := entity.User{
		Email:               request.Email,
		Login:               request.Login,
		Password:            request.Password,
		NotificationMethods: notificationMethods,
	}

	userUUID, err := a.iamService.Register(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return &iam_v1.RegisterResponse{
		UserUuid: userUUID,
	}, nil

}
