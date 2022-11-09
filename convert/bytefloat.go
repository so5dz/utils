package convert

import (
	"encoding/binary"
	"math"
)

type ByteFloatBuffer struct {
	buffer []byte
}

func (b *ByteFloatBuffer) Initialize() {
	b.buffer = make([]byte, 0)
}

func (b *ByteFloatBuffer) Put(bytes []byte) {
	b.buffer = append(b.buffer, bytes...)
}

func (b *ByteFloatBuffer) GetAll() []float64 {
	bufferFloatSize := len(b.buffer) / 8
	return b.Get(bufferFloatSize)
}

func (b *ByteFloatBuffer) Get(limit int) []float64 {
	bufferFloatSize := len(b.buffer) / 8
	if limit < 0 {
		limit = 0
	}
	if limit > bufferFloatSize {
		limit = bufferFloatSize
	}

	floats := make([]float64, limit)
	for i := 0; i < limit; i++ {
		floatBytes := b.buffer[i*8 : i*8+8]
		floatBits := binary.BigEndian.Uint64(floatBytes)
		floats[i] = math.Float64frombits(floatBits)
	}

	b.buffer = b.buffer[limit*8:]
	return floats
}
