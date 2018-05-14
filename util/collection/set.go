/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

type set struct {
	dict  map[Element]uint8
	items []Element
	dirty bool
}

func NewSet() Set {
	return &set{
		dict:  make(map[Element]uint8),
		items: nil,
		dirty: true,
	}
}

type setIterator struct {
	array []Element
	len   int
	index int
}

func (sit *setIterator) HasNext() bool {
	return sit.index >= 0 && sit.index < sit.len
}

func (sit *setIterator) Next() Element {
	elem := sit.array[sit.index]
	sit.index++
	return elem
}

func (s *set) getItemArray() []Element {
	array := s.items
	if s.dirty {
		count := len(s.dict)
		array = make([]Element, count)
		index := 0
		for key, _ := range s.dict {
			array[index] = key
			index++
		}
		s.items = array
		s.dirty = false
	}
	return array
}

func (s *set) GetIterator() Iterator {
	array := s.getItemArray()
	return &setIterator{
		array: array,
		len:   len(array),
		index: 0,
	}
}

func (s *set) IsEmpty() bool {
	return len(s.dict) <= 0
}

func (s *set) Add(e Element) {
	s.dict[e] = 0
	s.dirty = true
}

func (s *set) AddAll(iter Iterable) {
	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			s.Add(it.Next())
		}
		s.dirty = true
	}
}

func (s *set) Clear() {
	s.dict = make(map[Element]uint8)
	s.dirty = true
}

func (s *set) Contains(e Element) bool {
	_, exists := s.dict[e]
	return exists
}

func (s *set) ContainsAll(iter Iterable) bool {
	exists := true

	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			e := it.Next()
			if !s.Contains(e) {
				exists = false
				break
			}
		}
	}

	return exists
}

func (s *set) ContainsIf(matchMethod MatchMethod, param interface{}) bool {
	exists := false
	for key, _ := range s.dict {
		if matchMethod(key, param) {
			exists = true
			break
		}
	}
	return exists
}

func (s *set) Remove(e Element) {
	delete(s.dict, e)
	s.dirty = true
}

func (s *set) RemoveAll(iter Iterable) {
	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			s.Remove(it.Next())
			s.dirty = true
		}
	}
}

func (s *set) RemoveIf(matchMethod MatchMethod, param interface{}) {
	// Find items matched
	list := NewLinkedList()
	for key, _ := range s.dict {
		if matchMethod(key, param) {
			list.Add(key)
		}
	}

	// Remove items matched
	for it := list.GetIterator(); it.HasNext(); {
		s.Remove(it.Next())
	}

	if !list.IsEmpty() {
		s.dirty = true
	}
}

func (s *set) Size() int {
	return len(s.dict)
}
