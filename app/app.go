package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getWithTimeZone).Queries("tz", "{tz}")
	router.HandleFunc("/api/time", getDefaultTime)

	http.ListenAndServe("localhost:8080", router)
}
