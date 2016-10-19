package main

import (
	"os"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	const j = 36
	done := make(chan int)

	for i := 0; i < j; i++ {
		// this sleep is added to more or less preserve an order
		// of the goroutines id on creation
		time.Sleep(20 * time.Millisecond)
		go th(done, i)
		time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)

	for i := 0; i < j; i++ {
		<-done
	}

	trace.Stop()
}

func th(done chan int, i int) {
	done <- i
	time.Sleep(10 * time.Millisecond)
}
