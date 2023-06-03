package lab0

// A HashTable that uses the default map type for Go
type HashTable_Default struct {
	m map[int]int
}

func (ht *HashTable_Default) Init() {
	ht.m = make(map[int]int)
}

func (ht *HashTable_Default) Name() string {
	return "Default"
}

func (ht *HashTable_Default) Set(key int, value int) {
	ht.m[key] = value
}

func (ht *HashTable_Default) Delete(key int) {
	delete(ht.m, key)
}

func (ht *HashTable_Default) Get(key int) (int, bool) {
	value, found := ht.m[key]
	return value, found
}
