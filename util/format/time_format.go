/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package format

import (
	"regexp"
	"errors"
	"time"
	"strings"
)

var (
	TimeFormatRegex = regexp.MustCompile(`(y+|M+|d+|H+|h+|m+|S+|s+|AM|am)`)
	TimeFormatPatternDict = map[string]string {
		"yy": "06",
		"yyyy": "2006",
		"M": "1",
		"MM": "01",
		"d": "2",
		"dd": "02",
		"h": "3",
		"hh": "03",
		"H": "15",
		"HH": "15",
		"m": "4",
		"mm": "04",
		"s": "5",
		"ss": "05",
		"am": "pm",
		"AM": "PM",
	}
	ErrorInvalidFormatInput = errors.New("invalid object to format, only time is supported")
	ErrorInvalidParseInput  = errors.New("invalid object to format, only time is supported")
)

type TimeFormat struct {
	Pattern string
}

func (tf *TimeFormat) Format(object interface{}) (string, error) {
	var text string
	var err error

	// Check input type
	var tm *time.Time = nil
	if t, ok := object.(time.Time); ok {
		tm = &t
	}
	if t, ok := object.(*time.Time); ok {
		tm = t
	}

	// Format time
	if tm == nil {
		err = ErrorInvalidFormatInput
	} else {
		pattern := tf.normalizePattern(tf.Pattern)
		text = tm.Format(pattern)
	}
	return text, err
}

func (tf *TimeFormat) normalizePattern(pattern string) string {
	ret := TimeFormatRegex.ReplaceAllStringFunc(pattern, func(s string) string {
		s = strings.Replace(s, "S", "0", -1)
		if r, ok := TimeFormatPatternDict[s]; ok {
			return r
		}
		return s
	})
	return ret
}

func (tf *TimeFormat) Parse(text string) (interface{}, error) {
	pattern := tf.normalizePattern(tf.Pattern)
	tm, err := time.Parse(pattern, text)
	return tm, err
}
