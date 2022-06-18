package client

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Request(query *url.Values) {
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Panicln(err)
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[ERROR]", err.Error())
	} else {
		b, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Println("[ERROR]", err.Error())
		}
		fmt.Println(string(b))
	}
}
