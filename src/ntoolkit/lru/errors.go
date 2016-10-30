package lru

// ErrEvictionFailed is raised when an attempt is made to evict a record and the cache is empty.
type ErrEvictionFailed struct{}
