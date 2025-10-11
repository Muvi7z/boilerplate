package server

import (
	"context"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

type Server struct {
}

func (s *Server) PostApiV1Orders(ctx context.Context, request generated.PostApiV1OrdersRequestObject) (generated.PostApiV1OrdersResponseObject, error) {
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

func (s *Server) GetApiV1OrdersOrderUuid(ctx context.Context, request generated.GetApiV1OrdersOrderUuidRequestObject) (generated.GetApiV1OrdersOrderUuidResponseObject, error) {
	return generated.GetApiV1OrdersOrderUuid500JSONResponse{
		N5xxJSONResponse: struct {
			Code      *int    `json:"code,omitempty"`
			Message   string  `json:"message"`
			RequestId *string `json:"request_id,omitempty"`
		}{
			Message: "not implemented",
		},
	}, nil
}

func (s *Server) PostApiV1OrdersOrderUuidCancel(ctx context.Context, request generated.PostApiV1OrdersOrderUuidCancelRequestObject) (generated.PostApiV1OrdersOrderUuidCancelResponseObject, error) {
	return generated.PostApiV1OrdersOrderUuidCancel500JSONResponse{
		N5xxJSONResponse: struct {
			Code      *int    `json:"code,omitempty"`
			Message   string  `json:"message"`
			RequestId *string `json:"request_id,omitempty"`
		}{
			Message: "not implemented",
		},
	}, nil
}

func (s *Server) PostApiV1OrdersOrderUuidPay(ctx context.Context, request generated.PostApiV1OrdersOrderUuidPayRequestObject) (generated.PostApiV1OrdersOrderUuidPayResponseObject, error) {
	return generated.PostApiV1OrdersOrderUuidPay500JSONResponse{
		N5xxJSONResponse: struct {
			Code      *int    `json:"code,omitempty"`
			Message   string  `json:"message"`
			RequestId *string `json:"request_id,omitempty"`
		}{
			Message: "not implemented",
		},
	}, nil
}
