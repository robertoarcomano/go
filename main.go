package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("DATA TYPES AND BASIC")
	const my_const = 123
	a, s, b, c := 2, "Hello", true, complex(1, 2)
	conv_a := strconv.Itoa(a)
	mul, add := ops(2, 3)
	fmt.Println("my_const:", my_const, "| a:", a, "| conv_a:", conv_a, "| s:", s, "| b:", b, "| c:", c, "| math.Pi:", math.Pi, "| ops(2,3):", mul, add)
	fmt.Println()

	fmt.Println("ARRAY AND SLICE")
	var arr = [...]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Print("arr:", arr, " | slice:", slice, " | ")
	slice = append(slice, 6)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("slice:", slice, " | len(slice):", len(slice), "| cap(slice):", cap(slice))
	fmt.Println()

	fmt.Println("MAP")
	m := make(map[string]int)
	m["John"] = 1
	m["Sam"] = 2
	for k, v := range m {
		fmt.Print(k, ":", v, " | ")
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("FUNCTIONS / STRUCTS / POINTERS")
	var u user
	u.setName("John")
	fmt.Println("u.setName(\"John\") | u.getName():", u.getName(), "| u:", u, "| ")
	var u1 *user
	u1 = &u
	u1.setName("Sam")
	fmt.Println("u1 = &u", "| u1.setName(\"Sam\") | u1.getName():", u1.getName(), "| u.getName():", u.getName())
	fmt.Println()

	fmt.Println("GOROUTINES")
	n := 100
	fmt.Println("runtime.NumCPU():", runtime.NumCPU())
	fmt.Println("runtime.GOMAXPROCS(0):", runtime.GOMAXPROCS(0))
	fmt.Print("n:", n, "| Threads: ")
	stop := make(chan int, n)
	for i := 0; i < n; i++ {
		go show(i, stop)
	}
	num_go_routines := runtime.NumGoroutine()
	for i := 0; i < n; i++ {
		<-stop
	}
	fmt.Println("\nruntime.NumGoroutine():", num_go_routines)
}
