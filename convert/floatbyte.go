package convert

import (
	"encoding/binary"
	"math"
)

func FloatsToBytes(floats []float64) []byte {
	bytes := make([]byte, 8*len(floats))
	for i, f := range floats {
		fBits := math.Float64bits(f)
		binary.BigEndian.PutUint64(bytes[i*8:i*8+8], fBits)
	}
	return bytes
}
