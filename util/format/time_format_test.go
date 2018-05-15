/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package format

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	tf := TimeFormat{}
	tm := time.Date(2006, 1, 2, 15, 4, 5, 987654321, time.Local)
	//t.Logf("target time: %v", tm)

	inputs := []interface{} {
		tm, &tm,
	}

	patterns := []string {
		`yyyy/MM/dd hh:mm:ss (AM)`,
		`yy/M/d HH:m:s.SSSSSSSSS Z07 (am)`,
	}
	expecteds := []string {
		`2006/01/02 03:04:05 (PM)`,
		`06/1/2 15:4:5.987654321 +08 (pm)`,
	}
	converteds := []time.Time{
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		tm,
	}

	for i, pattern := range patterns {
		expected := expecteds[i]
		tf.Pattern = pattern
		if result, err := tf.Format(inputs[i % 2]); err != nil || result != expected {
			t.Errorf("failed to format for pattern %v. expected: %v, result: %v, err: %v",
				pattern, expected, result, err)
		} else {
			if obj, err := tf.Parse(result); err != nil {
				t.Errorf("failed to parse %v, err: %v", result, err)
			} else if tm2, ok := obj.(time.Time); tm != converteds[i % 2] || !ok {
				t.Errorf("parse result for %v is not correct. expected: %v, result: %v", result, converteds[i % 2], tm2)
			}
		}
	}
}