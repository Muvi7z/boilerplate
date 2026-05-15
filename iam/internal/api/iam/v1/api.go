package v1

import "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"

type api struct {
	iam_v1.UnimplementedIAMServiceServer
	sessionService sessionService
}

func NewApi(sessionService sessionService) *api {
	return &api{
		sessionService: sessionService,
	}
}
