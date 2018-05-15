/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"time"
	"sync"
)

type simpleCacheExpirable struct {
	name  string
	dict  map[Key]Value

	lastAccessTime time.Time
	maxIdle time.Duration
	expiresOn time.Time

	mutex sync.Mutex
}

func NewSimpleCacheExpirable(name string, maxIdle time.Duration) ExpirableCache {
	cache := simpleCacheExpirable{
		name: name,
		dict: make(map[Key]Value),
		maxIdle: maxIdle,
		mutex: sync.Mutex{},
	}
	cache.updateLastAccessTime(nil)
	return &cache
}

func (cache *simpleCacheExpirable) updateLastAccessTime(lastAccessTime *time.Time) {
	if lastAccessTime == nil {
		cache.lastAccessTime = time.Now()
	} else {
		cache.lastAccessTime = *lastAccessTime
	}
	cache.expiresOn = cache.lastAccessTime.Add(cache.maxIdle)
}

func (cache *simpleCacheExpirable) GetKeys() []Key {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	count := len(cache.dict)
	keys := make([]Key, count)
	index := 0
	for key, _ := range cache.dict {
		keys[index] = key
		index++
	}

	return keys
}

func (cache *simpleCacheExpirable) GetName() string {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	return cache.name
}

func (cache *simpleCacheExpirable) Get(key Key) (Value, bool, error) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	return cache.getCore(key)
}

func (cache *simpleCacheExpirable) getCore(key Key) (Value, bool, error) {
	value, ok := cache.dict[key]
	return value, ok, nil
}

func (cache *simpleCacheExpirable) Put(key Key, value Value) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	return cache.putCore(key, value)
}

func (cache *simpleCacheExpirable) putCore(key Key, value Value) error {
	cache.dict[key] = value
	return nil
}

func (cache *simpleCacheExpirable) Compute(key Key, param Param, callback func(key Key, value Value, exists bool, err error, param Param) (Value, bool)) (Value, error) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	value, exists, err := cache.getCore(key)
	newValue, shouldUpdate := callback(key, value, exists, err, param)
	if shouldUpdate {
		err = cache.putCore(key, newValue)
	}

	return newValue, err
}

func (cache *simpleCacheExpirable) Remove(key Key) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	return cache.removeCore(key)
}

func (cache *simpleCacheExpirable) removeCore(key Key) error {
	delete(cache.dict, key)
	return nil
}

func (cache *simpleCacheExpirable) Clear() error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(nil)

	return cache.clearCore()
}

func (cache *simpleCacheExpirable) clearCore() error {
	cache.dict = make(map[Key]Value)
	return nil
}

func (cache *simpleCacheExpirable) GetLastAccessTime() time.Time {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	return cache.lastAccessTime
}

func (cache *simpleCacheExpirable) SetLastAccessTime(lastAccess time.Time) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(&lastAccess)

	cache.lastAccessTime = lastAccess
}

func (cache *simpleCacheExpirable) IsExpired() bool {
	return cache.expiresOn.After(time.Now())
}

func (cache *simpleCacheExpirable) GetMaxIdleTime() time.Duration {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	return cache.maxIdle
}

func (cache *simpleCacheExpirable) SetMaxIdleTime(maxIdle time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	defer cache.updateLastAccessTime(&cache.lastAccessTime)

	cache.maxIdle = maxIdle
}