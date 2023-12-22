package http

import "github.com/monegim/simple-microservice-in-go/rating/internal/controller/rating"

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl}
}
