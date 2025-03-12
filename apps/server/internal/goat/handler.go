package handler

import (
	"context"
	"fmt"
	goatv1 "goat/proto/goat/v1"
	"goat/proto/goat/v1/goatv1connect"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
)

type goatHandler struct{}

func NewGofastHandler() *goatHandler {
	return &goatHandler{}
}

func (h *goatHandler) Goat(ctx context.Context, req *connect.Request[goatv1.GoatRequest]) (*connect.Response[goatv1.GoatResponse], error) {

	fmt.Println("Received request")
	res := connect.NewResponse(&goatv1.GoatResponse{
		Sentence: "This Stack is Goat 🐐",
	})
	return res, nil
}

func RegisterConnect(r *chi.Mux) {
	goatServer := NewGofastHandler()

	path, handler := goatv1connect.NewGoatServiceHandler(goatServer)

	r.Group(func(r chi.Router) {
		r.Mount(path, handler)
	})
}
