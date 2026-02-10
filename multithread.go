package main

import (
	"time"
)

func show(i int, status chan int) {
	start := time.Now()
	x := 0
	for time.Since(start) < 10*time.Second {
		// calcolo inutile per tenere occupata la CPU
		x += 1
		x *= 2
		x /= 2
	}
	status <- i
}
