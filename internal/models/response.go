package models

import (
	"encoding/json"
	"net/http"
)

// Response struct
type Response struct {
	Total   int         `json:"total,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	ErrorCode
}

// Send ...
func (resp *Response) Send(w http.ResponseWriter, errCode ErrorCode, payload interface{}) {
	resp.SendWithTotal(w, errCode, 0, payload)
}

// SendWithTotal ...
func (resp *Response) SendWithTotal(w http.ResponseWriter, errCode ErrorCode, total int, payload interface{}) {
	resp.Code = httpCodes[errCode.Code]
	resp.Message = errCode.Message
	if len(resp.Message) == 0 {
		resp.Message = "Неизвестная ошибка"
	}
	resp.Total = total
	resp.Payload = payload
	w.WriteHeader(resp.Code)
	w.Header().Set("Content-type", "application/json; charset=utf8")
	json.NewEncoder(w).Encode(resp)
}

var BadJSONerrorCode = ErrorCode{
	Code:    BADREQUEST.Code,
	Message: "Неправильный json!",
}

var httpCodes = map[int]int{
	SUCCESS.Code:       http.StatusOK,                  // 200
	DUPLICATE.Code:     http.StatusAlreadyReported,     // 208
	BADREQUEST.Code:    http.StatusBadRequest,          // 400
	UNAUTHORIZED.Code:  http.StatusUnauthorized,        // 401
	FORBIDDEN.Code:     http.StatusForbidden,           // 403
	NOTFOUND.Code:      http.StatusNotFound,            // 404
	INTERNALERROR.Code: http.StatusInternalServerError, // 500
	0:                  http.StatusNotImplemented,      // 501
}

// ErrorCode ...
type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error Codes
var (
	SUCCESS       = ErrorCode{Code: 1, Message: "Успешно"}
	DUPLICATE     = ErrorCode{Code: 2, Message: "Уже существует"}
	BADREQUEST    = ErrorCode{Code: 3, Message: "Неверный запрос"}
	UNAUTHORIZED  = ErrorCode{Code: 4, Message: "Неавторизован"}
	FORBIDDEN     = ErrorCode{Code: 5, Message: "Доступ ограничен"}
	NOTFOUND      = ErrorCode{Code: 6, Message: "Не найдено"}
	INTERNALERROR = ErrorCode{Code: 7, Message: "Внутренняя ошибка"}
)
