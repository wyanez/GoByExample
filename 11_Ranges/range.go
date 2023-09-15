package main

import "fmt"

/*
range iterates over elements in a variety of data structures.
Let's see how to use range with some of the data structures
we've already learned
*/

func main() {
	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	fmt.Println("Iterating over a slice getting the values only")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	fmt.Println("Iterating over a slice getting both index and value")
	for index, num := range nums {
		fmt.Printf("index: %d -> value: %d \n", index, num)
	}

	// range on map iterates over key/value pairs.
	fmt.Println("Iterating over the keys,values of a map")
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	//range can also iterate over just the keys of a map.
	fmt.Println("Iterating over just the keys of a map")
	for key := range kvs {
		fmt.Printf("key: %s\n", key)
	}

	//range on strings iterates over Unicode code points.
	//The first value is the starting byte index of the rune and the second
	//the rune itself. See Strings and Runes for more details.
	fmt.Println("Iterating over strings")
	for index, c := range "golang" {
		fmt.Println(index, c)
	}

}
