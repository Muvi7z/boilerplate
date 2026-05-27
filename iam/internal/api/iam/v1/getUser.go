package v1

import (
	"context"

	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
)

func (a *Api) GetUser(ctx context.Context, request *iam_v1.GetUserRequest) (*iam_v1.GetUserResponse, error) {
	user, err := a.iamService.GetUser(ctx, request.UserUuid)
	if err != nil {
		return nil, err
	}

	var notificationMethods []*iam_v1.NotificationMethod
	for _, notification := range user.NotificationMethods {
		notificationMethods = append(notificationMethods, &iam_v1.NotificationMethod{
			ProviderName: notification.ProviderName,
			Target:       notification.Target,
		})
	}

	return &iam_v1.GetUserResponse{
		UserUuid:            user.Uuid,
		Login:               user.Login,
		Email:               user.Email,
		NotificationMethods: notificationMethods,
	}, nil
}
