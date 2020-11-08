package utils

func FilterByte(bt []byte, f uint8) []byte {
	offs := 0

	for i, b := range bt {
		if b == f {
			offs++
			continue
		}

		bt[i-offs] = b
	}

	totalLen := len(bt) - offs

	return bt[:totalLen:totalLen]
}
