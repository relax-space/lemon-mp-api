package main

import (
	"lemon-mp-api/auth"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	RegApi(e)
	e.Start(":5000")
}

func RegApi(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("v1")
	v1.GET("/mpDoor", auth.OpenIdGate)
	v1.GET("/openId/:appId", auth.OpenIdRequest)
}
