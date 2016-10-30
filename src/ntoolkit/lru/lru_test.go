package lru_test

import (
	"fmt"
	"ntoolkit/assert"
	"ntoolkit/lru"
	"testing"
)

func TestNewCache(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := lru.New(10)
		T.Assert(instance != nil)
	})
}

func TestSetGet(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := lru.New(10)
		T.Assert(instance.Set("1", 100) == nil)

		value, ok := instance.Get("1")
		T.Assert(ok)
		T.Assert(value.(int) == 100)
	})
}

func TestSetGetEvicted(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := lru.New(1)
		T.Assert(instance.Set("1", 100) == nil)
		T.Assert(instance.Set("2", 200) == nil)

		value, ok := instance.Get("1")
		T.Assert(!ok)
		T.Assert(value == nil)

		value, ok = instance.Get("2")
		T.Assert(ok)
		T.Assert(value.(int) == 200)
	})
}

func TestCacheEviction(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := lru.New(10)
		for i := 0; i < 20; i++ {
			instance.Set(fmt.Sprintf("%d", i), i)
		}
		for i := 0; i < 20; i++ {
			value, _ := instance.Get(fmt.Sprintf("%d", i))
			if i >= 10 {
				T.Assert(value != nil)
				T.Assert(value.(int) == i)
			} else {
				T.Assert(value == nil)
			}
		}
	})
}

func TestResize(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := lru.New(20)
		for i := 0; i < 20; i++ {
			instance.Set(fmt.Sprintf("%d", i), i)
		}

		T.Assert(instance.Used() == 20)
		T.Assert(instance.Free() == 0)

		instance.Resize(10)

		T.Assert(instance.Used() == 10)
		T.Assert(instance.Free() == 0)

		instance.Resize(20)

		T.Assert(instance.Used() == 10)
		T.Assert(instance.Free() == 10)

		instance.Clear()
		T.Assert(instance.Used() == 0)
		T.Assert(instance.Free() == 20)
	})
}

func TestErrors(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := lru.New(2)
		for i := 0; i < 2; i++ {
			instance.Set(fmt.Sprintf("%d", i), i)
		}
		T.Assert(instance.Resize(-1) != nil)

		instance = lru.New(0)
		T.Assert(instance.Set("1", 1) != nil)
	})
}
