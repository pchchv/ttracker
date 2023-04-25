package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pchchv/env"
	"github.com/pchchv/golog"
)

type Day struct {
	Month   string
	Date    int16
	Weekend bool
	Holiday bool
	Workday bool
}

type Calendar struct {
	Year int16
	Days []Day
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

func newCalendar(year int) (c Calendar) {
	c.Year = int16(year)
	for i := 1; i <= 12; i++ {
		m := newMonth(i)
		c.Months = append(c.Months, m)
	}
	return
}

func newWorker(name string, dept string, year Calendar) (w Worker) {
	w.Name = name
	w.Department = dept
	w.PersonalCalendar = year

	return
}

func createPerson(jsonMap map[string]interface{}) Worker {
	name := fmt.Sprint(jsonMap["Name"])
	department := fmt.Sprint(jsonMap["Department"])
	year, err := strconv.Atoi(fmt.Sprint(jsonMap["Year"]))
	if err != nil {
		golog.Panic(err.Error())
	}

	// TODO: Implement the receipt of the calendar.
	// Creating a new one is a bad idea.
	calendar := newCalendar(year)

	return newWorker(name, department, calendar)
}

func main() {
	server()
}
