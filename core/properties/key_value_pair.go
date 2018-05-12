/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

type keyValuePairImpl struct {
	key   Key
	value Value
}

func NewKeyValuePair() KeyValuePair {
	return &keyValuePairImpl{}
}

func (kv *keyValuePairImpl) GetKey() Key {
	return kv.key
}

func (kv *keyValuePairImpl) SetKey(key Key) {
	kv.key = key
}

func (kv *keyValuePairImpl) GetValue() Value {
	return kv.value
}

func (kv *keyValuePairImpl) SetValue(value Value) {
	kv.value = value
}
