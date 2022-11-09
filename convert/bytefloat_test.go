package convert

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

func TestByteFloatBuffer(t *testing.T) {
	testCase(t, []float64{}, 0, 0, []float64{}, 0)
	testCase(t, []float64{}, 0, 10, []float64{}, 0)
	testCase(t, []float64{}, 0, -100, []float64{}, 0)
	testCase(t, []float64{}, 5, 0, []float64{}, 5)
	testCase(t, []float64{}, 6, 10, []float64{}, 6)
	testCase(t, []float64{}, 7, -100, []float64{}, 7)
	testCase(t, []float64{}, 16, 2, []float64{0, 0}, 0)
	testCase(t, []float64{1}, 0, 0, []float64{}, 8)
	testCase(t, []float64{1}, 0, 10, []float64{1}, 0)
	testCase(t, []float64{1}, 0, -100, []float64{}, 8)
	testCase(t, []float64{1}, 5, 0, []float64{}, 13)
	testCase(t, []float64{1}, 6, 1, []float64{1}, 6)
	testCase(t, []float64{1}, 7, -100, []float64{}, 15)
	testCase(t, []float64{3, 4, 5}, 0, 3, []float64{3, 4, 5}, 0)
	testCase(t, []float64{3, 4, 5}, 1, 3, []float64{3, 4, 5}, 1)
	testCase(t, []float64{3, 4, 5.5, 6.6, 7.7, 8.0}, 7, 2, []float64{3, 4}, 4*8+7)
	testCase(t, []float64{3.3, 4.4, 5.5, 6.6, 7.7, 8.0, 9.0}, 8, 5, []float64{3.3, 4.4, 5.5, 6.6, 7.7}, 2*8+8)
	testCase(t, []float64{1, 2, 3, 4}, 32, 123456, []float64{1, 2, 3, 4, 0, 0, 0, 0}, 0)
}

func getEmptyByteFloatBuffer() *ByteFloatBuffer {
	b := &ByteFloatBuffer{}
	b.Initialize()
	return b
}

func getByteFloatBufferWithFloatsAndNZeroBytes(xArr []float64, n int) *ByteFloatBuffer {
	b := getEmptyByteFloatBuffer()
	var bytes [8]byte
	for _, x := range xArr {
		floatBits := math.Float64bits(x)
		binary.BigEndian.PutUint64(bytes[:], floatBits)
		b.Put(bytes[:])
	}
	b.Put(make([]byte, n))
	return b
}

func arraysEqual(a []float64, b []float64) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, aElement := range a {
		bElement := b[i]
		if aElement != bElement {
			return false
		}
	}
	return true
}

func testCase(t *testing.T, putting []float64, extraBytes int, getting int, expecting []float64, remainingBytes int) {
	t.Run(fmt.Sprintf("(%df,%db).Get(%d)=(%df,%db)", len(putting), extraBytes, getting, len(expecting), remainingBytes),
		func(t *testing.T) {
			b := getByteFloatBufferWithFloatsAndNZeroBytes(putting, extraBytes)

			got := b.Get(getting)
			if !arraysEqual(got, expecting) {
				t.Errorf("got %v, want %v", got, expecting)
			}

			if len(b.buffer) != remainingBytes {
				t.Errorf("%v bytes remain in the buffer instead of %v", len(b.buffer), remainingBytes)
			}
		})
}
