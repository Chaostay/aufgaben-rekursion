package search

import "github.com/tel23a-inf/aufgaben-rekursion/lists"

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.

func FindSorted2(list []int, x int) int {
	// TODO
	if lists.Empty(list) {
		return -1
	}
	if Find(list, x) == -1 {
		return -1
	}

	if list[len(list)/2] < x {
		return len(list)/2 + Find(list[len(list)/2:], x)
	}
	//Schaut ob x größer als alle Elemente in der ersten Hälfte der gegebenen Liste ist, wenn ja,
	//wiederholt es sich mit der ersten Hälfte der ersten Hälfte etc. bis x gleich oder größer der ersten Hälfte ist.

	if list[len(list)/2] > x {
		return FindSorted2(list[:len(list)/2], x)
	}
	//Schaut ob x kleiner als alle Elemente in der zweiten Hälfte der gegebenen Liste ist, wenn ja,
	//wiederholt es sich mit der zweiten Hälfte der zweiten Hälfte etc. bis x gleich oder kleiner der zweiten Hälfte ist.
	//Bei jeder Wiederholung wird die Halbe Länge der Liste len(list)/2 zum Zähler addiert um an den richtigen Index zu kommen.

	return len(list)
	// Wenn keine der Oberen if Bedingungen greift ist x gleich des Elementes an der Hälfte der gegebenen Liste
	//und es wird die Länge der Liste zum Zähler addiert
}

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	//Fall 1 (wenn Liste leer ist)
	if len(list) == 0 {
		return -1
	}

	if Find(list, x) == -1 {
		return -1
	}
	c := len(list) / 2
	//Fall 2 (wenn wert an mitte genau x)
	if list[c] == x {
		return c
	}
	sublist := list[:c]
	offset := 0 // Offset für den Index in der Teilliste
	//Fall 3 (wenn x größer als wert an mitte)
	if x > list[c] {
		sublist = list[c+1:]
		offset = c + 1
	}
	return offset + FindSorted(sublist, x)
}
