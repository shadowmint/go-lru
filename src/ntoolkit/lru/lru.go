package lru

import (
	"container/list"
	"ntoolkit/errors"
)

// Cache implements a simple LRU strategy
type Cache struct {
	limit     int
	objectMap map[string]*list.Element
	objects   *list.List
}

// New returns a new cache instance specifying the maximum number of objects to store.
func New(limit int) *Cache {
	return &Cache{
		limit:     limit,
		objects:   list.New(),
		objectMap: make(map[string]*list.Element)}
}

// Get returns a record, and moves it to the back of the eviction list
func (cache *Cache) Get(key string) (interface{}, bool) {
	element, ok := cache.objectMap[key]
	if !ok {
		return nil, false
	}

	rtn := element.Value
	if element != cache.objects.Back() {
		cache.objects.MoveToBack(element)
	}

	record, _ := rtn.(*cacheRecord)
	return record.Value, true
}

// Set a cache record to the given value, if not enough space, evic a spot
func (cache *Cache) Set(key string, value interface{}) error {
	if cache.objects.Len() >= cache.limit {
		err := cache.evict()
		if err != nil {
			return err
		}
	}

	element := cache.objects.PushBack(&cacheRecord{
		Key:   key,
		Value: value})
	cache.objectMap[key] = element
	return nil
}

// Clear the entire cache
func (cache *Cache) Clear() {
	cache.objects = list.New()
	cache.objectMap = make(map[string]*list.Element)
}

// Used returns the number of items currently in the cache
func (cache *Cache) Used() int {
	return cache.objects.Len()
}

// Free returns the number of items that can be cached without an eviction
func (cache *Cache) Free() int {
	return cache.limit - cache.Used()
}

// Resize changes the cache size limit, and evics items until the constraint it met
func (cache *Cache) Resize(size int) error {
	cache.limit = size
	for cache.Used() > cache.limit {
		err := cache.evict()
		if err != nil {
			return err
		}
	}
	return nil
}

// evic finds the least recently used cache entry and evics it from the list
func (cache *Cache) evict() error {
	element := cache.objects.Front()
	if element == nil {
		return errors.Fail(&ErrEvictionFailed{}, nil, "No records in the cache to evict")
	}
	record := element.Value.(*cacheRecord)
	cache.objects.Remove(element)
	delete(cache.objectMap, record.Key)
	return nil
}
