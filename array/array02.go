package main

import "fmt"

func main() {
	var a [3]int = [3]int{3, 6, 9}

	for i, v := range a {
		fmt.Println(i, v)
	}
}
