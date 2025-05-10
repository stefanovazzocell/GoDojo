package bench

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func makeByteSlice(size int) []byte {
	bs := make([]byte, size)
	for i := range size {
		bs[i] = byte(i)
	}
	return bs
}

func BenchmarkSha256three(b *testing.B) {
	smallSlices := [][]byte{
		makeByteSlice(16),
		makeByteSlice(30),
		makeByteSlice(8),
	}
	largeSlices := [][]byte{
		makeByteSlice(512),
		makeByteSlice(1000),
		makeByteSlice(256),
	}
	for _, testSize := range []string{"small", "large"} {
		testSlices := smallSlices
		if testSize == "large" {
			testSlices = largeSlices
		}

		b.Run(testSize, func(b *testing.B) {
			b.Run("manyWrites", func(b *testing.B) {
				for range b.N {
					hash := sha256.New()
					hash.Write(testSlices[0])
					hash.Write(testSlices[1])
					hash.Write(testSlices[2])
					_ = hash.Sum(nil)
				}
			})
			b.Run("append", func(b *testing.B) {
				for range b.N {
					hash := sha256.New()
					hash.Write(append(testSlices[0], append(testSlices[1], testSlices[2]...)...))
					_ = hash.Sum(nil)
				}
			})
			b.Run("buffer", func(b *testing.B) {
				for range b.N {
					buffer := bytes.Buffer{}
					buffer.Grow(len(testSlices[0]) + len(testSlices[1]) + len(testSlices[2]))
					buffer.Write(testSlices[0])
					buffer.Write(testSlices[1])
					buffer.Write(testSlices[2])
					hash := sha256.New()
					hash.Write(buffer.Bytes())
					_ = hash.Sum(nil)
				}
			})
			b.Run("preprocess", func(b *testing.B) {
				// This is a bit unfair compared to the other tests but it can
				// be helpful to get a baseline
				data := append(testSlices[0], testSlices[1]...)
				data = append(data, testSlices[2]...)

				for range b.N {
					hash := sha256.New()
					hash.Write(data)
					_ = hash.Sum(nil)
				}
			})
		})
	}
}

// Validates that all examples above produce the same result
func TestValidateEqual(t *testing.T) {
	testSlices := [][]byte{
		makeByteSlice(151),
		makeByteSlice(3),
		makeByteSlice(17),
	}

	// 1
	hash := sha256.New()
	hash.Write(testSlices[0])
	hash.Write(testSlices[1])
	hash.Write(testSlices[2])
	_manyWrites := hash.Sum(nil)

	// 2
	hash = sha256.New()
	hash.Write(append(testSlices[0], append(testSlices[1], testSlices[2]...)...))
	_append := hash.Sum(nil)

	// 3
	buffer := bytes.Buffer{}
	buffer.Grow(len(testSlices[0]) + len(testSlices[1]) + len(testSlices[2]))
	buffer.Write(testSlices[0])
	buffer.Write(testSlices[1])
	buffer.Write(testSlices[2])
	hash = sha256.New()
	hash.Write(buffer.Bytes())
	_buffer := hash.Sum(nil)

	// 4
	data := append(testSlices[0], testSlices[1]...)
	data = append(data, testSlices[2]...)
	hash = sha256.New()
	hash.Write(data)
	_preprocess := hash.Sum(nil)

	// Assert they all match
	if !bytes.Equal(_manyWrites, _append) ||
		!bytes.Equal(_append, _buffer) ||
		!bytes.Equal(_buffer, _preprocess) {
		t.Fatalf("The results differ:\n_manyWrites=%x\n_append=%x\n_buffer=%x\n_preprocess=%x",
			_manyWrites, _append, _buffer, _preprocess)
	}
}
