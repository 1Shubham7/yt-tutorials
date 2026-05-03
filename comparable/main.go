package main

import (
	"fmt"
)

type ComparableStruct struct {
	Value int
	Name string
	A int
}

type NonComparableStruct struct {
	B int
	name string
	A []int 
}

func Contains[T comparable](slice []T, item T) bool {
	for _,v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GiveMeFirstItem[T any](slice []T) T {
	return slice[0]
}

func main() {
	fmt.Println(Contains([]string{"apple", "banana", "cherry"}, "banana"))
	fmt.Println(Contains([]int{1,2,3,4,5,6,7,8}, 1))
	fmt.Println(Contains([]ComparableStruct{{1, "Alice", 10}, {2, "Bob", 20}}, ComparableStruct{1, "Alice", 1}))

	// fmt.Println(Contains([]NonComparableStruct{{1, "Alice", []int{1,2,3}}, {2, "Bob", []int{4,5,6}}}, NonComparableStruct{1, "Alice", []int{1,2,3}}))
}