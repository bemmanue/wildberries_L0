package service

import (
	"github.com/bemmanue/wildberries_L0/internal/store"
	"io"
	"net/http"
)

type server struct {
	store  store.Store
	router *http.ServeMux
}

func newServer(store store.Store) *server {
	s := &server{
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleHello())
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
