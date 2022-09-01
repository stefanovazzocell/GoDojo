package bench

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkTextConcatSmall(b *testing.B) {
	b.Run("Sprint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprint("first", "second", "third")
		}
	})
	b.Run("Plus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = "first" + "second" + "third"
		}
	})
	b.Run("Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b := strings.Builder{}
			b.WriteString("first")
			b.WriteString("second")
			b.WriteString("third")
			_ = b.String()
		}
	})
	b.Run("Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Join([]string{"first", "second", "third"}, "")
		}
	})
}

func BenchmarkTextConcatLarge(b *testing.B) {
	b.Run("Plus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out := ""
			for j := 0; j < 100; j++ {
				out += "string"
			}
			_ = out
		}
	})
	b.Run("Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b := strings.Builder{}
			for i := 0; i < 100; i++ {
				b.WriteString("string")
			}
			_ = b.String()
		}
	})
	b.Run("BuilderPrealloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b := strings.Builder{}
			b.Grow(600) // len("string") * 100
			for i := 0; i < 100; i++ {
				b.WriteString("string")
			}
			_ = b.String()
		}
	})
	b.Run("Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strs := []string{}
			for i := 0; i < 100; i++ {
				strs = append(strs, "string")
			}
			_ = strings.Join(strs, "")
		}
	})
	b.Run("JoinReserved", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strs := make([]string, 0, 100)
			for i := 0; i < 100; i++ {
				strs = append(strs, "string")
			}
			_ = strings.Join(strs, "")
		}
	})
	b.Run("JoinPrealloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strs := make([]string, 100)
			for i := 0; i < 100; i++ {
				strs[i] = "string"
			}
			_ = strings.Join(strs, "")
		}
	})
}
