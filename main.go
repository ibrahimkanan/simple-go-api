package main

import (
	"github.com/labstack/echo/v4"
	"myapp/cmd/api/handlers"
)

	


func main(){
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHanlder)
	e.GET("posts:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}