package lru

import (
	"container/list"
	"ntoolkit/errors"
	"ntoolkit/iter"
)

type keyValueIter struct {
	values map[string]*list.Element
	keys   []string
	offset int
	err    error
}

// KeyValue is a single entry in the lru cache.
type KeyValue struct {
	Key   string
	Value interface{}
}

func newKeyValueIter(values map[string](*list.Element)) iter.Iter {
	rtn := &keyValueIter{values: values}
	rtn.Init()
	return rtn
}

func (iterator *keyValueIter) Init() {
	iterator.offset = -1
	iterator.keys = make([]string, 0, len(iterator.values))
	for k := range iterator.values {
		iterator.keys = append(iterator.keys, k)
	}
}

func (iterator *keyValueIter) Next() (interface{}, error) {
	if iterator.err != nil {
		return nil, iterator.err
	}

	iterator.offset++
	if iterator.offset >= len(iterator.keys) {
		iterator.err = errors.Fail(iter.ErrEndIteration{}, nil, "No more values")
		return nil, iterator.err
	}

	element := iterator.values[iterator.keys[iterator.offset]]
	record := element.Value.(*cacheRecord)
	keyValuePair := &KeyValue{
		iterator.keys[iterator.offset],
		record.Value}
	return keyValuePair, nil
}
