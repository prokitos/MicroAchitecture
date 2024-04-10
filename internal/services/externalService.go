package services

import (
	"encoding/json"
	"io/ioutil"
	"module/internal/database"
	"module/internal/models"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// отправка номеров на внешний сервер по роуту /info
func sendRequestToGet(p_num string) models.Car {

	var strConnect string = ""
	database.GetExternalRoutes(&strConnect)

	baseURL, _ := url.Parse(strConnect + "/info")
	params := url.Values{}
	params.Add("regNum", p_num)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Debug("Error connecting to external api")
		log.Error("Error getting data from api")

		var empty models.Car
		return empty
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var carExemp models.Car
	json.Unmarshal(body, &carExemp)
	return carExemp
}
