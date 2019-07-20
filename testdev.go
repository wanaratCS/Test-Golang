package main

import (
	"fmt"
	"app"
)

func main(){
	t1 := Date(1995, 12, 3)
	t2 := Date(2019, 7, 20)
	days := t2.Sub(t1).Hours() / 24
	ExpectedDay := 8507
	fmt.Println(ExpectedDay)
	app.Date()
}