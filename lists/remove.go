package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	// TODO
	if Empty(list) {
		return list
	}
	if pos == 0 {
		return list[1:]
	}
	return append([]int{list[0]}, RemoveElement(list[1:], pos-1)...)
}

/*
// Diese Funktion entfernt ein Element aus der Liste, indem es rekursiv
// die Elemente vor und nach der Position `pos` neu anordnet.
// Wenn `pos` außerhalb des Bereichs der Liste liegt, wird die ursprüngliche Liste zurückgegeben.
func RemoveElement(list []int, pos int) []int {
    // Basisfall: Wenn die Liste leer ist oder `pos` außerhalb des Bereichs liegt,
    // geben wir die ursprüngliche Liste zurück.
    if len(list) == 0 || pos < 0 || pos >= len(list) {
        return list
    }

    // Rekursiver Fall: Wir erstellen eine neue Liste, die alle Elemente außer dem
    // an der Position `pos` enthält.
    // Wir verwenden die Rekursion, um die Elemente vor und nach `pos` zu kombinieren.
    return append(RemoveElement(list[:pos], pos), list[pos+1:]...)
}

*/
