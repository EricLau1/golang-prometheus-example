package handlers

import (
	"golang-prometheus-example/app/httpext"
	"golang-prometheus-example/app/types"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var (
		status = http.StatusOK
		sleep  time.Duration
		start  = time.Now()
		query  = r.URL.Query()
		err    error
	)
	if query.Get("sleep") != "" {
		sleep, err = time.ParseDuration(query.Get("sleep"))
		if err != nil {
			httpext.WriteError(w, http.StatusBadRequest, err)
			return
		}
		time.Sleep(sleep)
	}
	if query.Get("status") != "" {
		status, err = strconv.Atoi(query.Get("status"))
		if err != nil {
			httpext.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	httpext.Write(w, status, types.Map{
		"message":  http.StatusText(status),
		"status":   status,
		"duration": time.Since(start).String(),
	})
}

func NotFound() {

}
