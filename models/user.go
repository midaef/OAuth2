package models

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type User struct {
	GUID string `json:"guid"`
}

type Auth struct {
	AccessToken string `json:"access"`
	Refresh     string `json:"refresh"`
	LifeTime    string `json:"lifetime"`
	Status      string `json:"status"`
}

type Refresh struct {
	Refresh string `json:"refresh"`
}

func NewAuth(user *User, w http.ResponseWriter) []byte {
	if len(user.GUID) == 32 {
		h := sha512.New()
		io.WriteString(h, user.GUID)
		jsonAuth, _ := json.Marshal(Auth{
			AccessToken: base64.URLEncoding.EncodeToString(h.Sum(nil)),
			Refresh:     base64.StdEncoding.EncodeToString([]byte(user.GUID)),
			LifeTime:    "30",
			Status:      "OK",
		})
		w.WriteHeader(http.StatusOK)
		return jsonAuth
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "StatusUnauthorized",
	})
	w.WriteHeader(http.StatusUnauthorized)
	return jsonMessage
}
