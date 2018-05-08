/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

func ToProperties(target Properties, source map[Key]Value) {
	if target != nil && source != nil {
		for key, value := range source {
			target.SetProp(key, value)
		}
	}
}

func ToMap(target map[Key]Value, source Properties) {
	if target != nil && source != nil {
		keys := source.GetProps()
		for _, key := range keys {
			if value, ok := source.GetProp(key); ok {
				target[key] = value
			} else {
				delete(target, key)
			}
		}
	}
}
