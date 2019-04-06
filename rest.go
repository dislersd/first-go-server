package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	data map[string]string
)

func main() {
	flag.Parse()
	data = map[string]string{}
	r := httprouter.New()
	r.GET("/entry/:key", show)
	r.GET("/list", show)
	r.PUT("/entry/:key/:value", update)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
