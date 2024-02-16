package search

import "github.com/tel23a-inf/aufgaben-rekursion/lists"

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {
	// TODO

	if lists.Empty(list) {
		return -1
	}

	if list[0] == x {
		return 0
	}
	if Find(list[1:], x) == -1 {
		return -1
	}
	return 1 + Find(list[1:], x)
}
