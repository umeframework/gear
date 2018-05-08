/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"sort"
	"sync"
)

type simpleCacheManager struct {
	dict  map[string]Cache
	mutex sync.RWMutex
}

func NewSimpleCacheManager() CacheManager {
	return &simpleCacheManager{
		dict:  make(map[string]Cache),
		mutex: sync.RWMutex{},
	}
}

func (cacheMgr *simpleCacheManager) AddCache(cache Cache) error {
	cacheMgr.mutex.Lock()
	defer cacheMgr.mutex.Unlock()

	cacheMgr.dict[cache.GetName()] = cache
	return nil
}

func (cacheMgr *simpleCacheManager) GetCache(name string) (Cache, bool, error) {
	cacheMgr.mutex.RLock()
	defer cacheMgr.mutex.RUnlock()

	cache, ok := cacheMgr.dict[name]
	return cache, ok, nil
}

func (cacheMgr *simpleCacheManager) RemoveCache(name string) error {
	cacheMgr.mutex.Lock()
	defer cacheMgr.mutex.Unlock()

	delete(cacheMgr.dict, name)
	return nil
}

func (cacheMgr *simpleCacheManager) ClearCaches() error {
	cacheMgr.mutex.Lock()
	defer cacheMgr.mutex.Unlock()

	cacheMgr.dict = make(map[string]Cache)
	return nil
}

func (cacheMgr *simpleCacheManager) GetCacheNames() []string {
	cacheMgr.mutex.RLock()
	defer cacheMgr.mutex.RUnlock()

	count := len(cacheMgr.dict)
	names := make([]string, count)
	index := 0
	for key, _ := range cacheMgr.dict {
		names[index] = key
		index++
	}

	sort.Strings(names)
	return names
}
