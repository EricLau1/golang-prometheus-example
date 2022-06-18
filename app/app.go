package app

import (
	"flag"
	"github.com/gorilla/handlers"
	"golang-prometheus-example/app/exit"
	"golang-prometheus-example/app/httpext"
	"golang-prometheus-example/app/router"
	"log"
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
	handler := handlers.LoggingHandler(os.Stdout, router.New())
	httpext.Listen(port, router.WithCORS(handler), func() {
		log.Printf("Listening http://localhost:%d\n\n", port)
	})
}
