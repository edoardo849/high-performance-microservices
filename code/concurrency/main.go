package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func handler(c chan int, ch chan int) {

	time.Sleep(50 * time.Millisecond)
	val := <-c
	ch <- val
	fmt.Println("Handler:", val)
}

func worker(wch chan int, results chan int) {
	for {
		time.Sleep(10 * time.Millisecond)
		data := <-wch
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		<-results
	}
}

func pool(ch chan int, n int) {
	wch := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		go worker(wch, results)
		time.Sleep(1 * time.Millisecond)
	}
	go parse(results)
	for {
		val := <-ch
		wch <- val
	}
}

func server(l chan int, ch chan int) {
	for {
		time.Sleep(10 * time.Millisecond)
		p := <-l
		fmt.Println("Server:", p)

		j := make(chan int)
		go handler(j, ch)
		j <- p
	}
}

func main() {
	trace.Start(os.Stderr)

	l := make(chan int)
	ch := make(chan int)
	go pool(ch, 36)
	go server(l, ch)
	time.Sleep(2 * time.Second)

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		l <- i
	}

	time.Sleep(2 * time.Second)
	trace.Stop()
}
