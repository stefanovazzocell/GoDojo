package bench

import "testing"

func BenchmarkSliceLoopValue(b *testing.B) {
	testSlice := make([]string, 1000000)
	b.Run("rangeKey", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for k := range testSlice {
				_ = testSlice[k]
			}
		}
	})
	b.Run("rangeVal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range testSlice {
				_ = v
			}
		}
	})
	b.Run("len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(testSlice); j++ {
				_ = testSlice[j]
			}
		}
	})
	b.Run("lenOptimizedA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := len(testSlice)
			for j := 0; j < l; j++ {
				_ = testSlice[j]
			}
		}
	})
	b.Run("lenOptimizedB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := len(testSlice) - 1; j >= 0; j-- {
				_ = testSlice[j]
			}
		}
	})
}

func BenchmarkSliceLoopCount(b *testing.B) {
	testSlice := make([]string, 1000000)
	var count int
	b.Run("rangeKey", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count = 0
			for k := range testSlice {
				_ = k
				count++
			}
		}
	})
	b.Run("rangeVal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count = 0
			for v := range testSlice {
				_ = v
				count++
			}
		}
	})
	b.Run("rangeBlank", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count = 0
			for range testSlice {
				count++
			}
		}
	})
	b.Run("len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count = 0
			for j := 0; j < len(testSlice); j++ {
				count++
			}
		}
	})
	b.Run("lenOptimizedA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count = 0
			l := len(testSlice)
			for j := 0; j < l; j++ {
				count++
			}
		}
	})
	b.Run("lenOptimizedB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count = 0
			for j := len(testSlice) - 1; j >= 0; j-- {
				count++
			}
		}
	})
}
