package main

import (
	"log"
	"net/http"

	"github.com/Shalqarov/real-time-comments/delivery/web"
)

func main() {
	mux := http.NewServeMux()
	web.NewRouter(mux)
	log.Println(http.ListenAndServe(":8080", mux))
}
