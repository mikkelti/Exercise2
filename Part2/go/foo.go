// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

// Control signals, GetNumber is set to 0, Exit to 1
const (
	GetNumber = iota // See https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3.
	Exit
)

// Q: What is the difference between addNumber and e.g. conrol as arguments?
func numberServer(addNumber <-chan int, control <-chan int, number chan<- int) {
	var i = 0
	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		// TODO: receive different messages and handle them correctly
		// You will at least need to update the number and handle control signals.
		case j := <-addNumber:
			i += j
		case j := <-control:
			if j == GetNumber {
				number <- i
			} else {
				return
			}
		}
	}
}

func incrementing(addNumber chan<- int, finished chan<- bool) {
	for j := 0; j < 1000001; j++ {
		addNumber <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementing(addNumber chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		addNumber <- -1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	control := make(chan int)
	addNumber := make(chan int)
	finished := make(chan bool)
	number := make(chan int)

	// TODO: Spawn the required goroutines
	go numberServer(addNumber, control, number)
	go incrementing(addNumber, finished)
	go decrementing(addNumber, finished)

	// TODO: block on finished from both "worker" goroutines
	for j := 0; j < 2; j++ {
		Println(<-finished)
	}

	control <- GetNumber
	Println("The magic number is:", <-number)
	control <- Exit
}
