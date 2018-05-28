package main

import (
	"time"
	"fmt"
)

func main(){
	current := time.Now()
	aprilDate := current.AddDate(1,1,0)

	fmt.Println(current.String())
	fmt.Println(aprilDate.String())

	// aprilDate.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))

	oneLessYears := aprilDate.AddDate(-1,0,0)
	fmt.Println(oneLessYears.String())

	tenMoreMinutes := aprilDate.Add(10 * time.Hour)
	fmt.Println(tenMoreMinutes)
}
