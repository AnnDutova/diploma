package server

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type Router struct {
	Handler *Handler
	Router  *mux.Router
}

func NewRouter(h *Handler) (*Router, error) {
	return &Router{
		Handler: h,
		Router:  mux.NewRouter(),
	}, nil
}

func (r *Router) configurateRouter(addr string) error {
	log.Println("Configurate Router")

	r.Router.HandleFunc("/", r.Handler.backupHandler)
	r.Router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	return http.ListenAndServe(addr, r.Router)
}
