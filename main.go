package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

var n int = 123

func main() {
	const my_const = 123
	a, b := 2, 3
	a = 3
	s := "Hello"
	q := true
	p := math.Pi
	c := complex(1, 2)
	fmt.Println(s+" "+strconv.Itoa(sum(a, b)), q, c, p, my_const)
	var u user
	u.setName("John")
	fmt.Println(u.getName())
	fmt.Println(u)

	var u1 *user
	u1 = &u
	u1.setName("Sam")
	fmt.Println(u1.getName())
	fmt.Println(u1)

	var u2 *user
	fmt.Println(u2)

	u3 := user{"Danny"}
	fmt.Println(u3)

	n := 100
	fmt.Println(n)
	status := make(chan int, n)
	for i := 0; i < n; i++ {
		go show(i, status)
	}
	fmt.Println("runtime.NumCPU():", runtime.NumCPU())
	fmt.Println("runtime.NumGoroutine():", runtime.NumGoroutine())
	fmt.Println("runtime.GOMAXPROCS(0):", runtime.GOMAXPROCS(0))
	for i := 0; i < n; i++ {
		a := <-status
		fmt.Print(a)
		fmt.Print(" ")
	}
	fmt.Println()
}
