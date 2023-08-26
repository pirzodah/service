package http

import (
	"avesto-service/cmd/app/controllers"
	"avesto-service/internal/models"
	"avesto-service/internal/usecase"
	pkghttp "avesto-service/pkg/controller/http"
	"context"
	"net/http"
)

type server struct {
	usecase usecase.Usecase
	srv     pkghttp.Server
}

func NewController(usecase usecase.Usecase, srv pkghttp.Server) controllers.Controller {
	return &server{usecase: usecase, srv: srv}
}

func (s *server) Serve(ctx context.Context, config *models.Configuration) error {
	return s.srv.Serve(ctx, config, []pkghttp.Route{
		//PING
		{Method: http.MethodGet, Path: "/ping", Handler: s.ping},

		{Method: http.MethodPost, Path: "/service", Handler: s.AddService},
		{Method: http.MethodGet, Path: "/service", Handler: s.GetService},
		{Method: http.MethodPut, Path: "/service", Handler: s.UpdateService},
		{Method: http.MethodDelete, Path: "/service", Handler: s.DeleteService},

		{Method: http.MethodPost, Path: "/inputs", Handler: s.AddInputs},
		{Method: http.MethodGet, Path: "/inputs", Handler: s.GetInputs},
		{Method: http.MethodPut, Path: "/inputs", Handler: s.UpdateInputs},
		{Method: http.MethodDelete, Path: "/inputs", Handler: s.DeleteInputs},

		{Method: http.MethodPost, Path: "/serviceInputs", Handler: s.AddServiceInputs},
		//{Method: http.MethodPut, Path: "/serviceInputs", Handler: s.UpdateServiceInputs},
		{Method: http.MethodDelete, Path: "/serviceInputs", Handler: s.DeleteServiceInputs},
		{Method: http.MethodGet, Path: "/serviceInputs", Handler: s.GetServiceInputs},
	})
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
