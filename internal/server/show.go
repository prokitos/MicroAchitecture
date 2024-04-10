package server

import (
	"module/internal/services"
	"net/http"
)

// Cars godoc
// @Summary Show cars and person
// @Description Show cars and his owner
// @Tags cars
// @Accept  json
// @Produce  json
// @Param sort query string false "Sort records"
// @Param limit query int false "Show max limit records"
// @Param offset query int false "Show records with current offset"
// @Param cars query carShow false "Show cars"
// @Param name query string false "Example: Ivan"
// @Param surname query string false "Example: Ivanov"
// @Param patronymic query string false "Example: Ivanovich"
// @Failure 400 "Invalid parameters supplied"
// @Failure 404 "Cars not found"
// @Success 200 {object} carShow "successful operation"
// @Router /show [get]
func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// здесь провекра доступа

	services.CarShow(&w, r)
}

// копия оригинальных моделей, но с "example" для добавления в сваггер
type carShow struct {
	Id     string     `json:"id" example:"12"`
	RegNum string     `json:"regNum" example:"x16xx150"`
	Mark   string     `json:"mark" example:"lada"`
	Model  string     `json:"model" example:"kalina"`
	Year   string     `json:"year" example:"2000"`
	Owner  peopleShow `json:"owner"`
}

type peopleShow struct {
	Id         string `json:"id" example:"14"`
	Name       string `json:"name" example:"james"`
	Surname    string `json:"surname" example:"johnson"`
	Patronymic string `json:"patronymic" example:"petrovich"`
}
