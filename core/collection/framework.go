/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

// Element represents for any element used in collection
type Element interface{}

// MatchMethod matches specified element
type MatchMethod func(e Element, param interface{}) bool

// Iterator represents for an iterator over collection.
type Iterator interface {
	HasNext() bool
	Next() Element
}

// PrevIterator represent for an iterator over collection in reserved order.
type PrevIterator interface {
	HasPrev() bool
	Prev() Element
}

// Iterable allows an object to be accessed via iterator
type Iterable interface {
	GetIterator() Iterator
}

// Iterable allows an object to be accessed via reversed iterator
type PrevIterable interface {
	GetPrevIterator() PrevIterator
}

// Collection represents a group of objects.
type Collection interface {
	Iterable

	IsEmpty() bool
	Add(e Element)
	AddAll(iter Iterable)

	Clear()

	Contains(e Element) bool
	ContainsAll(iter Iterable) bool
	ContainsIf(matchMethod MatchMethod, param interface{}) bool

	Remove(e Element)
	RemoveAll(iter Iterable)
	RemoveIf(matchMethod MatchMethod, param interface{})

	Size() int
}

// List represents a sequenced collection.
type List interface {
	Collection

	AddAt(index int, e Element)
	AddAllAt(index int, iter Iterable)

	GetAt(index int) (Element, bool)
	SetAt(index int, e Element)

	RemoveAt(index int)

	IndexOf(e Element) int
	IndexIf(matchMethod MatchMethod, param interface{}) int
	LastIndexOf(e Element) int
	LastIndexIf(matchMethod MatchMethod, param interface{}) int
}

// Queue represents a FIFO collection.
type Queue interface {
	Collection

	Push(e Element)
	Peek() (Element, bool)
	Poll() (Element, bool)
}

// Stack represents a FILO collection.
type Stack interface {
	Collection

	Push(e Element)
	Peek() (Element, bool)
	Pop() (Element, bool)
}
