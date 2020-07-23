package server

import (
	"net/http"
)

func (s *Server) NewAuthController() {
	s.Router.HandleFunc("/login", s.NewHandleLogin())
}

func (s *Server) NewHandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
