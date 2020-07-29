package server

import (
	"encoding/json"
	"net/http"
	"packages/models"
)

func (s *Server) NewAuthController() {
	s.Router.HandleFunc("/reg", s.NewHandleReg())
	s.Router.HandleFunc("/login", s.NewHandleLogin())
}

func (s *Server) NewHandleLogin() http.HandlerFunc {
	var user models.User
	session := s.ConnectionDB.Session
	client := s.ConnectionDB.Client
	ctx := s.ConnectionDB.Context
	return func(w http.ResponseWriter, r *http.Request) {
		body := NewRequestReader(r)
		json.Unmarshal(body, &user)
		jsonByte := models.NewAuth(&user, w, session, client, ctx)
		NewResponseWriter(jsonByte, w)
	}
}

func (s *Server) NewHandleReg() http.HandlerFunc {
	var user models.User
	session := s.ConnectionDB.Session
	client := s.ConnectionDB.Client
	ctx := s.ConnectionDB.Context
	return func(w http.ResponseWriter, r *http.Request) {
		body := NewRequestReader(r)
		json.Unmarshal(body, &user)
		jsonByte := models.NewReg(&user, w, session, client, ctx)
		NewResponseWriter(jsonByte, w)
	}
}
