package server

import (
	"module/internal/services"
	"net/http"
)

// Cars godoc
// @Summary Insert person and car
// @Description Insert reg nums for car
// @Tags cars
// @Accept  json
// @Produce  json
// @Param regNum body models.CarNumber true "Add reg nums"
// @Success 200 "successful operation"
// @Router /insert [post]
func insertGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// здесь провекра доступа

	services.CarCreate(&w, r)
}
