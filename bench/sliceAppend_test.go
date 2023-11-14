package bench

import "testing"

// Testing the scenario where we have a slice with 3 elements
// and we know we're going to have to add 1 more element
func BenchmarkSliceAppendOne(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Initial
			s := []string{
				"item:0",
				"item:2",
				"item:3",
			}
			// Add one more
			s = append(s, "item:4")
			_ = s
		}
	})
	b.Run("appendEachWithHint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Initial
			s := make([]string, 0, 4)
			s = append(s, "item:1")
			s = append(s, "item:2")
			s = append(s, "item:3")
			// Add one more
			s = append(s, "item:4")
			_ = s
		}
	})
	b.Run("appendSliceWithHint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Initial
			s := make([]string, 0, 4)
			s = append(s, []string{
				"item:1",
				"item:2",
				"item:3",
			}...)
			// Add one more
			s = append(s, "item:4")
			_ = s
		}
	})
	b.Run("appendTogetherWithHint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Initial
			s := make([]string, 0, 4)
			s = append(s,
				"item:1",
				"item:2",
				"item:3",
			)
			// Add one more
			s = append(s, "item:4")
			_ = s
		}
	})
	b.Run("preSized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Initial
			s := make([]string, 3, 4)
			s[0] = "item:1"
			s[1] = "item:2"
			s[2] = "item:3"
			// Add one more
			s = append(s, "item:4")
			_ = s
		}
	})
	b.Run("preSizedAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Initial
			s := make([]string, 4)
			s[0] = "item:1"
			s[1] = "item:2"
			s[2] = "item:3"
			// Add one more
			s[3] = "item:4"
			_ = s
		}
	})
}
