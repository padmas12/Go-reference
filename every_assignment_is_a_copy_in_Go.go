package main

import "fmt"

func increment(x *int) {
	*x++
}
func main() {
	v := 1
	increment(&v)
	fmt.Println("Main's local copy of v is", v)
}
