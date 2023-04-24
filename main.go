package main

import (
	"os"

	"github.com/pchchv/env"
	"github.com/pchchv/golog"
)

type Month struct {
	Title        string
	NumberOfDays int16
	Days         []Day
}

type Day struct {
	title   string
	date    int16
	weekend bool
	holiday bool
	workday bool
}

type Calendar struct {
	Year   int16
	Months []Month
}

type Worker struct {
	Name             string
	Department       string
	PersonalCalendar Calendar
}

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

func newMonth(num int) (m Month) {
	// TODO
	return
}

func newCalendar(year int) (c Calendar) {
	c.Year = int16(year)
	for i := 1; i <= 12; i++ {
		m := newMonth(i)
		c.Months = append(c.Months, m)
	}
	return
}

func main() {
	server()
}
