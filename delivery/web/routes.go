package web

import "net/http"

func NewRouter(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/", serveHome)
	mux.HandleFunc("/ws", webSocket)
}
