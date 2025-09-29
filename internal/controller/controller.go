package controller

import "github.com/Nonameipal/P2P/internal/contracts"

type Controller struct {
	service contracts.ServiceI
}

func NewController(svc contracts.ServiceI) *Controller {
	return &Controller{
		service: svc,
	}
}
