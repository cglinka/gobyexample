package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "bananna"}
	for k, v := range kvs {
		// Does not interpret anything, just prints the string, then the two values
		// fmt.Println("%s -> %s\n", k, v)
		// Does interperet the string, prints "a -> apple", etc
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
