package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", getSuccess)
	e.POST("/", echoRequest)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func echoRequest(context echo.Context) error {
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, string(body))
}

func getSuccess(context echo.Context) error {
	return context.String(http.StatusOK, "success")
}
