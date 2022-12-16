package main

import (
	"fmt"
)

func main() {
	var strslice []string
	stu := Student{18, "male", "cocoliu"}
	strslice = []string{stu.gender}
	fmt.Println(strslice)
}

type Person interface {
	say()
}
type Student struct {
	age    int
	gender string
	name   string
}

func (s *Student) say() {
	fmt.Printf("my name is %v", s.name)
	// cache.Reflector()
	// cache.DeltaFIFO{}
	// cache.SharedIndexInformer()
}
