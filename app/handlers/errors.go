package handlers

import (
	"fmt"
	"golang-prometheus-example/app/httpext"
	"net/http"
)

func InvalidMethod(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("invalid method: %s", r.Method)
	httpext.WriteError(w, http.StatusMethodNotAllowed, err)
}

func InvalidPath(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("invalid URI: %s", r.RequestURI)
	httpext.WriteError(w, http.StatusNotFound, err)
}
