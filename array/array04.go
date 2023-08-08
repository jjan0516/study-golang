package main

import "fmt"

func main() {
	var a [2][5]int = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	b := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	// 2중 for문
	for _, arr := range a {
		for _, v := range arr {
			fmt.Println(v)
		}
	}
	for _, arr := range b {
		for _, v := range arr {
			fmt.Println(v)
		}
	}
}
