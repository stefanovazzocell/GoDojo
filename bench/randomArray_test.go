package bench

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func random64SliceBuffer() []byte {
	b := bytes.Buffer{}
	b.Grow(64)
	for i := 0; i < 64; i++ {
		b.WriteByte(byte(rand.Intn(256)))
	}
	return b.Bytes()
}

func randomSlice64Make() []byte {
	b := make([]byte, 64)
	for i := 0; i < 64; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

func random64SliceReader() []byte {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func random64SliceAppend() []byte {
	b := bytes.Buffer{}
	b.Grow(64)
	for i := 0; i < 64; i++ {
		b.WriteByte(byte(rand.Intn(256)))
	}
	return b.Bytes()
}

func random64Array() (arr [64]byte) {
	for i := 0; i < 64; i++ {
		arr[i] = byte(rand.Intn(256))
	}
	return
}

func random64ArrayExpanded() [64]byte {
	return [64]byte{
		byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)),
	}
}

func BenchmarkRandom64Bytes(b *testing.B) {
	tests := map[string]func(){
		"random64SliceBuffer":   func() { _ = random64SliceBuffer() },
		"randomSlice64Make":     func() { _ = randomSlice64Make() },
		"random64SliceAppend":   func() { _ = random64SliceAppend() },
		"random64SliceReader":   func() { _ = random64SliceReader() },
		"random64Array":         func() { _ = random64Array() },
		"random64ArrayExpanded": func() { _ = random64ArrayExpanded() },
	}
	rand.Seed(time.Now().Unix())
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test()
			}
		})
	}
}
