/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var (
	propertiesWriterRegexMatches = regexp.MustCompile(`=|\\|[[:cntrl:]]|[^[:ascii:]]`)
)

// PropertiesWriter implements Writer interface for properties file format
type PropertiesWriter struct {
}

func NewWriter() Writer {
	return &PropertiesWriter{}
}

// Write() writes a Properties object to writer in ini file format.
// Properties.Props will be written as key-value pairs.
func (pw *PropertiesWriter) Write(p Properties, w io.Writer) error {
	var err error = nil
	writer := bufio.NewWriter(w)
	defer writer.Flush()
	keys := p.GetProps()
	for _, key := range keys {
		value, _ := p.GetProp(key)
		if err = pw.writeKeyValue(writer, key, value); err != nil {
			break
		}
	}
	return err
}

// writeKeyValue() writes specified key & value to writer
func (pw *PropertiesWriter) writeKeyValue(w *bufio.Writer, key Key, value Value) error {
	keyText := pw.escape(fmt.Sprintf("%v", key))
	valueText := pw.escape(fmt.Sprintf("%v", value))
	_, err := fmt.Fprintln(w, fmt.Sprintf("%s = %s", keyText, valueText))
	return err
}

// escape() returns ini-escaped text for specified input.
func (pw *PropertiesWriter) escape(text string) string {
	escaped := propertiesWriterRegexMatches.ReplaceAllStringFunc(text, func(s string) string {
		ret := s
		switch s {
		case "\r":
			ret = `\r`
		case "\n":
			ret = `\n`
		case "\t":
			ret = `\t`
		case "\\":
			ret = `\\`
		case "=":
			ret = `\=`
		default:
			runes := []rune(s)
			ret = pw.escapeUnicode(runes[0])
		}
		return ret
	})
	return escaped
}

// escapeUnicode() returns \uxxxx format of specified unicode rune.
func (pw *PropertiesWriter) escapeUnicode(r rune) string {
	return fmt.Sprintf("\\u%04X", r)
}
