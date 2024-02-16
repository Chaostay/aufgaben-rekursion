package search

import "fmt"

func ExampleFindSorted() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(FindSorted(l1, 386))

	// Output:
	// -1
}

func ExampleFindSorted2() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(FindSorted2(l1, 100))

	// Output:
	// -1
}
