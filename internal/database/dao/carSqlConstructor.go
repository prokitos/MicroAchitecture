package dao

import (
	"module/internal/models"
	"strconv"
)

// конструктор для создания where части запроса
func constructWhere(current models.Car) string {

	var result string = "where "
	if current.Id != "" {
		result += "car.id = '" + current.Id + "' and "
	}
	if current.Mark != "" {
		result += "mark = '" + current.Mark + "' and "
	}
	if current.Model != "" {
		result += "model = '" + current.Model + "' and "
	}
	if current.RegNum != "" {
		result += "regNum = '" + current.RegNum + "' and "
	}
	if current.Year != "" {
		result += "year = '" + current.Year + "' and "
	}
	if current.Owner.Name != "" {
		result += "name = '" + current.Owner.Name + "' and "
	}
	if current.Owner.Surname != "" {
		result += "surname = '" + current.Owner.Surname + "' and "
	}
	if current.Owner.Patronymic != "" {
		result += "patronymic = '" + current.Owner.Patronymic + "' and "
	}

	newRes := result[:len(result)-4]
	return newRes
}

// конструктор для создания set части запроса машин
func constructSetCar(current *models.Car) string {

	var result string = "SET "
	if current.Id != "" {
		result += "id = '" + current.Id + "' , "
	}
	if current.Mark != "" {
		result += "mark = '" + current.Mark + "' , "
	}
	if current.Model != "" {
		result += "model = '" + current.Model + "' , "
	}
	if current.RegNum != "" {
		result += "regNum = '" + current.RegNum + "' , "
	}
	if current.Year != "" {
		result += "year = '" + current.Year + "' , "
	}
	newRes := result[:len(result)-2]
	return newRes
}

// конструктор для создания set части запроса людей
func constructSetPeople(current *models.Car) string {

	var result string = "SET "
	if current.Owner.Name != "" {
		result += "name = '" + current.Owner.Name + "' , "
	}
	if current.Owner.Surname != "" {
		result += "surname = '" + current.Owner.Surname + "' , "
	}
	if current.Owner.Patronymic != "" {
		result += "patronymic = '" + current.Owner.Patronymic + "' , "
	}
	newRes := result[:len(result)-2]
	return newRes
}

// конструктор show
func conStringShowSpec(car *models.Car, limit string, offset string, sort string) string {

	var stroka string = "SELECT car.id,regnum,mark,model,year,people.id,name,surname,patronymic FROM car left join people on car.owner = people.id "

	newRes := stroka + constructWhere(*car)

	if sort == "id" || sort == "Id" {
		sort = "car.id"
	}

	if len(sort) > 0 {
		newRes += "order by " + sort + " "
	}
	test, _ := strconv.Atoi(limit)
	if test > 0 {
		newRes += "limit " + limit + " "
	}
	test, _ = strconv.Atoi(offset)
	if test > 0 {
		newRes += "offset " + offset + " "
	}

	return newRes
}

// конструктор update car
func conStringUpdateCar(car *models.Car) string {

	if car.Mark == "" && car.Model == "" && car.RegNum == "" && car.Year == "" {
		return ""
	}

	var stroka string = "UPDATE car "

	newRes := stroka + constructSetCar(car)
	newRes += "where id = " + car.Id

	return newRes
}

// конструктор update people
func conStringUpdatePeople(car *models.Car) string {

	if car.Owner.Name == "" && car.Owner.Surname == "" && car.Owner.Patronymic == "" {
		return ""
	}

	var stroka string = "UPDATE people "

	newRes := stroka + constructSetPeople(car)
	newRes += "from car "
	newRes += "where car.owner = people.id and car.id = " + car.Id

	return newRes
}

// конструктор insert people
func conStringInsertPeople(car *models.Car) string {

	var stroka string = "INSERT INTO people "
	var first = "("
	var second = "("

	if car.Owner.Name != "" {
		first += "name, "
		second += "'" + car.Owner.Name + "', "
	}
	if car.Owner.Surname != "" {
		first += "surname, "
		second += "'" + car.Owner.Surname + "', "
	}
	if car.Owner.Patronymic != "" {
		first += "patronymic, "
		second += "'" + car.Owner.Patronymic + "', "
	}

	first = first[:len(first)-2]
	second = second[:len(second)-2]

	first += ")"
	second += ")"

	stroka += first + " values " + second + " returning id"

	return stroka
}

// конструктор insert car
func conStringInsertCar(car *models.Car) string {

	var stroka string = "INSERT INTO car "
	var first = "("
	var second = "("

	if car.RegNum != "" {
		first += "regnum, "
		second += "'" + car.RegNum + "', "
	}
	if car.Mark != "" {
		first += "mark, "
		second += "'" + car.Mark + "', "
	}
	if car.Model != "" {
		first += "model, "
		second += "'" + car.Model + "', "
	}
	if car.Year != "" {
		first += "year, "
		second += "'" + car.Year + "', "
	}
	if car.Owner.Id != "" {
		first += "owner, "
		second += "'" + car.Owner.Id + "', "
	}

	first = first[:len(first)-2]
	second = second[:len(second)-2]

	first += ")"
	second += ")"

	stroka += first + " values " + second

	return stroka
}
