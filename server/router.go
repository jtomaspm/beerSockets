package server

import (
	"log"
	"strings"
)

type RouterMessage struct {
	Route   string
	Payload []byte
}

func NewRouterMessage(rawPayload []byte) RouterMessage {
	msg := string(rawPayload)
	lines := strings.Split(msg, "\n")
	var p []byte
	if len(lines) == 1 {
		p = []byte{}
	} else {
		p = []byte(strings.Join(lines[1:], "\n"))
	}
	return RouterMessage{
		Route:   strings.Trim(lines[0], " \n\t\r"),
		Payload: p,
	}
}

type Router struct {
	routes map[string]func([]byte) []byte
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func([]byte) []byte),
	}
}

func (self *Router) AddRoute(route string, handler func([]byte) []byte) {
	self.routes[route] = handler
}

func (self *Router) HandleRoute(message RouterMessage) []byte {
	log.Println("avalible routes:", self.routes)
	log.Println("route:", message.Route)
	log.Println("payload:", string(message.Payload))
	handler := self.routes[message.Route]
	if handler == nil {
		return []byte("ERROR\nInvalid Route")
	}
	res := handler(message.Payload)
	if res == nil {
		return []byte("ERROR\nInternal Server Error")
	}
	return res
}
