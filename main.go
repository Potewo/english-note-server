package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Note struct {
	English string `json:"english"`
	Japanese string `json:"japanese"`
	Description string `json:"description"`
	Examples string `json:"examples"`
	Similar string `json:"similar"`
	Tags []string `json:"tags"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/add", func(c echo.Context) error {
		u := new([]Note)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
