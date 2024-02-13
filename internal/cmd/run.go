package cmd

import (
	"gin-choes-server/internal/middleware"
	"gin-choes-server/internal/route/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(port string) {
	app := gin.Default()

	app.Use(middleware.RateLimiter())

	app.StaticFS("static", http.Dir("static"))

	api := app.Group("api")
	client.AuthRoute(api)
	client.UserRoute(api)

	if err := app.Run(port); err != nil {
		panic("Server run error -> " + err.Error())
	}
}
