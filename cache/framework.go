/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package cache

type Key interface{}
type Value interface{}
type Param interface{}

type Cache interface {
	GetName() string
	GetKeys() []Key
	Get(key Key) (Value, bool, error)
	Put(key Key, value Value) error
	Compute(key Key, param Param, callback func(key Key, value Value, exists bool, err error, param Param) (Value, bool)) (Value, error)
	Remove(key Key) error
	Clear() error
}

type CacheManager interface {
	AddCache(cache Cache) error
	GetCache(name string) (Cache, bool, error)
	RemoveCache(name string) error
	ClearCaches() error
	GetCacheNames() []string
}
