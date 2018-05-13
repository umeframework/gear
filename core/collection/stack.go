/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

type stack struct {
	list List
}

func NewStack() Stack {
	return &stack{
		list: NewLinkedList(),
	}
}

func (stk *stack) GetIterator() Iterator {
	return stk.list.GetIterator()
}

func (stk *stack) IsEmpty() bool {
	return stk.list.IsEmpty()
}

func (stk *stack) Add(e Element) {
	stk.list.AddAt(0, e)
}

func (stk *stack) AddAll(iter Iterable) {
	if iter != nil {
		for it := iter.GetIterator(); it.HasNext(); {
			e := it.Next()
			stk.Add(e)
		}
	}
}

func (stk *stack) Clear() {
	stk.list.Clear()
}

func (stk *stack) Contains(e Element) bool {
	return stk.list.Contains(e)
}

func (stk *stack) ContainsAll(iter Iterable) bool {
	return stk.list.ContainsAll(iter)
}

func (stk *stack) ContainsIf(matchMethod MatchMethod, param interface{}) bool {
	return stk.list.ContainsIf(matchMethod, param)
}

func (stk *stack) Remove(e Element) {
	stk.list.Remove(e)
}

func (stk *stack) RemoveAll(iter Iterable) {
	stk.list.RemoveAll(iter)
}

func (stk *stack) RemoveIf(matchMethod MatchMethod, param interface{}) {
	stk.list.RemoveIf(matchMethod, param)
}

func (stk *stack) Size() int {
	return stk.list.Size()
}

func (stk *stack) Push(e Element) {
	stk.list.AddAt(0, e)
}

func (stk *stack) Peek() (Element, bool) {
	return stk.list.GetAt(0)
}

func (stk *stack) Pop() (Element, bool) {
	elem, ok := stk.list.GetAt(0)
	stk.list.RemoveAt(0)
	return elem, ok
}
