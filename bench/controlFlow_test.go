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
				continue
			case 1:
				continue
			case 2:
				continue
			case 3:
				continue
			case 4:
				continue
			case 5:
				continue
			case 6:
				continue
			case 7:
				continue
			case 8:
				continue
			case 9:
				continue
			default:
				continue
			}
		}
	})
	b.Run("if", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 11
			if x == 0 {
				continue
			}
			if x == 1 {
				continue
			}
			if x == 2 {
				continue
			}
			if x == 3 {
				continue
			}
			if x == 4 {
				continue
			}
			if x == 5 {
				continue
			}
			if x == 6 {
				continue
			}
			if x == 7 {
				continue
			}
			if x == 8 {
				continue
			}
			if x == 9 {
				continue
			}
			continue
		}
	})
	b.Run("ifElse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = i % 11
			if x == 0 {
				continue
			} else if x == 1 {
				continue
			} else if x == 2 {
				continue
			} else if x == 3 {
				continue
			} else if x == 4 {
				continue
			} else if x == 5 {
				continue
			} else if x == 6 {
				continue
			} else if x == 7 {
				continue
			} else if x == 8 {
				continue
			} else if x == 9 {
				continue
			} else {
				continue
			}
		}
	})
}
