# go-lru

A super simple lru cache implementation.

This cache is not thread safe. 

# Usage

    npm install shadowmint/go-lru --save

Then use `Set` and `Get` to set and access records in the cache:

    cache := lru.New(10)  // Maximum 10 items in cache.
    cache.Set("Key1", 1)

    value, ok := cache.Get("Key1")
    if ok {
      iv := value.(int)
    }
