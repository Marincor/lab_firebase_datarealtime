package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/marincor/lab_firebase_datarealtime/handlers"
)

func main() {
	fmt.Println("init server...")
	e := echo.New()
	e.GET("/teams", handlers.ListAll)
	e.GET("/teams/:id", handlers.SearchOne)
	e.POST("/team", handlers.CreateOne)
	e.Logger.Fatal(e.Start(":2000"))
}
