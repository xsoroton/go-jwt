package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"go-jwt/models"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
)

const numberOfEvents = 5

func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware := GinJWTMiddleware()

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/events", GetUserEvents)
	}

	http.ListenAndServe(":"+port, r)
}

// GetUserEvents ...
func GetUserEvents(c *gin.Context) {
	// swagger:route GET /auth/events get events
	//
	// Get Events by JWT location value
	// ---
	// produces:
	// - application/json
	//
	// schemes: Events
	//
	// responses:
	//  200: models.Events
	//
	var location string
	claims := jwt.ExtractClaims(c)
	logrus.Infof("Extracted data from token: %#v", claims)
	//if reflect.TypeOf(claims["location"]) == reflect.Interface {
	//	c.JSON(200, GetEvents(``))
	//}
	if len(claims["location"].(string)) > 0 {
		logrus.Infof("Get events for location: %v", claims["location"])
		location = claims["location"].(string)
	}
	c.JSON(200, GetEvents(location))
}

// GetEvents return array of Events
func GetEvents(location string) models.Events {
	return models.Events{GenerateFakeEvents(numberOfEvents, location)}
}

// GinJWTMiddleware handle JWT for GIN
func GinJWTMiddleware() *jwt.GinJWTMiddleware{
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret"),
		Timeout:    time.Hour * 24,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string, c *gin.Context) (interface{}, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return models.User{
					Name:     fake.FullName(),
					Location: fake.City(),
					Sub:      fake.CharactersN(10),
					Admin:    userId == "admin",
				}, true
			}
			return nil, false
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			if user.(string) == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		PayloadFunc: func(data interface{}) map[string]interface{} {
			logrus.Infof("Fake data %#v", data)
			user := data.(models.User)
			return map[string]interface{}{"name": user.Name, "location": user.Location, "sub": user.Sub, "admin": user.Admin}
		},
		TokenLookup: "header:Authorization",
		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}
}

// GenerateFakeEvents ...
func GenerateFakeEvents(n int, location string) (events []models.Event) {
	// if location is not set make make it random
	if len(location) == 0 {
		location = fake.City()
	}
	for i := 1; i <= n; i++ {
		events = append(events, models.Event{
			Title:          fake.Title(),
			Date:           time.Now().UTC().AddDate(0, 0, rand.Intn(100)).Format(time.RFC3339),
			ImageURL:       fake.CharactersN(10) + `.jpg`,
			AvailableSeats: rand.Intn(10000000),
			Location:       location,
		})
	}
	return
}
