/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
)

func TestCacheManager(t *testing.T) {
	cacheMgr := NewSimpleCacheManager()

	// Add caches
	wg := sync.WaitGroup{}
	count := rand.Intn(100) + 5
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(n int) {
			name := fmt.Sprintf("cache no. %04d", n)
			cache := NewSimpleCache(name)
			cacheMgr.AddCache(cache)

			if rand.Int()%3 == 0 {
				wg.Add(1)
				go func(index int) {
					name2 := fmt.Sprintf("cache no. %04d", index)
					if cache, ok, err := cacheMgr.GetCache(name2); (!ok) || err != nil {
						t.Errorf("failed to get cache, name = %v, ok = %v, err = %v", name2, ok, err)
					} else if cache.GetName() != name2 {
						t.Errorf("invalid cache name. expected: %v, real: %v", name2, cache.GetName())
					}
					wg.Done()
				}(n)
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

	// Check caches
	cacheNames := cacheMgr.GetCacheNames()
	sort.Strings(cacheNames)
	if len(cacheNames) != count {
		t.Errorf("invalid cache count")
	} else {
		for i := 0; i < count; i++ {
			expected := fmt.Sprintf("cache no. %04d", i)
			real := cacheNames[i]
			if expected != real {
				t.Errorf("failed to get cache names. expected: %v, real: %v", expected, real)
			}
		}
	}

	// Remove cache
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			switch rand.Int() % 20 {
			case 1:
				cacheMgr.ClearCaches()
			default:
				index := rand.Intn(count)
				name := fmt.Sprintf("cache no. %04d", index)
				cacheMgr.RemoveCache(name)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
