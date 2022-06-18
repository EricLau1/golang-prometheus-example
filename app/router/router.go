package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang-prometheus-example/app/handlers"
	"golang-prometheus-example/app/middlewares"
	"net/http"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.Metrics)
	router.NotFoundHandler = http.HandlerFunc(handlers.InvalidPath)
	router.MethodNotAllowedHandler = http.HandlerFunc(handlers.InvalidMethod)
	router.Path("/metrics").Handler(promhttp.Handler())
	router.Path("/").HandlerFunc(handlers.Index)
	return router
}
