package bench

import (
	"testing"
	"unsafe"
)

//go:noescape
//go:linkname strhash runtime.strhash
func strhash(p unsafe.Pointer, h uintptr) uintptr

func BenchmarkStrHash(b *testing.B) {
	str := "Hello World"
	for i := 0; i < b.N; i++ {
		_ = strhash(unsafe.Pointer(&str), 0)
	}
}

func TestStrHash(t *testing.T) {
	strA0 := "Hello World"
	hA0 := strhash(unsafe.Pointer(&strA0), 0)
	strA1 := "Hello World"
	hA1 := strhash(unsafe.Pointer(&strA1), 0)
	strB0 := "Hello World!"
	hB0 := strhash(unsafe.Pointer(&strB0), 0)
	if hA0 != hA1 {
		t.Fatalf("hA0 != hA1 [%d != %d]", hA0, hA1)
	}
	if hA0 == hB0 {
		t.Fatalf("hA0 == hB0 [%d == %d]", hA0, hB0)
	}
	t.Logf("%d, %d, %d", hA0, hA1, hB0)
	t.FailNow()
}
