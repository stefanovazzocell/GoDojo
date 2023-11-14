package bench

import "testing"

// Switch between 2 cases
func BenchmarkControlFlowSimple(b *testing.B) {
	var x int
	b.Run("switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 2
			switch x {
			case 0:
				continue
			case 1:
				continue
			}
		}
	})
	b.Run("if", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 2
			if x == 0 {
				continue
			}
			if x == 1 {
				continue
			}
		}
	})
	b.Run("ifElse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 2
			if x == 0 {
				continue
			} else if x == 1 {
				continue
			}
		}
	})
}

// Switch between 10 cases and default
func BenchmarkControlFlowComplex(b *testing.B) {
	var x int
	b.Run("switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 11
			switch x {
			case 0:
				_ = 0
				continue
			case 1:
				_ = 0
				continue
			case 2:
				_ = 0
				continue
			case 3:
				_ = 0
				continue
			case 4:
				_ = 0
				continue
			case 5:
				_ = 0
				continue
			case 6:
				_ = 0
				continue
			case 7:
				_ = 0
				continue
			case 8:
				_ = 0
				continue
			case 9:
				_ = 0
				continue
			default:
				_ = 0
				continue
			}
		}
	})
	b.Run("if", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 11
			if x == 0 {
				_ = 0
				continue
			}
			if x == 1 {
				_ = 0
				continue
			}
			if x == 2 {
				_ = 0
				continue
			}
			if x == 3 {
				_ = 0
				continue
			}
			if x == 4 {
				_ = 0
				continue
			}
			if x == 5 {
				_ = 0
				continue
			}
			if x == 6 {
				_ = 0
				continue
			}
			if x == 7 {
				_ = 0
				continue
			}
			if x == 8 {
				_ = 0
				continue
			}
			if x == 9 {
				_ = 0
				continue
			}
			_ = 0
			continue
		}
	})
	b.Run("ifElse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 11
			if x == 0 {
				_ = 0
			} else if x == 1 {
				_ = 0
			} else if x == 2 {
				_ = 0
			} else if x == 3 {
				_ = 0
			} else if x == 4 {
				_ = 0
			} else if x == 5 {
				_ = 0
			} else if x == 6 {
				_ = 0
			} else if x == 7 {
				_ = 0
			} else if x == 8 {
				_ = 0
			} else if x == 9 {
				_ = 0
			} else {
				_ = 0
			}
		}
	})
}
