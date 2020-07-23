package server

import (
	"encoding/json"
	"net/http"
	"packages/models"
)

func (s *Server) NewAuthController() {
	s.Router.HandleFunc("/login", s.NewHandleLogin())
}

func (s *Server) NewHandleLogin() http.HandlerFunc {
	var user models.User
	return func(w http.ResponseWriter, r *http.Request) {
		body := NewRequestReader(r)
		json.Unmarshal(body, &user)
		jsonByte := models.NewAuth(&user, w)
		NewResponseWriter(jsonByte, w)
	}
}
