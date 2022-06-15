package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type TimeResponse struct {
	Zone  string `json:"time_zone"`
	Value string `json:"time_value"`
}

type ErrorStruct struct {
	Message string `json:"message"`
}

func getDefaultTime(w http.ResponseWriter, r *http.Request) {
	actualTime := time.Now()
	getTime := TimeResponse{"UTC", actualTime.String()}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getTime)
}

func getWithTimeZone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramSplit := strings.Split(params["tz"], ",")
	if len(paramSplit) == 1 {
		response, err := getTimeWithTZ(paramSplit[0])
		if err != nil {
			w.WriteHeader(404)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		var response []TimeResponse
		for _, v := range paramSplit {
			actualTime, _ := getTimeWithTZ(v)
			response = append(response, actualTime)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func getTimeWithTZ(timezone string) (TimeResponse, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		response := TimeResponse{"Error time zone invalid", ""}
		return response, err
	} else {
		actualTime := time.Now().In(loc).String()
		response := TimeResponse{loc.String(), actualTime}
		return response, err
	}
}
