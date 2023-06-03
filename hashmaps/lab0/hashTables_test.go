package lab0

import (
	"testing"
)

var hashTables = []HashMap{
	&HashTable_SC{},
	&HashTable_OA{},
	&HashTable_Default{},
}

const (
	// 1048576
	basicTests int = 1 << 16
)

// Run a check to make sure there's some basic consistency in this specific HashMap implementation
func TestBasics(t *testing.T) {
	for _, ht := range hashTables {
		t.Run(ht.Name(), func(t *testing.T) {
			// Initialize
			ht.Init()
			// Perform a large insert + read back
			for i := 0; i < basicTests*3; i += 3 {
				if value, found := ht.Get(i); value != 0 || found {
					t.Fatalf("Get(%d) returned (%d, %v) but expected no entry here", i, value, found)
				}
				ht.Set(i, i+1)
				if value, found := ht.Get(i); value != i+1 || !found {
					t.Fatalf("Get(%d) returned (%d, %v) but expected (%d, true) as it was just Set", i, value, found, i+1)
				}
			}
			// Perform an override + read back
			for i := 0; i < basicTests*3; i += 3 {
				ht.Set(i, i+2)
				if value, found := ht.Get(i); value != i+2 || !found {
					t.Fatalf("Get(%d) returned (%d, %v) but expected (%d, true) as it was just (re-)Set", i, value, found, i+2)
				}
			}
			// Perform delete
			for i := 0; i < basicTests*3; i += 3 {
				ht.Delete(i)
				if value, found := ht.Get(i); value != 0 || found {
					t.Fatalf("Get(%d) returned (%d, %v) but expected no entry here as it was just Delete(d)", i, value, found)
				}
			}
		})
	}
}

// Benchmark Get
func BenchmarkGet(b *testing.B) {
	for _, ht := range hashTables {
		b.Run(ht.Name(), func(b *testing.B) {
			b.StopTimer()
			ht.Init()
			for i := 0; i < 3*b.N; i += 3 {
				ht.Set(i, i)
			}
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < 3*b.N; i += 3 {
				_, _ = ht.Get(i)
			}
		})
	}
}

// Benchmark Set
func BenchmarkSet(b *testing.B) {
	for _, ht := range hashTables {
		b.Run(ht.Name(), func(b *testing.B) {
			b.StopTimer()
			ht.Init()
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < 3*b.N; i += 3 {
				ht.Set(i, i)
			}
		})
	}
}

// Benchmark Set (where value exists already override)
func BenchmarkSetOverride(b *testing.B) {
	for _, ht := range hashTables {
		b.Run(ht.Name(), func(b *testing.B) {
			b.StopTimer()
			ht.Init()
			for i := 0; i < 3*b.N; i += 3 {
				ht.Set(i, i+1)
			}
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < 3*b.N; i += 3 {
				ht.Set(i, i)
			}
		})
	}
}

// Benchmark Delete (where value exists already override)
func BenchmarkDelete(b *testing.B) {
	for _, ht := range hashTables {
		b.Run(ht.Name(), func(b *testing.B) {
			b.StopTimer()
			ht.Init()
			for i := 0; i < 3*b.N; i += 3 {
				ht.Set(i, i)
			}
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < 3*b.N; i += 3 {
				ht.Delete(i)
			}
		})
	}
}
