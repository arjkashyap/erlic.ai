package handlers

import "net/http"

type HealthCheckHandler struct {
}

func (hc *HealthCheckHandler) HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Ok\n"))
}
