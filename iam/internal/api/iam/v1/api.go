package v1

import "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"

type Api struct {
	iam_v1.UnimplementedIAMServiceServer
	iamService IamService
}

func NewApi(iamService IamService) *Api {
	return &Api{
		iamService: iamService,
	}
}
