package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wlcm")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
}