package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

//ResponseJSON response http request with application/json
func ResponseJSON(data interface{}, writer http.ResponseWriter) (err error) {
	writer.Header().Set("Content-type", "application/json")

	d, err := json.Marshal(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("ResponseJSON: Failed to response")
		return
	}

	writer.Write(d)
	return
}
