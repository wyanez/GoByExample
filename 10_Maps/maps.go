package main

import (
	"fmt"
	"maps"
)

/*
Maps are Go's built-in associative data type
(sometimes called hashes or dicts in other languages).
*/

func main() {
	//To create an empty map, use builtin make:
	//make( map[ley-type]val-type)
	m := make(map[string]int)

	//Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

	//Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("get m[k1]:", v1)

	//If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("get m[k3]:", v3)

	//The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	//The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map after delete k2:", m, len(m))

	//To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("map after clear:", m, len(m))

	//The optional second return value when getting a value from a map indicates
	//if the key was present in the map. This can be used to disambiguate between
	//missing keys and keys with zero values like 0 or "".
	//Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, present := m["k2"]
	fmt.Println("k2 present ?", present)

	//You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	n["xyz"] = 3
	fmt.Println("map n:", n)

	//The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"bar": 2, "xyz": 3}
	n2["foo"] = 1
	fmt.Println("map n2:", n2)
	if maps.Equal(n, n2) {
		fmt.Println("map n == n2")
	}
}
