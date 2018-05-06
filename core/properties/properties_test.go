/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

import (
	"math/rand"
	"testing"
	"time"
)

func TestPropertiesImpl_Props(t *testing.T) {
	p := New()

	// Get props for initial status
	if props := p.GetProps(); len(props) != 0 {
		t.Errorf("props count for initial value is not correct. real: %v", props)
	}

	// Add some values
	testDict := map[Key]Value{
		100:    "tom",
		'A':    time.Now(),
		"name": "hello, world!",
	}
	for key, value := range testDict {
		p.SetProp(key, value)
	}

	// Test whether values are correctly set
	if props := p.GetProps(); len(props) != len(testDict) {
		t.Errorf("props count is not correct. expected: %v, real: %v", len(testDict), len(props))
	} else {
		for _, key := range props {
			if value, ok := p.GetProp(key); !ok {
				t.Errorf("failed to retrive prop for key: %v", key)
			} else {
				value2 := testDict[key]
				if value != value2 {
					t.Errorf("prop value is not correct. key = %v, expected value = %v, real value = %v",
						key, value2, value)
				}
			}
		}
	}

	// Remove some props
	keysToRemove := make([]Key, 0, len(testDict))
	for key := range testDict {
		if rand.Int()%2 == 0 {
			keysToRemove = append(keysToRemove, key)
		}
	}
	for _, key := range keysToRemove {
		p.RemoveProp(key)
	}
	for _, key := range keysToRemove {
		if _, ok := p.GetProp(key); ok {
			t.Errorf("failed to remove key: %v", key)
		}
	}

	// Clear all props
	p.ClearProps()
	if props := p.GetProps(); len(props) != 0 {
		t.Errorf("failed to clear props")
	}
}

