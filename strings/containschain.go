package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {
	// TODO
	if s == symbol {
		return true
	}

	if len(s) < count {
		return false
	}

	if count == 0 {
		return true
	}

	if Contains(s, Chain(symbol, count)) {
		return true
	}

	return ContainsChain(s[1:], symbol, count)
}
