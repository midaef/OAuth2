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

	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
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

func NewAuth(user *User, w http.ResponseWriter, s *mongo.Session, client *mongo.Client, ctx context.Context) []byte {
	if len(user.GUID) == 32 {
		h := sha512.New()
		token := string(securecookie.GenerateRandomKey(16))
		io.WriteString(h, user.GUID+token)
		var auth Auth
		auth = Auth{
			GUID:        user.GUID,
			AccessToken: base64.URLEncoding.EncodeToString(h.Sum(nil)),
			Refresh:     base64.StdEncoding.EncodeToString([]byte(user.GUID + token)),
			LifeTime:    "30",
		}
		jsonAuth, _ := json.Marshal(auth)
		c := client.Database("DataBaseTest").Collection("CollectionTest")
		err := mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
			sessionContext.StartTransaction()
			var guid User
			err := c.FindOne(ctx, bson.M{"GUID": user.GUID}).Decode(&guid)
			checkError(err, sessionContext)
			if guid.GUID == user.GUID {
				_, err := c.InsertOne(ctx, auth)
				checkError(err, sessionContext)
				sessionContext.CommitTransaction(sessionContext)
				return nil
			}
			return errors.Errorf("GUID is not registered")
		})
		if err != nil {
			jsonMessage, _ := json.Marshal(Message{
				DateTime: time.Now().Format("01-02-2006 15:04:05"),
				Status:   "Unauthorized",
			})
			w.WriteHeader(http.StatusUnauthorized)
			return jsonMessage
		}
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

func NewReg(user *User, w http.ResponseWriter, s *mongo.Session, client *mongo.Client, ctx context.Context) []byte {
	if len(user.GUID) == 32 {
		jsonMessage, _ := json.Marshal(Message{
			DateTime: time.Now().Format("01-02-2006 15:04:05"),
			Status:   "Registered",
		})
		c := client.Database("DataBaseTest").Collection("CollectionTest")
		err := mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
			sessionContext.StartTransaction()
			var guid User
			err := c.FindOne(ctx, bson.M{"GUID": user.GUID}).Decode(&guid)
			checkError(err, sessionContext)
			if guid.GUID != user.GUID {
				_, err := c.InsertOne(ctx, bson.M{"GUID": user.GUID})
				checkError(err, sessionContext)
				sessionContext.CommitTransaction(sessionContext)
				return nil
			}
			return errors.Errorf("GUID REGISTERED")
		})
		if err != nil {
			jsonMessage, _ := json.Marshal(Message{
				DateTime: time.Now().Format("01-02-2006 15:04:05"),
				Status:   "GuidRegistered",
			})
			w.WriteHeader(http.StatusBadRequest)
			return jsonMessage
		}
		w.WriteHeader(http.StatusOK)
		return jsonMessage
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "NotRegistered",
	})
	w.WriteHeader(http.StatusBadRequest)
	return jsonMessage
}

func NewRefresh(refresh *Refresh, w http.ResponseWriter, s *mongo.Session, client *mongo.Client, ctx context.Context) []byte {
	c := client.Database("DataBaseTest").Collection("CollectionTest")
	err := mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
		sessionContext.StartTransaction()
		var user Auth
		err := c.FindOne(ctx, bson.M{"refresh": refresh.Refresh}).Decode(&user)
		checkError(err, sessionContext)
		h := sha512.New()
		token := string(securecookie.GenerateRandomKey(16))
		io.WriteString(h, user.GUID+token)
		user = Auth{
			GUID:        user.GUID,
			AccessToken: base64.URLEncoding.EncodeToString(h.Sum(nil)),
			Refresh:     base64.StdEncoding.EncodeToString([]byte(user.GUID + token)),
			LifeTime:    "30",
		}
		atualizacao := bson.D{{Key: "$set", Value: user}}
		_, e := c.UpdateOne(ctx, bson.M{"refresh": refresh.Refresh}, atualizacao)
		checkError(e, sessionContext)
		sessionContext.CommitTransaction(sessionContext)
		return err
	})
	if err != nil {
		jsonMessage, _ := json.Marshal(Message{
			DateTime: time.Now().Format("01-02-2006 15:04:05"),
			Status:   "RefreshTokenNotChanged",
		})
		w.WriteHeader(http.StatusBadRequest)
		return jsonMessage
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "RefreshTokenChanged",
	})
	w.WriteHeader(http.StatusOK)
	return jsonMessage
}

func NewDelRefreshToken(refresh *Refresh, w http.ResponseWriter, s *mongo.Session, client *mongo.Client, ctx context.Context) []byte {
	c := client.Database("DataBaseTest").Collection("CollectionTest")
	err := mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
		sessionContext.StartTransaction()
		var user Auth
		err := c.FindOne(ctx, bson.M{"refresh": refresh.Refresh}).Decode(&user)
		checkError(err, sessionContext)
		h := sha512.New()
		token := string(securecookie.GenerateRandomKey(16))
		io.WriteString(h, user.GUID+token)
		user = Auth{
			GUID:        user.GUID,
			AccessToken: user.AccessToken,
			Refresh:     "",
			LifeTime:    user.LifeTime,
		}
		atualizacao := bson.D{{Key: "$set", Value: user}}
		_, e := c.UpdateOne(ctx, bson.M{"refresh": refresh.Refresh}, atualizacao)
		checkError(e, sessionContext)
		sessionContext.CommitTransaction(sessionContext)
		return err
	})
	if err != nil {
		jsonMessage, _ := json.Marshal(Message{
			DateTime: time.Now().Format("01-02-2006 15:04:05"),
			Status:   "RefreshTokenNotDeleted",
		})
		w.WriteHeader(http.StatusBadRequest)
		return jsonMessage
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "RefreshTokenDeleted",
	})
	w.WriteHeader(http.StatusOK)
	return jsonMessage
}

func NewDelRefreshTokens(guid *User, w http.ResponseWriter, s *mongo.Session, client *mongo.Client, ctx context.Context) []byte {
	c := client.Database("DataBaseTest").Collection("CollectionTest")
	err := mongo.WithSession(ctx, *s, func(sessionContext mongo.SessionContext) error {
		sessionContext.StartTransaction()
		var user Auth
		err := c.FindOne(ctx, bson.M{"guid": guid.GUID}).Decode(&user)
		checkError(err, sessionContext)
		h := sha512.New()
		token := string(securecookie.GenerateRandomKey(16))
		io.WriteString(h, user.GUID+token)
		user = Auth{
			GUID:        user.GUID,
			AccessToken: user.AccessToken,
			Refresh:     "",
			LifeTime:    user.LifeTime,
		}
		atualizacao := bson.D{{Key: "$set", Value: user}}
		_, e := c.UpdateMany(ctx, bson.M{"guid": guid.GUID}, atualizacao)
		checkError(e, sessionContext)
		sessionContext.CommitTransaction(sessionContext)
		return err
	})
	if err != nil {
		jsonMessage, _ := json.Marshal(Message{
			DateTime: time.Now().Format("01-02-2006 15:04:05"),
			Status:   "RefreshTokensNotDeleted",
		})
		w.WriteHeader(http.StatusBadRequest)
		return jsonMessage
	}
	jsonMessage, _ := json.Marshal(Message{
		DateTime: time.Now().Format("01-02-2006 15:04:05"),
		Status:   "RefreshTokensDeleted",
	})
	w.WriteHeader(http.StatusOK)
	return jsonMessage
}

func checkError(err error, sessionContext mongo.SessionContext) error {
	if err != nil {
		sessionContext.AbortTransaction(sessionContext)
		log.Println(err)
		return err
	}
	return nil
}
