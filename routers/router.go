package routers

import (
	"fmt"
	"log"
	"net/http"

	jwtModelsClaims "github.com/ismailraqi/echoSQLboiler/jwtmodelsclaims"

	"github.com/ismailraqi/echoSQLboiler/handlers"
	_ "github.com/ismailraqi/echoSQLboiler/statik"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rakyll/statik/fs"
)

// declare & inisialize new instance of echo
var e = echo.New()
var conf Config

// function that read configuration from Environement before StartRouters() func
func init() {
	cfg := Configuration()
	conf = cfg
}

//StartRouters is a func that's give you access to all routes
func StartRouters() {
	e.Debug = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/pilots", handlers.GetAllPilots)
	e.GET("/pilot/:id", handlers.GetOnePilots)
	e.POST("/pilot", handlers.CreatePilot)
	e.DELETE("/pilot/:id", handlers.DeletePilot)
	e.PUT("/pilot/:id", handlers.UpdatePilot)
	e.POST("/register", handlers.UserRegister)
	e.POST("/login", handlers.UserLogin)
	r := e.Group("/restricted")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtModelsClaims.JwtCustomClaims{},
		SigningKey: []byte("secretToken"),
	}
	r.Use(middleware.JWTWithConfig(config))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	h := http.FileServer(statikFS)
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))
	// // Serve the contents over HTTP.
	// http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(statikFS)))
	// http.ListenAndServe(":8080", nil)

	//e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.Logger.Print(fmt.Sprintf("Listening on prot: %d\n", conf.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%d", conf.Port)))
}
