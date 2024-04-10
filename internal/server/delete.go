package server

import (
	"module/internal/services"
	"net/http"
)

// Cars godoc
// @Summary Delete car and person by car id
// @Description Delete car and perso by car id
// @Tags cars
// @Accept  json
// @Produce  json
// @Param id query int true "delete car by id"
// @Failure 400 "Invalid id supplied"
// @Failure 404 "Car not found"
// @Router /delete [delete]
func deleteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// здесь провекра доступа

	services.CarDelete(&w, r)
}
