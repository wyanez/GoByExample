package main

import "fmt"

/*
	In Go, an array is a numbered sequence of elements of a specific length.
	In a typical Go code, slices are much more common; arrays are useful
	in some special scenarios.
*/

func main() {
	/*
		Here we create an array a that will hold exactly 5 ints.
		The type of elements and length are both part of the array's type.
		By default an array is zero-valued, which for ints means 0s.
	*/
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the array[index] = value syntax
	a[4] = 100
	fmt.Println("set:", a)
	// and get a value with the array[index].
	fmt.Println("get:", a[4])

	//the builtin len returns the length of an array
	fmt.Println("len:", len(a))

	// We can use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can compose types to build
	// multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	var threeD [2][3][4]int
	for i := 0; i < len(threeD); i++ {
		for j := 0; j < len(threeD[i]); j++ {
			for k := 0; k < len(threeD[i][j]); k++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("3d:", threeD)
}
