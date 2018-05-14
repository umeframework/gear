/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

import "testing"

func TestStack(t *testing.T) {
	stk := NewStack()

	if !stk.IsEmpty() {
		t.Errorf("initial status is not empty")
	}
	if size := stk.Size(); size != 0 {
		t.Errorf("initial size is not zero. real: %v", size)
	}
	if e, ok := stk.Peek(); ok {
		t.Errorf("initial peek not correct. value: %v, ok: %v", e, ok)
	}

	i := 100
	stk.Push(i)
	if stk.IsEmpty() {
		t.Errorf("status is empty")
	}
	if size := stk.Size(); size != 1 {
		t.Errorf("size is not correct. expected: %v, real: %v", 1, size)
	}
	if e, ok := stk.Peek(); e != i || !ok {
		t.Errorf("peek not correct. expected: %v, real: %v, ok: %v", i, e, ok)
	}
	if e, ok := stk.Pop(); e != i || !ok {
		t.Errorf("poll not correct. expected: %v, real: %v, ok: %v", i, e, ok)
	}
	if e, ok := stk.Peek(); ok {
		t.Errorf("peek not correct. expected: %v, real: %v, ok: %v", nil, e, ok)
	}
	if e, ok := stk.Pop(); ok {
		t.Errorf("poll not correct. expected: %v, real: %v, ok: %v", nil, e, ok)
	}
}

func TestStack_2(t *testing.T) {
	stk := NewStack()

	count := 5
	array := make([]Element, count)
	for i := 0; i < count; i++ {
		array[i] = i
	}
	iter := FromArray(array...)

	stk.AddAll(iter)
	//walk(t, stk)
	//t.Log("")

	for i, it := 0, stk.GetIterator(); it.HasNext(); {
		if e := it.Next(); e != count-i-1 {
			t.Errorf("failed to get item from iterator. expected: %v, real: %v", count-i-1, e)
		}
		i++
	}

	for i := 0; i < count; i++ {
		if !stk.Contains(i) {
			t.Errorf("item not found by mistake. %v", i)
		}
	}
	if stk.Contains("0") {
		t.Errorf("item found by mistake. %v", "0")
	}
	if stk.Contains(-1) {
		t.Errorf("item found by mistake. %v", -1)
	}
	if stk.Contains(count) {
		t.Errorf("item found by mistake. %v", count)
	}
	if !stk.ContainsAll(iter) {
		t.Errorf("iter not found by mistake")
	}

	stk.Remove(count - 1)
	//walk(t, stk)
	//t.Log("")

	if stk.Contains(count - 1) {
		t.Errorf("item found by mistake. %v", count-1)
	}
	if stk.ContainsAll(iter) {
		t.Errorf("iter found by mistake")
	}

	for i := 0; i < count-1; i++ {
		if e, ok := stk.Peek(); e != count-2 || !ok {
			t.Errorf("failed to peek item. expected: %v, real: %v", count-2, e)
		}
	}
	for i := 0; i < count-1; i++ {
		if e, ok := stk.Pop(); e != count-i-2 || !ok {
			t.Errorf("failed to poll item. expected: %v, real: %v", count-i-2, e)
		}
	}
}
