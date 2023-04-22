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

type Calendar struct {
	Year   int16
	Months []Month
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

func main() {}
