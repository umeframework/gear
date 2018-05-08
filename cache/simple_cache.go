/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package cache

import "sync"

type simpleCache struct {
	name  string
	dict  map[Key]Value
	mutex sync.RWMutex
}

func NewSimpleCache(name string) Cache {
	return &simpleCache{
		name: name,
		dict: make(map[Key]Value),
	}
}

func (cache *simpleCache) GetKeys() []Key {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	count := len(cache.dict)
	keys := make([]Key, count)
	index := 0
	for key, _ := range cache.dict {
		keys[index] = key
		index++
	}

	return keys
}

func (cache *simpleCache) GetName() string {
	return cache.name
}

func (cache *simpleCache) Get(key Key) (Value, bool, error) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	return cache.getCore(key)
}

func (cache *simpleCache) getCore(key Key) (Value, bool, error) {
	value, ok := cache.dict[key]
	return value, ok, nil
}

func (cache *simpleCache) Put(key Key, value Value) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	return cache.putCore(key, value)
}

func (cache *simpleCache) putCore(key Key, value Value) error {
	cache.dict[key] = value
	return nil
}

func (cache *simpleCache) Compute(key Key, param Param, callback func(key Key, value Value, exists bool, err error, param Param) (Value, bool)) (Value, error) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	value, exists, err := cache.getCore(key)
	newValue, shouldUpdate := callback(key, value, exists, err, param)
	if shouldUpdate {
		err = cache.putCore(key, newValue)
	}

	return newValue, err
}

func (cache *simpleCache) Remove(key Key) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	return cache.removeCore(key)
}

func (cache *simpleCache) removeCore(key Key) error {
	delete(cache.dict, key)
	return nil
}

func (cache *simpleCache) Clear() error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	return cache.clearCore()
}

func (cache *simpleCache) clearCore() error {
	cache.dict = make(map[Key]Value)
	return nil
}
