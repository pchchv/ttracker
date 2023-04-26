package main

import (
	"fmt"
	"os"

	"github.com/pchchv/env"
	"github.com/pchchv/golog"
	"github.com/rickar/cal/v2"
)

var testURL string

func init() {
	// Load values from .env into the system
	if err := env.Load(); err != nil {
		golog.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		golog.Panic("Value %v does not exist", v)
	}
	return value
}

func newCalendar(year int) *cal.BusinessCalendar {
	c := cal.NewBusinessCalendar()
	c.Name = fmt.Sprintf("%v year", year)
	c.Description = "Default calendar"

	return c
}

func main() {
	server()
}
