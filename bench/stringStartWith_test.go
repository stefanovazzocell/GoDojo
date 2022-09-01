package bench

import (
	"strings"
	"testing"
)

// Returns true if str is prefixed by prefix, otherwise false.
// Compares the strings char-by-char
func LoopStartsWith(str string, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if str[i] != prefix[i] {
			return false
		}
	}
	return true
}

func TestLoopStartsWith(t *testing.T) {
	cases := map[string]bool{
		"hello world;":           true,
		";hello world":           false,
		"hello world;hello":      true,
		"hello world;world":      false,
		"hello world;helloworld": false,
	}
	for strs, expected := range cases {
		params := strings.Split(strs, ";")
		if len(params) != 2 {
			t.Fatalf("Got %d parameters, but expected 2: %+s", len(params), params)
		}
		t.Logf("Testing %q with prefix %q (expecting %v)", params[0], params[1], expected)
		if strings.HasPrefix(params[0], params[1]) != expected {
			t.Fatal("strings.HasPrefix differs from expected value")
		}
		if LoopStartsWith(params[0], params[1]) != expected {
			t.Fatal("LoopStartsWith differs from expected value")
		}
	}
}

func BenchmarkStringStartWithShort(b *testing.B) {
	b.Run("stringsHasPrefix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.HasPrefix("hello world", "")
			strings.HasPrefix("", "hello world")
			strings.HasPrefix("hello world", "hello")
			strings.HasPrefix("hello world", "world")
			strings.HasPrefix("hello world", "helloworld")
		}
	})
	b.Run("LoopStartsWith", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LoopStartsWith("hello world", "")
			LoopStartsWith("", "hello world")
			LoopStartsWith("hello world", "hello")
			LoopStartsWith("hello world", "world")
			LoopStartsWith("hello world", "helloworld")
		}
	})
}

func BenchmarkStringStartWithLong(b *testing.B) {
	b.Run("stringsHasPrefix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.HasPrefix(
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non. Rerum ratione sed est dolorem dolorem. Provident voluptatibus sunt voluptas velit.",
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non.")
			strings.HasPrefix(
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non. Rerum ratione sed est dolorem dolorem. Provident voluptatibus sunt voluptas velit.",
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non. Rerum ratione sed est dolorem dolorem. Provident voluptatibus sunt voluptas velit.")
		}
	})
	b.Run("LoopStartsWith", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LoopStartsWith(
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non. Rerum ratione sed est dolorem dolorem. Provident voluptatibus sunt voluptas velit.",
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non.")
			LoopStartsWith(
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non. Rerum ratione sed est dolorem dolorem. Provident voluptatibus sunt voluptas velit.",
				"Incidunt nihil qui reprehenderit praesentium. Sit consectetur voluptas non. Rerum ratione sed est dolorem dolorem. Provident voluptatibus sunt voluptas velit.")
		}
	})
}
