package main

import "fmt"

func main() {
	m := 0
	if true {
		m := 1
		m++
	}
	fmt.Println(m)
}
