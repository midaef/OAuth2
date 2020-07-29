package models

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	GUID string `json:"guid"`
}

type Auth struct {
	GUID        string `json:"guid"`
	AccessToken string `json:"access"`
	Refresh     string `json:"refresh"`
	LifeTime    string `json:"lifetime"`
}

type Refresh struct {
	Refresh string `json:"refresh"`
}

func NewAuth(user *User, w http.ResponseWriter, s *mongo.Session, c *mongo.Client, ctx context.Context) []byte {
	if len(user.GUID) == 32 {
		h := sha512.New()
		io.WriteString(h, user.GUID)
		var auth Auth
		auth = Auth{
			GUID:        user.GUID,
			AccessToken: base64.URLEncoding.EncodeToString(h.Sum(nil)),
			Refresh:     base64.StdEncoding.EncodeToString([]byte(user.GUID)),
			LifeTime:    "30",
		}
		jsonAuth, _ := json.Marshal(auth)
		c := c.Database("DataBaseTest").Collection("CollectionTest")
		mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
			sessionContext.StartTransaction()
			_, err := c.InsertOne(ctx, bson.M{"Session": auth})
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)
				log.Println(err)
				return err
			}
			sessionContext.CommitTransaction(sessionContext)
			return nil
		})
		w.WriteHeader(http.StatusOK)
		return jsonAuth
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "Unauthorized",
	})
	w.WriteHeader(http.StatusUnauthorized)
	return jsonMessage
}

func NewReg(user *User, w http.ResponseWriter, s *mongo.Session, c *mongo.Client, ctx context.Context) []byte {
	if len(user.GUID) == 32 {
		jsonMessage, _ := json.Marshal(Message{
			DateTime: time.Now().Format("01-02-2006 15:04:05"),
			Status:   "Registered",
		})
		c := c.Database("DataBaseTest").Collection("CollectionTest")
		mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
			sessionContext.StartTransaction()
			_, err := c.InsertOne(ctx, bson.M{"GUID": user.GUID})
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)
				log.Println(err)
				return err
			}
			sessionContext.CommitTransaction(sessionContext)
			return nil
		})
		w.WriteHeader(http.StatusOK)
		return jsonMessage
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "NotRegistered",
	})
	w.WriteHeader(http.StatusUnauthorized)
	return jsonMessage
}
