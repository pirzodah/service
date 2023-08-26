package http

import (
	"avesto-service/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *server) AddService(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp    models.Response
		service models.Services
	)
	service.Name = r.FormValue("name")
	errCode := s.usecase.AddServices(r.Context(), service)

	resp.Send(w, errCode.ErrorCode, nil)
	return
}

func (s *server) GetService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp     models.Response
		services models.Services
	)

	services.ID, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	service, errCode := s.usecase.GetServiceByID(r.Context(), services.ID)

	resp.Send(w, errCode.ErrorCode, service)
	return
}

func (s *server) UpdateService(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp    models.Response
		service models.Services
		err     error
	)

	// Получаем ID сервиса из параметров маршрута
	service.ID, err = strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid service ID")
		return
	}

	service.Name = r.FormValue("name")

	errCode := s.usecase.UpdateService(r.Context(), service)
	resp.Send(w, errCode.ErrorCode, nil)
}

func (s *server) DeleteService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp models.Response
	)

	// Получаем ID сервиса из параметров маршрута
	serviceID, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid service ID")
		return
	}

	// Вызываем метод usecase для удаления сервиса
	errCode := s.usecase.DeleteService(r.Context(), serviceID)
	resp.Send(w, errCode.ErrorCode, nil)
}
