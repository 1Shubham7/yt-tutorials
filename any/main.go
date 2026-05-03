package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x any

	var y int = 1
	fmt.Print(y)

	x = 1
	fmt.Printf("values of x is %v, type is %T\n", x, x)

	x = "hi"
	fmt.Printf("values of x is %v, type is %T\n", x, x)

	x = 1.2
	fmt.Printf("values of x is %v, type is %T\n", x, x)

	mixedSlice := []any{1, "hi", 1.2}
	fmt.Printf("values of mixedSlice is %v, type is %T\n", mixedSlice, mixedSlice)

	var result map[string]any

	json.Unmarshal([]byte(`{"name": "John", "age": 30}`), &result)
	fmt.Printf("values of result is %v, type is %T\n", result, result)
}