package slices

func EqualToByte(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func ContainsByte(a []byte, b ...[]byte) bool {
	for k, _ := range b {
		if EqualToByte(a, b[k]) {
			return true
		}
	}

	return false
}
