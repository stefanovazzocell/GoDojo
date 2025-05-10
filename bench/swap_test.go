package bench

import "testing"

func BenchmarkSwap(b *testing.B) {
	b.Run("uint8", func(b *testing.B) {
		b.Run("inline", func(b *testing.B) {
			for range b.N {
				n0 := uint8(12)
				n1 := uint8(34)
				n0, n1 = n1, n0
			}
		})
		b.Run("tmp", func(b *testing.B) {
			for range b.N {
				n0 := uint8(12)
				n1 := uint8(34)
				t := n0
				n0 = n1
				n1 = t
			}
		})
		b.Run("xor", func(b *testing.B) {
			for range b.N {
				n0 := uint8(12)
				n1 := uint8(34)
				n0 ^= n1
				n1 ^= n0
				n0 ^= n1
			}
		})
	})
	b.Run("int", func(b *testing.B) {
		b.Run("inline", func(b *testing.B) {
			for range b.N {
				n0 := 123
				n1 := 456
				n0, n1 = n1, n0
			}
		})
		b.Run("tmp", func(b *testing.B) {
			for range b.N {
				n0 := 123
				n1 := 456
				t := n0
				n0 = n1
				n1 = t
			}
		})
		b.Run("xor", func(b *testing.B) {
			for range b.N {
				n0 := 123
				n1 := 456
				n0 ^= n1
				n1 ^= n0
				n0 ^= n1
			}
		})
	})
}
