package order

import (
	"context"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func (h *Handler) Create(ctx context.Context, request generated.PostApiV1OrdersRequestObject) (generated.PostApiV1OrdersResponseObject, error) {
	return generated.PostApiV1Orders500JSONResponse{
		N5xxJSONResponse: struct {
			Code      *int    `json:"code,omitempty"`
			Message   string  `json:"message"`
			RequestId *string `json:"request_id,omitempty"`
		}{
			Message: "not implemented",
		},
	}, nil
}
