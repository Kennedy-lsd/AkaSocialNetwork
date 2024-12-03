package api

import "net/http"

func (a *App) statusCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
}
