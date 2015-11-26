package doitserver

import (
	"net/http"

	"github.com/gorilla/pat"
)

type Server struct {
	Consumers *Consumers
	Mux       http.Handler

	key string
}

func NewServer(key string) *Server {
	cc := NewConsumers()
	p := pat.New()

	a := NewAuth(key)
	ac := NewAuthCallback(cc, key)
	t := NewTokenGenerator(key)
	am := NewAuthMonitor(cc, key)

	p.Add("GET", "/auth/{provider}/callback", ac)
	p.Add("GET", "/auth/{provider}", a)
	p.Add("GET", "/token", t)
	p.Add("GET", "/status", am)

	return &Server{
		Consumers: cc,
		Mux:       p,
		key:       key,
	}
}
