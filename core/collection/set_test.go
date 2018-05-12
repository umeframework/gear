/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package collection

import "testing"

func TestSet(t *testing.T) {
	s := NewSet()

	if !s.IsEmpty() {
		t.Errorf("initial status is not empty")
	}
	if size := s.Size(); size != 0 {
		t.Errorf("initial size is not zero. real: %v", size)
	}

	i := 100
	s.Add(i)
	if s.IsEmpty() {
		t.Errorf("status is empty")
	}
	if size := s.Size(); size != 1 {
		t.Errorf("size is not correct. expected: %v, real: %v", 1, size)
	}
	if !s.Contains(i) {
		t.Errorf("item not found by mistake. %v", i)
	}
}

func TestSet_2(t *testing.T) {
	s := NewSet()

	count := 5
	for i := 0; i < count; i++ {
		s.Add(i)
	}
	if size := s.Size(); size != count {
		t.Errorf("size not correct. expected: %v, real: %v", count, size)
	}
	for i := 0; i < count; i++ {
		if !s.Contains(i) {
			t.Errorf("items not found by mistake. %v", i)
		}
	}
	walk(t, s)

	for i := 0; i < count; i++ {
		s.Add(i)
	}
	if size := s.Size(); size != count {
		t.Errorf("size not correct. expected: %v, real: %v", count, size)
	}
	s.Add(count)
	if size := s.Size(); size != count+1 {
		t.Errorf("size not correct. expected: %v, real: %v", count+1, size)
	}
	if !s.Contains(count) {
		t.Errorf("items not found by mistake. %v", count)
	}
}
