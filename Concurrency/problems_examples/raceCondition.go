package main

import "fmt"

/*
	In this program exist a race condition

	the program will can terminate with 3 possible cases:

		- A is executed first that B then not printed
		- B y C are executed first that A then print 0
		- B is executed before that A but A executed before C then print 1


*/

func main() {

	var data int = 0
	go func() {
		data++ //A
	}()

	if data == 0 { //B
		fmt.Printf("data is %v", data) //C
	}

}
