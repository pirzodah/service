package http

import (
	"avesto-service/internal/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *server) AddServiceInputs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp        models.Response
		serviceData struct {
			ServiceID int64   `json:"service_id"`
			Required  bool    `json:"required"`
			InputIDs  []int64 `json:"input_ids"`
		}
		serviceInputs models.ServiceInputs

		err error
	)
	// Декодирование JSON данных из тела запроса
	err = json.NewDecoder(r.Body).Decode(&serviceData)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid JSON data")
		return
	}

	serviceInputs.InputIDs = serviceData.InputIDs
	serviceInputs.ServiceID = serviceData.ServiceID
	serviceInputs.Required = serviceData.Required
	// Вызов метода use-case для добавления данных
	errCode := s.usecase.AddServiceInputs(r.Context(), serviceInputs)
	if errCode.Code != 1 {
		resp.Send(w, errCode.ErrorCode, errCode.Message)
		return
	}

	resp.Send(w, models.SUCCESS, nil) // Отправка успешного ответа
}

func (s *server) GetServiceInputs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp models.Response
	)

	serviceID, _ := strconv.ParseInt(r.FormValue("service_id"), 10, 64)
	input, errCode := s.usecase.GetServiceInputs(r.Context(), serviceID)

	resp.Send(w, errCode.ErrorCode, input)
	return
}

//func (s *server) UpdateServiceInputs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	var (
//		resp  models.Response
//		input models.ServiceInputs
//		err   error
//	)
//
//	// Парсинг данных из формы
//	input.ID, err = strconv.ParseInt(r.FormValue("id"), 10, 64)
//	if err != nil {
//		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid input ID")
//		return
//	}
//	input.InputIDs = []int64{input.ID} // Здесь вы можете передать несколько input IDs, если это актуально
//	input.InputIDs, err = strconv.ParseInt(r.FormValue("input_id"), 10, 64)
//	if err != nil {
//		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid input_id value")
//		return
//	}
//	input.ServiceID, err = strconv.ParseInt(r.FormValue("service_id"), 10, 64)
//	if err != nil {
//		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid service_id value")
//		return
//	}
//	input.Required, err = strconv.ParseBool(r.FormValue("required"))
//	if err != nil {
//		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid required value")
//		return
//	}
//
//	// Вызов метода use-case для обновления данных
//	errCode := s.usecase.UpdateServiceInputs(r.Context(), input)
//	if errCode != nil {
//		// Обработка ошибок и отправка соответствующего HTTP-статуса
//		resp.Send(w, errCode.ErrorCode, errCode.Message)
//		return
//	}
//
//	resp.Send(w, models.ErrorCode{}, nil) // Отправка успешного ответа
//}

func (s *server) DeleteServiceInputs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resp models.Response
	)

	inputID, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		resp.Send(w, models.ErrorCode{Code: 7}, "Invalid input ID")
		return
	}

	errCode := s.usecase.DeleteServiceInputs(r.Context(), inputID)
	resp.Send(w, errCode.ErrorCode, nil)
}
