package app

import (
	"flag"
	"github.com/gorilla/handlers"
	"golang-prometheus-example/app/exit"
	"golang-prometheus-example/app/httpext"
	"golang-prometheus-example/app/router"
	"log"
	"net/http"
	"os"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "api port")
	flag.Parse()
}

func New() {
	exit.Graceful(func() {
		log.Println("Stop.")
	})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "X-Requested-with"})
	origins := handlers.AllowedOrigins([]string{"*"})
	handler := handlers.LoggingHandler(os.Stdout, router.New())
	handler = handlers.CORS(methods, headers, origins)(handler)
	httpext.Listen(port, handler, func() {
		log.Printf("Listening http://localhost:%d\n\n", port)
	})
}
