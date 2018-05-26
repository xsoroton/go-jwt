package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"encoding/json"
	"go-jwt/models"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

func init() {
	gin.SetMode(gin.TestMode)
	logrus.SetLevel(logrus.WarnLevel)
}

func TestGetUserEvents(t *testing.T) {
	Convey("Test GetUserEvents", t, func() {
		router := gin.New()
		authMiddleware := GinJWTMiddleware()
		router.Use(authMiddleware.MiddlewareFunc())
		router.GET("/auth/events", GetUserEvents)

		token := generateToken()

		req, _ := http.NewRequest(http.MethodGet, "/auth/events", nil)
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("Content-Type:", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		var apiResponse models.Events
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			panic(err.Error())
		}
		Convey("Check correct number of events", func() {
			So(len(apiResponse.Events), ShouldEqual, numberOfEvents)
		})
		Convey("Check correct model", func() {
			So(apiResponse, ShouldHaveSameTypeAs, models.Events{})
		})
		Convey("Check that all events match Location from JWT", func() {
			jwtPayload, _ := extractClaims(token)
			for _, event := range apiResponse.Events {
				So(event.Location, ShouldEqual, jwtPayload["location"])
			}
		})
	})
}

func generateToken() string {
	router := gin.New()
	authMiddleware := GinJWTMiddleware()
	router.POST("/login", authMiddleware.LoginHandler)

	data := bytes.NewBuffer([]byte(`{"username": "admin", "password": "admin"}`))
	req, _ := http.NewRequest(http.MethodPost, "/login", data)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var i map[string]interface{}
	err = json.Unmarshal(body, &i)
	if err != nil {
		panic(err.Error())
	}
	return i["token"].(string)
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "secret"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
