package main

import "fmt"

func mul_add(a, b int) (c, d int) {
	c = a * b
	d = a + b
	return
}

func main() {
	const a = 22
	var mul int
	var add int
	mul, add = mul_add(2, 5)
	fmt.Printf("%v %T", mul, add)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
