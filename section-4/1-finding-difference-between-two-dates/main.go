package main

import (
	"time"
	"fmt"
)

func main(){
	first := time.Date(2010, 1,1,0,0,0,0,time.UTC)
	second := time.Date(2011, 1,1,0,0,0,0,time.UTC)

	difference := second.Sub(first)
	fmt.Printf("Difference %v", difference)
}
