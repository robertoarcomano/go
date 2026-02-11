package main

import (
	"fmt"
)

func show(i int, stop chan int) {
	fmt.Print(i, " ")
	stop <- 1
}
