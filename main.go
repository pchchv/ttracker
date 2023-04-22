package main

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
}
