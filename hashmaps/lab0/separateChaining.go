package lab0

const (
	bucketSize_SC = 1 << 12
)

// A separate chaining based Hashmap (implements HashMap interface)
type HashTable_SC struct {
	buckets [bucketSize_SC]*bucket_SC
}

type bucket_SC struct {
	head *node_SC
}

type node_SC struct {
	key  int
	data int
	next *node_SC
}

func (ht *HashTable_SC) Init() {
	for i := 0; i < bucketSize_SC; i++ {
		ht.buckets[i] = &bucket_SC{
			head: nil,
		}
	}
}

func (ht *HashTable_SC) Name() string {
	return "SeparateChaining"
}

func (ht *HashTable_SC) Set(key int, value int) {
	h := hasher(key, bucketSize_SC)
	ht.buckets[h].Set(key, value)
}

func (ht *HashTable_SC) Delete(key int) {
	h := hasher(key, bucketSize_SC)
	ht.buckets[h].Delete(key)
}

func (ht *HashTable_SC) Get(key int) (int, bool) {
	h := hasher(key, bucketSize_SC)
	return ht.buckets[h].Get(key)
}

func (b *bucket_SC) Set(key int, value int) {
	current := b.head
	// Update node (if exists)
	for current != nil {
		if current.key == key {
			// Found!
			current.data = value
			return
		}
		current = current.next
	}
	// Add new
	b.head = &node_SC{
		key:  key,
		data: value,
		next: b.head,
	}
}

func (b *bucket_SC) Delete(key int) {
	// Special cases
	if b.head == nil {
		return
	}
	if b.head.next == nil {
		if b.head.key == key {
			b.head = nil
		}
		return
	}
	// Complex cases
	previous := b.head
	for previous.next != nil {
		if previous.next.key == key {
			// Found!
			previous.next = previous.next.next
			return
		}
		previous = previous.next
	}
}

func (b *bucket_SC) Get(key int) (int, bool) {
	current := b.head
	for current != nil {
		if current.key == key {
			// Found!
			return current.data, true
		}
		current = current.next
	}
	return 0, false
}
