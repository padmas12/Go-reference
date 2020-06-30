package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Counting for a while ... ")
	count("Counting some more ...")
}

func count(howlong string) {
	for i := 1; true; i++ {
		fmt.Println(howlong)
		time.Sleep(time.Millisecond * 1010)
	}
}
