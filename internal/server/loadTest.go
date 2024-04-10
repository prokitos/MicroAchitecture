package server

import (
	"bytes"
	"net/http"
)

func loadTesst(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 100; i++ {
		go sendRequestToInsert()
	}

}

// отправка кучи запросов. стресс тест!!?
func sendRequestToInsert() {

	data := []byte(`{
		"regNum": [
		  "xxxxxx"
		]
	  }`)
	r := bytes.NewReader(data)
	_, err := http.Post("http://localhost:8888"+"/insert", "application/json", r)
	if err != nil {
		return
	}
}
