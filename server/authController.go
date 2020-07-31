package server

import (
	"encoding/json"
	"net/http"
	"packages/models"
)

func (s *Server) NewAuthController() {
	s.Router.HandleFunc("/reg", s.NewHandleReg())
	s.Router.HandleFunc("/login", s.NewHandleLogin())
	s.Router.HandleFunc("/refresh", s.NewHandleRefresh())
	s.Router.HandleFunc("/delRefreshToken", s.NewHandleDelRefreshToken())
	s.Router.HandleFunc("/delRefreshTokens", s.NewHandleDelRefreshTokens())
}

func (s *Server) NewHandleRefresh() http.HandlerFunc {
	var refresh models.Refresh
	session := s.ConnectionDB.Session
	client := s.ConnectionDB.Client
	ctx := s.ConnectionDB.Context
	return func(w http.ResponseWriter, r *http.Request) {
		body := NewRequestReader(r)
		json.Unmarshal(body, &refresh)
		jsonByte := models.NewRefresh(&refresh, w, session, client, ctx)
		NewResponseWriter(jsonByte, w)
	}
}

func (s *Server) NewHandleDelRefreshToken() http.HandlerFunc {
	var refresh models.Refresh
	session := s.ConnectionDB.Session
	client := s.ConnectionDB.Client
	ctx := s.ConnectionDB.Context
	return func(w http.ResponseWriter, r *http.Request) {
		body := NewRequestReader(r)
		json.Unmarshal(body, &refresh)
		jsonByte := models.NewDelRefreshToken(&refresh, w, session, client, ctx)
		NewResponseWriter(jsonByte, w)
	}
}

func (s *Server) NewHandleDelRefreshTokens() http.HandlerFunc {
	var user models.User
	session := s.ConnectionDB.Session
	client := s.ConnectionDB.Client
	ctx := s.ConnectionDB.Context
	return func(w http.ResponseWriter, r *http.Request) {
		body := NewRequestReader(r)
		json.Unmarshal(body, &user)
		jsonByte := models.NewDelRefreshTokens(&user, w, session, client, ctx)
		NewResponseWriter(jsonByte, w)
	}
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
