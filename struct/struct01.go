package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

func main() {
	student1 := Student{"길동", 10, 5, 99.5}
	fmt.Println(student1.Name)
}
