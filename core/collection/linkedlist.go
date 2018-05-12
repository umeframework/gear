/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"container/list"
)

type LinkedList struct {
	innerList *list.List
}

func NewLinkedList() List {
	ll := LinkedList{
		innerList: list.New(),
	}
	return &ll
}

type linkedListIterator struct {
	element *list.Element
}

func (iter *linkedListIterator) HasNext() bool {
	return iter.element != nil
}

func (iter *linkedListIterator) Next() Element {
	valule := iter.element.Value
	iter.element = iter.element.Next()
	return valule
}

func (iter *linkedListIterator) HasPrev() bool {
	return iter.element != nil && iter.element.Prev() != nil
}

func (iter *linkedListIterator) Prev() Element {
	valule := iter.element.Value
	iter.element = iter.element.Prev()
	return valule
}

func (ll *LinkedList) GetIterator() Iterator {
	iter := linkedListIterator{
		element: ll.innerList.Front(),
	}
	return &iter
}

func (ll *LinkedList) GetPrevIterator() PrevIterator {
	iter := linkedListIterator{
		element: ll.innerList.Back(),
	}
	return &iter
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.innerList.Len() < 1
}

func (ll *LinkedList) Add(e Element) {
	ll.innerList.PushBack(e)
}

func (ll *LinkedList) AddAll(iter Iterable) {
	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			e := it.Next()
			ll.Add(e)
		}
	}
}

func (ll *LinkedList) Clear() {
	ll.innerList = ll.innerList.Init()
}

func (ll *LinkedList) Contains(e Element) bool {
	return ll.ContainsIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (ll *LinkedList) ContainsAll(iter Iterable) bool {
	exists := true
	for it := iter.GetIterator(); it.HasNext(); {
		elem := it.Next()
		if !ll.Contains(elem) {
			exists = false
			break
		}
	}
	return exists
}

func (ll *LinkedList) ContainsIf(matchMethod MatchMethod, param interface{}) bool {
	exists := false
	for elem := ll.innerList.Front(); elem != nil; elem = elem.Next() {
		if matchMethod(elem.Value, param) {
			exists = true
			break
		}
	}
	return exists
}

func (ll *LinkedList) Remove(e Element) {
	ll.RemoveIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (ll *LinkedList) RemoveAll(iter Iterable) {
	ll.RemoveIf(func(elem Element, param interface{}) bool {
		if it, ok := param.(Iterable); ok {
			return Contains(it, elem)
		}
		return false
	}, iter)
}

func (ll *LinkedList) RemoveIf(matchMethod MatchMethod, param interface{}) {
	for it := ll.innerList.Front(); it != nil; {
		next := it.Next()
		if matchMethod(it.Value, param) {
			ll.innerList.Remove(it)
		}
		it = next
	}
}

func (ll *LinkedList) Size() int {
	return ll.innerList.Len()
}

func (ll *LinkedList) indexToElement(index int) (*list.Element, bool) {
	var elem *list.Element = nil
	var ok = false

	count := ll.innerList.Len()
	if index >= 0 && index < count {
		i := 0
		elem = ll.innerList.Front()
		for ; i < index && elem != nil; {
			i++
			elem = elem.Next()
		}
		ok = true
	}

	return elem, ok
}

func (ll *LinkedList) AddAt(index int, e Element) {
	if index >= ll.innerList.Len() {
		ll.innerList.PushBack(e)
	} else {
		if pos, ok := ll.indexToElement(index); ok {
			ll.innerList.InsertBefore(e, pos)
		}
	}
}

func (ll *LinkedList) AddAllAt(index int, iter Iterable) {
	if index >= ll.innerList.Len() {
		pos := ll.innerList.Back()
		for it := iter.GetIterator(); it.HasNext(); {
			pos = ll.innerList.InsertAfter(it.Next(), pos)
		}
	} else {
		pos, _ := ll.indexToElement(index - 1)
		for it := iter.GetIterator(); it.HasNext(); {
			value := it.Next()
			if pos == nil {
				pos = ll.innerList.PushBack(value)
			} else {
				pos = ll.innerList.InsertAfter(value, pos)
			}
		}
	}
}

func (ll *LinkedList) GetAt(index int) (Element, bool) {
	if pos, ok := ll.indexToElement(index); ok {
		return pos.Value, true
	}
	return nil, false
}

func (ll *LinkedList) SetAt(index int, e Element) {
	if pos, ok := ll.indexToElement(index); ok {
		pos.Value = e
	}
}

func (ll *LinkedList) IndexOf(e Element) int {
	return ll.IndexIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (ll *LinkedList) IndexIf(matchMethod MatchMethod, param interface{}) int {
	index := -1
	for i, it := 0, ll.innerList.Front(); it != nil; {
		if matchMethod(it.Value, param) {
			index = i
			break
		}
		i++
		it = it.Next()
	}
	return index
}

func (ll *LinkedList) LastIndexOf(e Element) int {
	return ll.LastIndexIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (ll *LinkedList) LastIndexIf(matchMethod MatchMethod, param interface{}) int {
	index := -1
	for i, it := ll.innerList.Len()-1, ll.innerList.Back(); it != nil; {
		if matchMethod(it.Value, param) {
			index = i
			break
		}
		i--
		it = it.Prev()
	}
	return index
}
