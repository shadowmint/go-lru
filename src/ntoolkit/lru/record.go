package lru

// cacheRecord is used internally to keep track of records
type cacheRecord struct {
	Key   string
	Value interface{}
}
