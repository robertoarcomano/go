package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const my_const = 123
	a, s, b, c := 2, "Hello", true, complex(1, 2)
	conv_a := strconv.Itoa(a)
	fmt.Println("my_const:", my_const, "| a:", a, "| conv_a:", conv_a, "| s:", s, "| b:", b, "| c:", c, "| pi:", math.Pi)

	var arr = [...]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Print("arr:", arr, " | slice:", slice, " | ")
	slice = append(slice, 6)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("slice:", slice, " | len(slice):", len(slice), "| cap(slice):", cap(slice))

	fmt.Print("map: ")
	m := make(map[string]int)
	m["John"] = 1
	m["Sam"] = 2
	for k, v := range m {
		fmt.Print(k, ":", v, " | ")
	}
	fmt.Println()

	// var u user
	// u.setName("John")
	// fmt.Println(u.getName())
	// fmt.Println(u)

	// var u1 *user
	// u1 = &u
	// u1.setName("Sam")
	// fmt.Println(u1.getName())
	// fmt.Println(u1)

	// var u2 *user
	// fmt.Println(u2)

	// u3 := user{"Danny"}
	// fmt.Println(u3)

	// n := 100
	// fmt.Println(n)
	// status := make(chan int, n)
	// for i := 0; i < n; i++ {
	// 	go show(i, status)
	// }
	// fmt.Println("runtime.NumCPU():", runtime.NumCPU())
	// fmt.Println("runtime.NumGoroutine():", runtime.NumGoroutine())
	// fmt.Println("runtime.GOMAXPROCS(0):", runtime.GOMAXPROCS(0))
	// for i := 0; i < n; i++ {
	// 	a := <-status
	// 	fmt.Print(a)
	// 	fmt.Print(" ")
	// }
	// fmt.Println()
}
