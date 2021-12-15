package app

import "net/http"

type app struct {
	server *http.Server
}

func NewApp(cfg *config) *app {
	return &app{
		server: &http.Server{
			Addr: cfg.Port,
		},
	}
}

func (a *app) Run() error {
	return a.server.ListenAndServe()
}
