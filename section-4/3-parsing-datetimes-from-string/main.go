package main

import (
	"time"
	"fmt"
)

func main(){
	str := "2018-03-12T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"

	t,err := time.Parse(layout, str)

	if err != nil{
		fmt.Println(err)
	}

	kelime := "neden calismak istiyorsun anlamadim"
	fmt.Println(kelime)

	fmt.Println("nedenlerin bu kadar calismasina yardimci olmak demek")
	fmt.Printf("hello varlik denilen %v", kelime)
	

	fmt.Println(t.String())
}
