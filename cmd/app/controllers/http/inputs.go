package http

import (
	"avesto-service/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *server) AddInputs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp    models.Response
		service models.Inputs
		err     error
	)
	service.Name = r.FormValue("name")
	service.TypeInputs, err = strconv.ParseInt(r.FormValue("type_inputs"), 10, 64)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid type_inputs")
		return
	}
	errCode := s.usecase.AddInputs(r.Context(), service)

	resp.Send(w, errCode.ErrorCode, nil)
	return
}

func (s *server) GetInputs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp     models.Response
		services models.Services
	)

	services.ID, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	service, errCode := s.usecase.GetInputs(r.Context(), services.ID)

	resp.Send(w, errCode.ErrorCode, service)
	return
}

func (s *server) UpdateInputs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp    models.Response
		service models.Inputs
		err     error
	)
	service.Name = r.FormValue("name")
	service.ID, err = strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid service ID")
		return
	}
	service.TypeInputs, err = strconv.ParseInt(r.FormValue("type_inputs"), 10, 64)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid type_inputs")
		return
	}

	errCode := s.usecase.UpdateInputs(r.Context(), service)
	resp.Send(w, errCode.ErrorCode, nil)
}

func (s *server) DeleteInputs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	errCode := s.usecase.DeleteInputs(r.Context(), serviceID)
	resp.Send(w, errCode.ErrorCode, nil)
}
