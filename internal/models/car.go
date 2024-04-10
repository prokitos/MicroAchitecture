package models

// структура таблицы машин
type Car struct {
	Id     string `json:"id" example:"12"`
	RegNum string `json:"regNum" example:""`
	Mark   string `json:"mark" example:""`
	Model  string `json:"model" example:"tesla"`
	Year   string `json:"year" example:""`
	Owner  People `json:"owner"`
}

// структура с массивом номеров
type CarNumber struct {
	RegNum []string `json:"regNum" example:"x15xx150,x24xx134"`
}
