package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pchchv/golog"
)

// Checks that the server is up and running.
func pingHandler(c echo.Context) error {
	message := "Time tracking service. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

// Creates new business calendar.
func createCalendarHandler(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Error when parsing the year")
	}

	calendar := newCalendar(year)

	return c.JSON(http.StatusOK, calendar)
}

// The declaration of all routes comes from it.
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	golog.Fatal(e.Start(":" + getEnvValue("PORT")).Error())
}
