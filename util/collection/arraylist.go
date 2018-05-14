/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

type ArrayList struct {
	array []Element
}

func NewArrayList(size, cap int) List {
	return &ArrayList{
		array: make([]Element, size, cap),
	}
}

type arrayListIterator struct {
	array []Element
	len   int
	index int
}

func (lit *arrayListIterator) HasPrev() bool {
	return lit.len > 0 && lit.index >= 0
}

func (lit *arrayListIterator) Prev() Element {
	elem := lit.array[lit.index]
	lit.index--
	return elem
}

func (lit *arrayListIterator) HasNext() bool {
	return lit.index >= 0 && lit.index < lit.len
}

func (lit *arrayListIterator) Next() Element {
	elem := lit.array[lit.index]
	lit.index++
	return elem
}

func (al *ArrayList) GetIterator() Iterator {
	return &arrayListIterator{
		array: al.array[:],
		len:   len(al.array),
		index: 0,
	}
}

func (al *ArrayList) GetPrevIterator() PrevIterator {
	return &arrayListIterator{
		array: al.array[:],
		len:   len(al.array),
		index: len(al.array) - 1,
	}
}

func (al *ArrayList) IsEmpty() bool {
	return len(al.array) < 1
}

func (al *ArrayList) Add(e Element) {
	al.array = append(al.array, e)
}

func (al *ArrayList) AddAll(iter Iterable) {
	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			e := it.Next()
			al.Add(e)
		}
	}
}

func (al *ArrayList) Clear() {
	al.array = al.array[:0]
}

func (al *ArrayList) Contains(e Element) bool {
	return al.ContainsIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (al *ArrayList) ContainsAll(iter Iterable) bool {
	exists := true
	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			elem := it.Next()
			if !al.Contains(elem) {
				exists = false
				break
			}
		}
	}
	return exists
}

func (al *ArrayList) ContainsIf(matchMethod MatchMethod, param interface{}) bool {
	exists := false
	for _, elem := range al.array {
		if matchMethod(elem, param) {
			exists = true
			break
		}
	}
	return exists
}

func (al *ArrayList) Remove(e Element) {
	al.RemoveIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (al *ArrayList) RemoveAll(iter Iterable) {
	if iter != nil {
		al.RemoveIf(func(elem Element, param interface{}) bool {
			if it, ok := param.(Iterable); ok {
				return Contains(it, elem)
			}
			return false
		}, iter)
	}
}

func (al *ArrayList) RemoveIf(matchMethod MatchMethod, param interface{}) {
	temp := make([]Element, 0, len(al.array))
	for _, elem := range al.array {
		if !matchMethod(elem, param) {
			temp = append(temp, elem)
		}
	}
	al.array = temp
}

func (al *ArrayList) Size() int {
	return len(al.array)
}

func (al *ArrayList) AddAt(index int, e Element) {
	count := len(al.array)
	if index >= count {
		// Push to back
		al.array = append(al.array, e)
	} else if index >= 0 {
		temp := make([]Element, count+1)
		copy(temp, al.array[:index])
		temp[index] = e
		copy(temp[index+1:], al.array[index:])
		al.array = temp
	}
}

func (al *ArrayList) AddAllAt(index int, iter Iterable) {
	if iter == nil {
		return
	}

	iterArray := ToArray(iter)
	iterLen := len(iterArray)
	if iterLen <= 0 {
		return
	}

	count := len(al.array)
	if index >= count {
		// Push to back
		al.array = append(al.array, iterArray...)
	} else if index >= 0 {
		temp := make([]Element, count+iterLen)
		copy(temp, al.array[:index])
		copy(temp[index:], iterArray)
		copy(temp[index+iterLen:], al.array[index:])
		al.array = temp
	}
}

func (al *ArrayList) GetAt(index int) (Element, bool) {
	count := len(al.array)
	if index >= 0 && index < count {
		return al.array[index], true
	}
	return nil, false
}

func (al *ArrayList) SetAt(index int, e Element) {
	count := len(al.array)
	if index >= 0 && index < count {
		al.array[index] = e
	}
}

func (al *ArrayList) RemoveAt(index int) {
	count := len(al.array)
	if index >= 0 && index < count {
		var temp []Element
		if index == 0 {
			temp = al.array[1:]
		} else if index == count-1 {
			temp = al.array[:index]
		} else {
			temp = make([]Element, count-1)
			copy(temp, al.array[:index])
			copy(temp[index:], al.array[index+1:])
		}
		al.array = temp
	}
}

func (al *ArrayList) IndexOf(e Element) int {
	return al.IndexIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (al *ArrayList) IndexIf(matchMethod MatchMethod, param interface{}) int {
	index := -1

	for i, elem := range al.array {
		if matchMethod(elem, param) {
			index = i
			break
		}
	}

	return index
}

func (al *ArrayList) LastIndexOf(e Element) int {
	return al.LastIndexIf(func(elem Element, param interface{}) bool {
		return elem == param
	}, e)
}

func (al *ArrayList) LastIndexIf(matchMethod MatchMethod, param interface{}) int {
	index := -1
	count := len(al.array)
	for i := count - 1; i >= 0; i-- {
		elem := al.array[i]
		if matchMethod(elem, param) {
			index = i
			break
		}
	}

	return index
}
