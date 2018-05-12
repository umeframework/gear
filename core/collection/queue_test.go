/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Errorf("initial status is not empty")
	}
	if size := q.Size(); size != 0 {
		t.Errorf("initial size is not zero. real: %v", size)
	}
	if e, ok := q.Peek(); ok {
		t.Errorf("initial peek not correct. value: %v, ok: %v", e, ok)
	}

	i := 100
	q.Add(i)
	if q.IsEmpty() {
		t.Errorf("status is empty")
	}
	if size := q.Size(); size != 1 {
		t.Errorf("size is not correct. expected: %v, real: %v", 1, size)
	}
	if e, ok := q.Peek(); e != i || !ok {
		t.Errorf("peek not correct. expected: %v, real: %v, ok: %v", i, e, ok)
	}
	if e, ok := q.Poll(); e != i || !ok {
		t.Errorf("poll not correct. expected: %v, real: %v, ok: %v", i, e, ok)
	}
	if e, ok := q.Peek(); ok {
		t.Errorf("peek not correct. expected: %v, real: %v, ok: %v", nil, e, ok)
	}
	if e, ok := q.Poll(); ok {
		t.Errorf("poll not correct. expected: %v, real: %v, ok: %v", nil, e, ok)
	}
}

func TestQueue_2(t *testing.T) {
	q := NewQueue()

	count := 5
	array := make([]Element, count)
	for i := 0; i < count; i++ {
		array[i] = i
	}
	iter := FromArray(array...)

	q.AddAll(iter)
	for i, it := 0, q.GetIterator(); it.HasNext(); {
		if e := it.Next(); e != i {
			t.Errorf("failed to get item from iterator. expected: %v, real: %v", i, e)
		}
		i++
	}

	for i := 0; i < count; i++ {
		if !q.Contains(i) {
			t.Errorf("item not found by mistake. %v", i)
		}
	}
	if q.Contains("0") {
		t.Errorf("item found by mistake. %v", "0")
	}
	if q.Contains(-1) {
		t.Errorf("item found by mistake. %v", -1)
	}
	if q.Contains(count) {
		t.Errorf("item found by mistake. %v", count)
	}
	if !q.ContainsAll(iter) {
		t.Errorf("iter not found by mistake")
	}

	q.Remove(count - 1)
	if q.Contains(count - 1) {
		t.Errorf("item found by mistake. %v", count - 1)
	}
	if q.ContainsAll(iter) {
		t.Errorf("iter found by mistake")
	}

	for i := 0; i < count - 1; i++ {
		if e, ok := q.Peek(); e != 0 || !ok {
			t.Errorf("failed to peek item. expected: %v, real: %v", 0, e)
		}
	}
	for i := 0; i < count - 1; i++ {
		if e, ok := q.Poll(); e != i || !ok {
			t.Errorf("failed to poll item. expected: %v, real: %v", i, e)
		}
	}
}
