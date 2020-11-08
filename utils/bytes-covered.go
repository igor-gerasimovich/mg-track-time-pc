package utils

import "bytes"

func BytesCoveredRune(bt []byte, cover uint8) [][]byte {
	parts := make([][]byte, 0)

	var (
		cPart   bytes.Buffer
		started bool
	)
	for _, b := range bt {
		if b == cover {
			if started {
				parts = append(parts, cPart.Bytes())

				// Need this to reallocate buffer ()
				cPart = bytes.Buffer{}
			}

			started = !started
			continue
		}

		if started {
			cPart.WriteByte(b)
		}
	}

	return parts
}
