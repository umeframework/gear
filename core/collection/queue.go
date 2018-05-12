/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

type queue struct {
	list List
}

func NewQueue() Queue {
	return &queue{
		list: NewLinkedList(),
	}
}

func (q *queue) GetIterator() Iterator {
	return q.list.GetIterator()
}

func (q *queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *queue) Add(e Element) {
	q.list.Add(e)
}

func (q *queue) AddAll(iter Iterable) {
	q.list.AddAll(iter)
}

func (q *queue) Clear() {
	q.list.Clear()
}

func (q *queue) Contains(e Element) bool {
	return q.list.Contains(e)
}

func (q *queue) ContainsAll(iter Iterable) bool {
	return q.list.ContainsAll(iter)
}

func (q *queue) ContainsIf(matchMethod MatchMethod, param interface{}) bool {
	return q.list.ContainsIf(matchMethod, param)
}

func (q *queue) Remove(e Element) {
	q.list.Remove(e)
}

func (q *queue) RemoveAll(iter Iterable) {
	q.list.RemoveAll(iter)
}

func (q *queue) RemoveIf(matchMethod MatchMethod, param interface{}) {
	q.list.RemoveIf(matchMethod, param)
}

func (q *queue) Size() int {
	return q.list.Size()
}

func (q *queue) Push(e Element) {
	q.list.Add(e)
}

func (q *queue) Peek() (Element, bool) {
	return q.list.GetAt(0)
}

func (q *queue) Poll() (Element, bool) {
	elem, ok := q.list.GetAt(0)
	if ok {
		q.list.RemoveAt(0)
	}
	return elem, ok
}
