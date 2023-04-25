package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pchchv/golog"
)

// Checks that the server is up and running.
func pingHandler(c echo.Context) error {
	message := "Time tracking service. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

func createCalendarHandler(c echo.Context) error {
	var jsonMap map[string]interface{}
	if err := c.Bind(&jsonMap); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Some input error: %e", err))
	}

	calendar, err := newCalendar(jsonMap)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Some input error: %e", err))
	}

	return c.JSON(http.StatusOK, calendar)
}

func createPersonHandler(c echo.Context) error {
	var jsonMap map[string]interface{}
	if err := c.Bind(&jsonMap); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Some input error: %e", err))
	}

	person := createPerson(jsonMap)
	return c.JSON(http.StatusOK, person)
}

// The declaration of all routes comes from it.
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.POST("/newyear", createCalendarHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	golog.Fatal(e.Start(":" + getEnvValue("PORT")).Error())
}
