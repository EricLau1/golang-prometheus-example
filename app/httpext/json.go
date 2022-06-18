package httpext

import (
	"encoding/json"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set(HeaderContentType, MimeJSON)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("error on write json:", err.Error())
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := new(Error)
	if err != nil {
		e.Description = err.Error()
	}
	Write(w, statusCode, e)
}
