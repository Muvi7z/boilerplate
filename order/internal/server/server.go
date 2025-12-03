package server

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/handler/order"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
	"log"
	"net/http"
)

type Server struct {
	orderHandler *order.Handler
	addr         string
	server       *http.Server
}

func NewServer(orderHandler *order.Handler, addr string) *Server {
	s := &Server{
		orderHandler: orderHandler,
		addr:         addr,
	}
	mux := http.NewServeMux()

	strictHandler := generated.NewStrictHandler(s, nil)
	handler := generated.Handler(strictHandler)

	mux.Handle("/", handler)

	s.server = &http.Server{Addr: s.addr, Handler: mux}

	return s
}

func (s *Server) Run() {
	log.Println("Starting server")

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) PostApiV1Orders(ctx context.Context, request generated.PostApiV1OrdersRequestObject) (generated.PostApiV1OrdersResponseObject, error) {
	return s.orderHandler.CreateOrder(ctx, request)
}

func (s *Server) GetApiV1OrdersOrderUuid(ctx context.Context, request generated.GetApiV1OrdersOrderUuidRequestObject) (generated.GetApiV1OrdersOrderUuidResponseObject, error) {
	return s.orderHandler.GetOrder(ctx, request)
}

func (s *Server) PostApiV1OrdersOrderUuidCancel(ctx context.Context, request generated.PostApiV1OrdersOrderUuidCancelRequestObject) (generated.PostApiV1OrdersOrderUuidCancelResponseObject, error) {
	return s.orderHandler.CancelOrder(ctx, request)
}

func (s *Server) PostApiV1OrdersOrderUuidPay(ctx context.Context, request generated.PostApiV1OrdersOrderUuidPayRequestObject) (generated.PostApiV1OrdersOrderUuidPayResponseObject, error) {
	return s.orderHandler.PayOrder(ctx, request)
}
