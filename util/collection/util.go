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
	for it := iter.GetIterator(); it.HasNext(); {
		if matchMethod(it.Next(), param) {
			found = true
			break
		}
	}
	return found
}

func ToArray(iter Iterable) []Element {
	array := make([]Element, 0, 0x10)
	for it := iter.GetIterator(); it.HasNext(); {
		e := it.Next()
		array = append(array, e)
	}
	return array
}

func FromArray(elements ...Element) List {
	return &ArrayList{
		array: elements[:],
	}
}

func AddArray(list List, elements ...Element) {
	if list != nil {
		elementList := FromArray(elements...)
		list.AddAll(elementList)
	}
}
