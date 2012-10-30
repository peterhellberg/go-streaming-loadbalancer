package main

import (
	"fmt"
	"net/http"

	"github.com/bmizerany/pat"
)

func startServer() {
	endpoint := "/loadbalancer.json"

	m := pat.New()
	m.Get(endpoint, http.HandlerFunc(Loadbalancer))
	m.Get("/", http.RedirectHandler(endpoint, http.StatusMovedPermanently))

	http.Handle("/", m)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
