package models

import (
	"encoding/json"	"net/http"
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
		jsonAuth, _ := json.Marshal(Auth{
			AccessToken: ,
			Refresh:,
			LifeTime:,
			Status:,
		})
		w.WriteHeader(http.OK)
		return jsonAuth
	}
	jsonMessage, _ = json.Marshal(models.Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "StatusUnauthorized",
	})
	w.WriteHeader(http.StatusUnauthorized)
	return jsonMessage
}
