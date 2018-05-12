/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

func Contains(iter Iterable, e Element) bool {
	return ContainsIf(iter, func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func ContainsIf(iter Iterable, matchMethod MatchMethod, param interface{}) bool {
	found := false
	if col, ok := iter.(Collection); ok {
		found = col.ContainsIf(matchMethod, param)
	} else {
		for it := iter.GetIterator(); it.HasNext(); {
			if matchMethod(it.Next(), param) {
				found = true
				break
			}
		}
	}
	return found
}


