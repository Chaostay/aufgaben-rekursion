package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	// TODO

	if s == seq {
		return true
	}
	n := 0
	m := n + len(seq)
	if len(s) < len(seq) {
		return false
	}

	if s[n:m] == seq {
		return true
	}
	return false
}
