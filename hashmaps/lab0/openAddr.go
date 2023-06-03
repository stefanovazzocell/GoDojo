package lab0

const (
	arraySize_OA = 1 << 16
	maxJumps     = 1 << 5
)

// A open addressing based Hashmap (implements HashMap interface)
// WARNING:
// As currently designed it does override entries
type HashTable_OA struct {
	nodes [arraySize_OA]*node_OA
}

type node_OA struct {
	key   int
	value int
	del   bool // If true this is just a placeholder
}

func (ht *HashTable_OA) Init() {
}

func (ht *HashTable_OA) Name() string {
	return "OpenAddressing"
}

func (ht *HashTable_OA) Set(key int, value int) {
	h := hasher(key, arraySize_OA)
	// Limited to maxJumps jumps due to lack of scaling logic
	lastValid := -1
	for i := 0; i < maxJumps; i++ {
		if ht.nodes[(h+i)%arraySize_OA] == nil {
			ht.nodes[(h+i)%arraySize_OA] = &node_OA{
				key:   key,
				value: value,
				del:   false,
			}
			return
		}
		if ht.nodes[(h+i)%arraySize_OA].key == key && !ht.nodes[(h+i)%arraySize_OA].del {
			ht.nodes[(h+i)%arraySize_OA].value = value
			return
		}
		if ht.nodes[(h+i)%arraySize_OA].del {
			lastValid = (h + i) % arraySize_OA
		}
	}
	if lastValid != -1 {
		ht.nodes[lastValid].del = false
		ht.nodes[lastValid].value = value
	}
}

func (ht *HashTable_OA) Delete(key int) {
	h := hasher(key, arraySize_OA)
	// Limited to maxJumps jumps due to lack of scaling logic
	for i := 0; i < maxJumps; i++ {
		if ht.nodes[(h+i)%arraySize_OA] == nil {
			return
		}
		if ht.nodes[(h+i)%arraySize_OA].key == key && !ht.nodes[(h+i)%arraySize_OA].del {
			ht.nodes[(h+i)%arraySize_OA].del = true
			return
		}
	}
}

func (ht *HashTable_OA) Get(key int) (int, bool) {
	h := hasher(key, arraySize_OA)
	// Limited to maxJumps jumps due to lack of scaling logic
	for i := 0; i < maxJumps; i++ {
		if ht.nodes[(h+i)%arraySize_OA] == nil {
			return 0, false
		}
		if ht.nodes[(h+i)%arraySize_OA].key == key && !ht.nodes[(h+i)%arraySize_OA].del {
			return ht.nodes[(h+i)%arraySize_OA].value, true
		}
	}
	return 0, false // faking "not found" (but we don't actually know)
}
