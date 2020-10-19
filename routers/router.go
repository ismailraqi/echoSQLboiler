package routers

import (
	"fmt"

	"github.com/ismailraqi/Golang-sqlboiler/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// declare & inisialize new instance of echo
var e = echo.New()
var conf Config

// function that read configuration from Environement before app start
func init() {
	cfg := Configuration()
	conf = cfg
}

//StartRouters is a func that's give you access to all routes
func StartRouters() {
	//e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/pilots", handlers.GetAllPilots)
	e.GET("/pilot/:id", handlers.GetOnePilots)
	e.POST("/pilot", handlers.CreatePilot)
	e.Logger.Print(fmt.Sprintf("Listening on prot: %d\n", conf.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%d", conf.Port)))
}
