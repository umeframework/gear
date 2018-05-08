/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSimpleCache(t *testing.T) {
	name := "simple-test"
	content := map[Key]Value{
		"id":  rand.Int(),
		"now": fmt.Sprintf("%v", time.Now()),
	}
	count := rand.Intn(100)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("no. %d", i)
		value := i
		content[key] = value
	}

	// Create cache
	cache := NewSimpleCache(name)
	for key, value := range content {
		cache.Put(key, value)
	}

	// Check name
	if nameGot := cache.GetName(); nameGot != name {
		t.Errorf("failed to get name. expected: %v, real: %v", name, nameGot)
	}

	// Check for keys & values
	keys := cache.GetKeys()
	if len(keys) != len(content) {
		t.Errorf("count of keys are differenct. expected = %v, real = %v", len(content), len(keys))
	}
	for _, key := range keys {
		valueExpected, ok := content[key]
		if !ok {
			t.Errorf("invalid key retrieved: %v", key)
		}
		valueGot, ok, err := cache.Get(key)
		if !ok || err != nil || valueGot != valueExpected {
			t.Errorf("invalid value retrieved (sync). key = %v, value expected = %v, value got = %v, ok = %v, err = %v",
				key, valueExpected, valueGot, ok, err)
		}
	}

	// Check for async reading
	wg := sync.WaitGroup{}
	wg.Add(len(content))
	for key, value := range content {
		go func(k Key, v Value) {
			if valueGot, ok, err := cache.Get(k); (!ok) || err != nil || valueGot != v {
				t.Errorf("invalid value retrieved (async). key = %v, value expected = %v, value got = %v, ok = %v, err = %v",
					k, v, valueGot, ok, err)
			}
			wg.Done()
		}(key, value)
	}
	wg.Wait()

	// Check for async writing & deleting
	for _, key := range keys {
		wg.Add(1)
		go func(k Key) {
			dice := rand.Int()
			switch dice % 10 {
			case 0: // Remove value
				cache.Remove(k)
			case 1, 2: // Update value
				cache.Put(k, fmt.Sprintf("no. %v (%v)", k, dice))
			case 3:
				cache.Compute(k, rand.Int(), func(key Key, value Value, exists bool, err error, param Param) (Value, bool) {
					if exists {
						value = fmt.Sprintf("%v computed (%v)", value, param)
					}
					return value, exists
				})
			default:
				cache.Get(k)
			}

			wg.Done()
		}(key)
	}
	wg.Wait()

	// Clear cache
	cache.Clear()
}
