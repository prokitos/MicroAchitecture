package server

import (
	"log"
	"net/http"
	"time"

	_ "module/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// запуск сервера. localhost:8888
func MainServer() {

	router := routers()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// все пути
func routers() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/delete", deleteRequest).Methods(http.MethodDelete)
	router.HandleFunc("/insert", insertGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/update", updateGetRequest).Methods(http.MethodPut)
	router.HandleFunc("/show", showsSpecGetRequest).Methods(http.MethodGet)

	router.HandleFunc("/load", loadTesst).Methods(http.MethodGet)

	return router
}
