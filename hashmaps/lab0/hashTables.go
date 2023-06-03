package lab0

// A generic HashMap interface to simplify testing
type HashMap interface {
	// Initialize the map
	Init()
	// Return the name for this HashMap implementation
	Name() string
	// Sets a value for a key
	Set(key int, value int)
	// Deletes an entry
	Delete(key int)
	// Gets the value for a key (or false if no such entry found)
	Get(key int) (int, bool)
}

// A trivial hash function
func hasher(i int, buckets int) int {
	return i % buckets
}
