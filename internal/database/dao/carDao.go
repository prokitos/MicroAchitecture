package dao

import (
	"module/internal/database"
	"module/internal/models"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Обновление записи. возвращает false если неудачно
func UpdateData(w *http.ResponseWriter, curModel *models.Car) bool {

	if curModel.Id == "" {
		return false
	}

	// сначала обновляем машину
	query := conStringUpdateCar(curModel)
	if query != "" {
		err := executeToDB(w, query)
		if !err {
			return false
		}
	}

	// потом обновляем владельца
	query = conStringUpdatePeople(curModel)
	if query != "" {
		err := executeToDB(w, query)
		if !err {
			return false
		}
	}

	// возвращаем true если всё прошло успешно
	return true
}

// Удаление записи по айди. возвращает false если неудачно
func DeleteData(w *http.ResponseWriter, id string) bool {

	// получаем айдишник владельца машины
	var query string = "SELECT owner from car where id = " + id
	temp, err := executeReturnToDB(w, query)
	if !err {
		return false
	}

	// каскадно удаляем владельца и машину
	query = "delete from people where id = " + temp
	err = executeToDB(w, query)
	if !err {
		return false
	}

	// возвращаем true если всё прошло успешно
	return true
}

// Показать записи. возвращает false если неудачно
func ShowData(w *http.ResponseWriter, curModel *models.Car, offset string, limit string, sort string) (*[]models.Car, bool) {

	// показать записи по разным параметрам
	var query string = conStringShowSpec(curModel, limit, offset, sort)
	carMas, err := showFromDB(w, query)
	if !err {
		return carMas, false
	}

	// возвращаем true и машины если всё прошло успешно
	return carMas, true
}

// Создать новую запись. возвращает false если неудачно
func CreateData(w *http.ResponseWriter, curCar *models.Car) bool {

	var err bool

	// добавляем сначала владельца, и получаем его айдишник
	var query string = conStringInsertPeople(curCar)
	curCar.Owner.Id, err = executeReturnToDB(w, query)
	if !err {
		return false
	}

	// добавляем машину, и указывает айдишник владельца
	query = conStringInsertCar(curCar)
	err = executeToDB(w, query)
	if !err {
		return false
	}

	// возвращаем true если всё прошло успешно
	return true
}

// просто выполнить запрос. вернуть false если неудачно
func executeToDB(w *http.ResponseWriter, conn string) bool {

	res, err := database.GlobalInstance.GetConnect().Exec(conn)
	models.CheckError(err)

	affected, _ := res.RowsAffected()
	if affected == 0 {
		log.Debug("nothind to execute. 0 records")
		models.BadClientResponse404(w)
		return false
	}

	return true
}

// выполнить запрос, и вернуть айди записи, над которой был запрос. также вернуть false при неудаче
func executeReturnToDB(w *http.ResponseWriter, conn string) (string, bool) {

	lastInsertId := "0"
	database.GlobalInstance.GetConnect().QueryRow(conn).Scan(&lastInsertId)

	if lastInsertId == "0" {
		log.Debug("nothind to execute. 0 records")
		models.BadClientResponse404(w)
		return lastInsertId, false
	}

	return lastInsertId, true
}

// показать данные из таблицы. вернуть false если неудачно
func showFromDB(w *http.ResponseWriter, stroka string) (*[]models.Car, bool) {

	rows, _ := database.GlobalInstance.GetConnect().Query(stroka)
	defer rows.Close()

	cars := []models.Car{}

	for rows.Next() {
		car := models.Car{}
		own := models.People{}
		err := rows.Scan(&car.Id, &car.RegNum, &car.Model, &car.Mark, &car.Year, &own.Id, &own.Name, &own.Surname, &own.Patronymic)
		if err != nil {
			log.Error("database show data error")
			log.Debug("errored query when show data: " + stroka)
			models.BadServerResponse(w)
			return &cars, false
		}
		car.Owner = own
		cars = append(cars, car)
	}

	if len(cars) == 0 {
		log.Debug("nothing to show")
		models.BadClientResponse404(w)
		return &cars, false
	}

	return &cars, true
}
