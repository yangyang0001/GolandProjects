package main

import (
	"fmt"
	"time"
)

func main() {
	minetime := time.Now()
	minetime = minetime.AddDate(0, 0, 1)
	fmt.Printf("minetime = %v \n", time.Unix(minetime.Unix(), 0).Format("2006-01-02 15:04:05"))
}
