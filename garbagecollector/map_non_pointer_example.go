package main

import (
	"fmt"
	"runtime"
)

func main() {
	var N = 40000000
	myMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}
	fmt.Printf("antes")
	runtime.GC()
	fmt.Printf("despues")
	_ = myMap[0]
}
