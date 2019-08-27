package main

import (
	f "fmt"
	"time"
)

func dumbAdder() int64 {
	var x int64
	var i int64
	for i = 0; i < 2147683647*20; i++ {
		x += i
	}
	return x
}
func adder(c chan int64, r int64) {
	var x int64
	var i int64
	for i = 2147683647 * (r); i < 2147683647*(r+1); i++ {
		x += i
	}
	c <- x
}
func main() {
	t1 := time.Now()
	tot := dumbAdder()
	t2 := time.Now()
	f.Println("The total is: ", tot)
	f.Println("Time taken by 1 thread to calculate sum of numbers from 0 to 42953672940: ", t2.Sub(t1))

	t1 = time.Now()

	var channels [20]chan int64
	for i := range channels {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go adder(channels[i], int64(i))
	}

	var total int64

	for i := range channels {
		temp := <-channels[i]
		total += temp
	}

	t2 = time.Now()

	f.Println("The total is: ", total)
	f.Println("Time taken by 20 go routines to calculate sum of numbers from 0 to 42953672940: ", t2.Sub(t1))

	if tot == total {
		f.Println("Total from concurrent goroutines same as total from one thread, no sync errors!")
	}
}
