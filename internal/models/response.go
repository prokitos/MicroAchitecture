package models

import (
	"encoding/json"
	"net/http"
)

// структура для вывода хорошего ответа
type GoodBasicResponse struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"               example:"status"`
}

// структура для вывода хорошего ответа с машиной
type GoodAdvancedResponse struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"               example:"status"`
	Cars        []Car  `json:"cars"               example:"...."`
}

// структура для вывода плохого ответа
type ErrorResponse struct {
	Description string `json:"description"`
	Code        int    `json:"code"`
}

// ошибка на сервере
func BadServerResponse(w *http.ResponseWriter) {
	badResponse := ErrorResponse{
		Description: "Internal server error",
		Code:        500,
	}
	json.NewEncoder(*w).Encode(badResponse)
}

// обишка на клиенте
func BadClientResponse400(w *http.ResponseWriter) {
	badResponse := ErrorResponse{
		Description: "Bad request",
		Code:        400,
	}
	json.NewEncoder(*w).Encode(badResponse)
}
func BadClientResponse404(w *http.ResponseWriter) {
	badResponse := ErrorResponse{
		Description: "Records not Found",
		Code:        404,
	}
	json.NewEncoder(*w).Encode(badResponse)
}

// вывод что все прошло хорошо
func GoodResponse(w *http.ResponseWriter) {
	goodResp := GoodBasicResponse{
		Description: "Ok",
		Code:        200,
	}
	json.NewEncoder(*w).Encode(goodResp)
}

// вывод что все прошло хорошо + показать машину
func GoodShowResponse(w *http.ResponseWriter, resultCars *[]Car) {
	goodResp := GoodAdvancedResponse{
		Description: "Ok",
		Code:        200,
		Cars:        *resultCars,
	}
	json.NewEncoder(*w).Encode(goodResp)
}
