package api

import "github.com/mohamedveron/paytabs_task/service"

type Server struct {
	svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}
