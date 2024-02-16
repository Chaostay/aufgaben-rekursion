package lists

import "fmt"

func ExamplePartition() {
	l1 := []int{10, 3, 17, 5, 23, 12, 11}

	smaller, larger := Partition(l1, 11)
	fmt.Println(smaller)
	fmt.Println(larger)

	// Output:
	// [10 3 5 11]
	// [17 23 12]
}
