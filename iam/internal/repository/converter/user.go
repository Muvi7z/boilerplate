package converter

import (
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	entity2 "github.com/Muvi7z/boilerplate/iam/internal/repository/entity"
)

func UserToRepository() {

}

func UserFromRepository(user entity2.User) entity.User {
	return entity.User{
		Uuid:                user.Uuid,
		Email:               user.Email,
		Login:               user.Login,
		Password:            user.Password,
		NotificationMethods: nil,
	}
}

func NotificationMethodFromRepository(notificationMethod entity2.NotificationMethod) entity.NotificationMethod {
	return entity.NotificationMethod{
		ProviderName: notificationMethod.ProviderName,
		Target:       notificationMethod.Target,
	}
}
