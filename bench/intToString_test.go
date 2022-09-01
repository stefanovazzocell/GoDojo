package bench

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkIntToString(b *testing.B) {
	b.Run("Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%d", i)
		}
	})
	b.Run("strconvItoa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strconv.Itoa(i)
		}
	})
	b.Run("strconvFormatInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strconv.FormatInt(int64(i), 10)
		}
	})
}
