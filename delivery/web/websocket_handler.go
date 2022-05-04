package web

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type Hub struct {
	clients    map[*client]bool
	comments   chan []byte
	register   chan *client
	unregister chan *client
}

func NewHub() *Hub {
	return &Hub{
		comments:   make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	// TODO serve ws
}
