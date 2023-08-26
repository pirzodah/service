package http

import (
	"avesto-service/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// @Summary Ping the server
// @Description Returns a "pong" response
// @Tags ping
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (s *server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp models.Response
	)
	resp.Send(w, models.SUCCESS, "Pong")
	return
}
