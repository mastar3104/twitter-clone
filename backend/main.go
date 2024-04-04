package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	fanout "github.com/mastar3104/twitter-clone/fanout/application"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	tweet "github.com/mastar3104/twitter-clone/tweet/application"
	user "github.com/mastar3104/twitter-clone/user/application"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	_ = godotenv.Load("./.env")

	timeZone, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = timeZone

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := adapter.GetDatabaseHandler()
	defer handler.Close()

	server := createServer(handler)

	server.Logger.Fatal(server.Start(":" + port))
}

func createServer(handler *adapter.MySQLHandler) *echo.Echo {
	server := echo.New()
	server.HideBanner = true

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_END_HOST")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	}))

	tweet.Route(server, handler)
	user.Route(server, handler)
	fanout.Route(server, handler)

	return server
}
