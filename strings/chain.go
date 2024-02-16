package strings

// Chain erwartet einen String und hängt ihn n mal hintereinander.
// Liefert das Ergebnis.
func Chain(s string, n int) string {
	// TODO
	if n == 1 {
		return s
	}
	return s + Chain(s, n-1)
}
