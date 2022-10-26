package main

import (
	"builder/builder/log"
	"builder/builder/views/router"
	"net/http"
)

func main() {
	r := router.Load()
	addr := "localhost:8080"
	log.L.Infof("Running Server on %v\n", addr)
	http.ListenAndServe(addr, r)
}
