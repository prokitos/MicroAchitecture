package services

import (
	"encoding/json"
	"io"
	"module/internal/database/dao"
	"module/internal/models"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// обновление записи
func CarUpdate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Car
	json.Unmarshal(reqBody, &curCar)

	// если пользователь не ввёл айди изменяемой записи, то ошибка
	if _, err := strconv.Atoi(curCar.Id); err != nil {
		log.Debug("don't correct id of updated car")
		models.BadClientResponse400(w)
		return
	}

	// возвращаем 200 если операции прошли хорошо
	res := dao.UpdateData(w, &curCar)
	if res {
		models.GoodResponse(w)
	}

}

// Удаление записи по айди
func CarDelete(w *http.ResponseWriter, r *http.Request) {

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	id := r.FormValue("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		models.BadClientResponse400(w)
		return
	}

	// возвращаем 200 если операции прошли хорошо
	res := dao.DeleteData(w, id)
	if res {
		models.GoodResponse(w)
	}

}

// Показать записи
func CarShow(w *http.ResponseWriter, r *http.Request) {

	var curCar models.Car
	curCar.Id = r.FormValue("id")
	curCar.RegNum = r.FormValue("regNum")
	curCar.Mark = r.FormValue("mark")
	curCar.Model = r.FormValue("model")
	curCar.Year = r.FormValue("year")
	curCar.Owner.Name = r.FormValue("name")
	curCar.Owner.Surname = r.FormValue("surname")
	curCar.Owner.Patronymic = r.FormValue("patronymic")

	var sorted string = r.FormValue("sort")
	var offset string = r.FormValue("offset")
	var limit string = r.FormValue("limit")

	// проверка что в offset,limit,id либо пустота либо числа.
	if _, err := strconv.Atoi(offset); offset != "" && err != nil {
		log.Debug("offset params couldn't convert to a number, offset = " + offset)

		models.BadClientResponse400(w)
		return
	}
	if _, err := strconv.Atoi(limit); limit != "" && err != nil {
		log.Debug("limit params couldn't convert to a number, limit = " + limit)

		models.BadClientResponse400(w)
		return
	}
	if _, err := strconv.Atoi(curCar.Id); curCar.Id != "" && err != nil {
		log.Debug("id params couldn't convert to a number, id = " + curCar.Id)

		models.BadClientResponse400(w)
		return
	}

	if curCar.Id == "" && curCar.Mark == "" && curCar.Model == "" && curCar.RegNum == "" && curCar.Year == "" &&
		curCar.Owner.Name == "" && curCar.Owner.Surname == "" && curCar.Owner.Patronymic == "" {
		log.Debug("empty all field in request")

		models.BadClientResponse400(w)
		return
	}

	// возвращаем 200 и машины если операции прошли хорошо
	cars, res := dao.ShowData(w, &curCar, offset, limit, sorted)
	if res {
		models.GoodShowResponse(w, cars)
	}

}

// Создать новую запись
func CarCreate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var nomera models.CarNumber
	json.Unmarshal(reqBody, &nomera)

	if len(nomera.RegNum) == 0 {
		log.Debug("empty regNum in insert request")
		models.BadClientResponse400(w)
		return
	}

	// так как внешнее апи получает только один номер, то отправляем номера по одному
	// в цикле каждый номер отправляем во внешнее апи, получаем carModel, и добавляем её в бд.
	for _, i := range nomera.RegNum {
		curCar := sendRequestToGet(i)

		// ошибка если внешний сервер вернул пустые поля машины
		if curCar.Mark == "" && curCar.Model == "" && curCar.Owner.Name == "" && curCar.Owner.Surname == "" {
			log.Error("error in external server request")
			log.Debug("external server return empty data")

			models.BadServerResponse(w)
			return
		}

		// выполнение каждоый операции, и проверка что они выполнены.
		res := dao.CreateData(w, &curCar)
		if !res {
			return
		}
	}

	duration := time.Second * 5
	time.Sleep(duration)

	//  возвращаем 200 если операции прошли хорошо
	models.GoodResponse(w)
}
