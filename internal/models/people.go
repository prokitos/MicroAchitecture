package models

// структура таблицы владельцев
type People struct {
	Id         string `json:"id" example:""`
	Name       string `json:"name" example:"jamson"`
	Surname    string `json:"surname" example:""`
	Patronymic string `json:"patronymic" example:""`
}
