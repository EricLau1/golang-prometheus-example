package httpext

import (
	"log"
	"net/http"
)

func Listen(port int, handler http.Handler, fn ...func()) {
	for _, f := range fn {
		f()
	}
	log.Fatalln(http.ListenAndServe(Port(port).Addr(), handler))
}
