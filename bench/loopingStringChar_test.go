package bench

import (
	"strings"
	"testing"
)

func BenchmarkStringCharLoop(b *testing.B) {
	b.Run("len10", func(b *testing.B) {
		builder := strings.Builder{}
		for i := 0; i < 10; i++ {
			builder.WriteByte(byte(i))
		}
		testStr := builder.String()
		b.Run("rangeKey", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for k := range testStr {
					_ = testStr[k]
				}
			}
		})
		b.Run("rangeVal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range testStr {
					_ = v
				}
			}
		})
		b.Run("len", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < len(testStr); j++ {
					_ = testStr[j]
				}
			}
		})
		b.Run("lenOptimizedA", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l := len(testStr)
				for j := 0; j < l; j++ {
					_ = testStr[j]
				}
			}
		})
		b.Run("lenOptimizedB", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := len(testStr) - 1; j >= 0; j-- {
					_ = testStr[j]
				}
			}
		})
		b.Run("stringsReader", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := strings.NewReader(testStr)
				for {
					_, err := r.ReadByte()
					if err != nil {
						break
					}
				}
			}
		})
	})
	b.Run("len1000000", func(b *testing.B) {
		builder := strings.Builder{}
		for i := 0; i < 1000000; i++ {
			builder.WriteByte(byte(i))
		}
		testStr := builder.String()
		b.Run("rangeKey", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for k := range testStr {
					_ = testStr[k]
				}
			}
		})
		b.Run("rangeVal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range testStr {
					_ = v
				}
			}
		})
		b.Run("len", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < len(testStr); j++ {
					_ = testStr[j]
				}
			}
		})
		b.Run("lenOptimizedA", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l := len(testStr)
				for j := 0; j < l; j++ {
					_ = testStr[j]
				}
			}
		})
		b.Run("lenOptimizedB", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := len(testStr) - 1; j >= 0; j-- {
					_ = testStr[j]
				}
			}
		})
		b.Run("stringsReader", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := strings.NewReader(testStr)
				for {
					_, err := r.ReadByte()
					if err != nil {
						break
					}
				}
			}
		})
	})
}
