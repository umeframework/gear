/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if size := ll.Size(); size != 0 {
		t.Errorf("initial size is not correct. expected: %v, real: %v", 0, size)
	}
	if !ll.IsEmpty() {
		t.Errorf("initial empty status is not correct")
	}

	count := 10
	for i := 0; i < count; i++ {
		ll.Add(i)
	}
	if size := ll.Size(); size != count {
		t.Errorf("initial size is not correct. expected: %v, real: %v", count, size)
	}
	if ll.IsEmpty() {
		t.Errorf("empty status is not correct")
	}
	for i := 0; i < count; i++ {
		if e, ok := ll.GetAt(i); e != i || !ok {
			t.Errorf("failed to get item at %v. expected: %v, real: %v", i, i, e)
		}
	}
	if e, ok := ll.GetAt(-1); ok {
		t.Errorf("failed to get item at %v. expected: nil, real: %v, ok: %v",
			-1, e, ok)
	}
	if e, ok := ll.GetAt(count); ok {
		t.Errorf("failed to get item at %v. expected: nil, real: %v, ok: %v",
			count, e, ok)
	}

	for i := 0; i < count; i++ {
		if !ll.Contains(i) {
			t.Errorf("failed to find contains %v", i)
		}
		if ll.Contains(fmt.Sprintf("%v", i)) {
			t.Errorf("failed to find contains string %v", i)
		}
	}
	if ll.Contains(-1) {
		t.Errorf("failed to find contains %v (should not exist)", -1)
	}
	if ll.Contains(count) {
		t.Errorf("failed to find contains %v (should not exist)", count)
	}

	for i, it := 0, ll.GetIterator(); it.HasNext(); {
		if e := it.Next(); e != i {
			t.Errorf("value of iterator at index %v is not correct. expected: %v, real: %v",
				i, i, e)
		}
		i++
	}

	{
		max := 0
		for i, it := 0, ll.GetIterator(); it.HasNext(); {
			if e := it.Next(); e != i {
				t.Errorf("value of iterator at index %v is not correct. expected: %v, real: %v",
					i, i, e)
			}
			i++
			max = i
		}
		if max != count {
			t.Errorf("max looped not correct. expected: %v, real: %v", count, max)
		}
	}

	if prevIterator, ok := ll.(PrevIterable); !ok {
		t.Errorf("linked list should also be a PrevIterable")
	} else {
		min := count
		for i, it := count-1, prevIterator.GetPrevIterator(); it.HasPrev(); {
			//t.Logf("no. %v", i)
			if e := it.Prev(); e != i {
				t.Errorf("value of iterator at index %v is not correct. expected: %v, real: %v",
					i, i, e)
			}
			i--
			min = i
		}
		if min != -1 {
			t.Errorf("min looped not correct. expected: %v, real: %v", -1, min)
		}
	}

	ll.Clear()
	if !ll.IsEmpty() {
		t.Errorf("failed to clear linked list")
	}
}

func TestLinkedList_2(t *testing.T) {
	ll := NewLinkedList()
	count := 5
	for i := 0; i < count; i++ {
		ll.AddAt(0, i)
	}
	for i, it := count-1, ll.GetIterator(); it.HasNext(); {
		if elem := it.Next(); elem != i {
			t.Errorf("failed to get elem at %v. expected: %v, real: %v", i, count-i-1, elem)
		}
		i--
	}

	ll2 := NewLinkedList()
	ll2.AddAll(ll)
	for i, it := 0, ll2.GetIterator(); i < count; i++ {
		if elem := it.Next(); elem != count-i-1 {
			t.Errorf("failed to get elem at %v. expected: %v, real: %v", i, count-i-1, elem)
		}
	}
	//walk(t, ll2)
	//t.Log("")

	ll2.AddAllAt(count, ll)
	for i := 0; i < ll2.Size(); i++ {
		expected := count - i%count - 1
		if elem, ok := ll2.GetAt(i); elem != expected || !ok {
			t.Errorf("value not correct at index %v. expected: %v, real: %v", i, expected, elem)
		}
	}
	//walk(t, ll2)
	//t.Log("")

	ll2.AddAllAt(count, ll)
	for i := 0; i < ll2.Size(); i++ {
		expected := count - i%count - 1
		if elem, ok := ll2.GetAt(i); elem != expected || !ok {
			t.Errorf("value not correct at index %v. expected: %v, real: %v", i, expected, elem)
		}
	}
	//walk(t, ll2)
	//t.Log("")

	ll2.AddAllAt(0, ll)
	for i := 0; i < ll2.Size(); i++ {
		expected := count - i%count - 1
		if elem, ok := ll2.GetAt(i); elem != expected || !ok {
			t.Errorf("value not correct at index %v. expected: %v, real: %v", i, expected, elem)
		}
	}
	//walk(t, ll2)
	//t.Log("")

	if ok := ll2.ContainsAll(ll); !ok {
		t.Errorf("contains all failed. ll2 should contains ll")
	}

	ll2.Remove(0)
	//walk(t, ll2)
	//t.Log("")
	if ok := ll2.ContainsAll(ll); ok {
		t.Errorf("contains all failed. ll2 should NOT contains ll")
	}

	ll2.RemoveIf(func(e Element, param interface{}) bool {
		if i, ok := e.(int); ok {
			return i%2 == 0
		}
		return false
	}, nil)
	//walk(t, ll2)
	//t.Log("")

	ll2.RemoveAll(ll)
	if !ll2.IsEmpty() {
		t.Errorf("failed to remove ll from ll2")
	}
	//walk(t, ll2)
	//t.Log("")

}

func TestLinkedList_3(t *testing.T) {
	ll := NewLinkedList()
	count := 5
	for i := 0; i < count; i++ {
		ll.Add(i)
	}
	for i := 0; i < count; i++ {
		ll.Add(i)
	}

	for i := 0; i < count; i++ {
		if index := ll.IndexOf(i); index != i {
			t.Errorf("failed to find index for %v. expected: %v, real: %v", i, i, index)
		}
	}
	for i := 0; i < count; i++ {
		if index := ll.LastIndexOf(i); index != i+count {
			t.Errorf("failed to find index for %v. expected: %v, real: %v", i, i+count, index)
		}
	}

	if index := ll.IndexOf(-1); index != -1 {
		t.Errorf("failed to find index for %v. expected: %v, real: %v", -1, -1, index)
	}
	if index := ll.IndexOf(count + 1); index != -1 {
		t.Errorf("failed to find index for %v. expected: %v, real: %v", count+1, -1, index)
	}

	count = ll.Size()
	for i := 0; i < count; i++ {
		ll.SetAt(i, i)
	}
	for i := 0; i < count; i++ {
		if e, ok := ll.GetAt(i); e != i || !ok {
			t.Errorf("failed to get value at index %v. expected: %v, real: %v", i, i, e)
		}
	}
}