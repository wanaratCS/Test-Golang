package main

import (
	"fmt"
	"time"
)

func main() {
	// The leap year 2016 had 366 days.
	t1 := Date(1995, 12, 3)
	t2 := Date(2019, 7, 20)
	days := t2.Sub(t1).Hours() / 24
	fmt.Println(days) // 366
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
