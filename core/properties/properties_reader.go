/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	propertiesReaderRegexComment        = regexp.MustCompile(`^[\s]*[#!]+`)
	propertiesReaderRegexContinuousLine = regexp.MustCompile(`[\\]+$`)
	propertiesReaderRegexSept           = regexp.MustCompile(`([\\]*)=`)
	propertiesReaderReplace             = regexp.MustCompile(`\\0|\\\\|\\=|\\[a-tv-z]|\\u[\w\W]{0,4}`)
)

// PropertiesReader implements Reader interface for properties file format
type PropertiesReader struct {
}

func NewReader() Reader {
	return &PropertiesReader{}
}

// Read() reads (deserialize) a Properties object from reader in properties file format.
// It reads each valid line to props.
func (pr *PropertiesReader) Read(p Properties, r io.Reader) error {
	reader := bufio.NewReader(r)

	var line string
	var err error = nil
	var key string
	var value *string = nil

	for loop := true; loop; {
		// Read a line
		if line, err = pr.readLine(reader); err != nil {
			loop = false
			if err != io.EOF {
				break
			}
			err = nil
		}

		if len(line) < 1 {
			continue
		}

		// Parse key & value
		if key, value, err = pr.parseKeyValue(line); err == nil {
			if value != nil {
				p.SetProp(key, *value)
			} else {
				p.SetProp(key, nil)
			}
		} else {
			break
		}
	}
	return err
}

// readLine() reads a valid line from reader.
// It ignores comments and blank lines, and combines continuous lines linked with ending \.
func (pr *PropertiesReader) readLine(r *bufio.Reader) (string, error) {
	var text string = ""
	var err error = nil
	var line string

	for loop := true; loop; {
		// Read a line
		if line, err = r.ReadString('\n'); err != nil {
			loop = false
			if err != io.EOF {
				break
			}
		}

		// Normalize line
		line = pr.normalizeLine(line)

		// Continue for blank & comment lines
		if len(line) < 1 || propertiesReaderRegexComment.MatchString(line) {
			continue
		}

		// Check for continuous line
		if pr.isContinuousLine(line) {
			text += line[:len(line)-1] // Remove the ending continuous mark ("\\")
		} else {
			text += line
			loop = false
		}
	}

	return text, err
}

// normalizeLine() normalizes a line (trimming spaces & line carriers)
func (pr *PropertiesReader) normalizeLine(line string) string {
	normalized := strings.TrimSpace(line)
	normalized = strings.TrimRight(normalized, "\r\n")
	return normalized
}

// isContinuousLine() checks whether there are succeeding lines after specified line.
// If current line ends with a single \, there are succeeding lines.
func (pr *PropertiesReader) isContinuousLine(line string) bool {
	matched := false
	if indexes := propertiesReaderRegexContinuousLine.FindStringIndex(line); len(indexes) == 2 {
		matched = (indexes[1]-indexes[0])%2 == 1
	}
	return matched
}

// parseKeyValue() extracts key and value settings from a line
func (pr *PropertiesReader) parseKeyValue(line string) (string, *string, error) {
	septPos := -1

	// Find key value separator
	allIndexes := propertiesReaderRegexSept.FindAllStringIndex(line, -1)
	for _, indexes := range allIndexes {
		sept := line[indexes[0]:indexes[1]]
		if pr.isValidSept(sept) {
			septPos = indexes[1]
			break
		}
	}

	// Parse key & value
	var key string
	var value *string = nil
	var err error = nil
	if septPos < 0 {
		key = strings.TrimSpace(line)
		key, err = pr.unescape(key)
	} else if septPos > 0 {
		key = strings.TrimSpace(line[:septPos-1])
		valueText := strings.TrimSpace(line[septPos:])
		if key, err = pr.unescape(key); err == nil {
			valueText, err = pr.unescape(valueText)
			value = &valueText
		}
	}
	if err == nil && len(key) < 1 {
		err = errors.New(fmt.Sprintf("key is not defined. line = %s", line))
	}

	// Exit
	return key, value, err
}

// isValidSept() checks whether the expression (containing =) is a valid separator between key and value
func (pr *PropertiesReader) isValidSept(sept string) bool {
	slackCount := strings.Count(sept, "\\")
	valid := slackCount%2 == 0
	return valid
}

// unescape() converts specified text to unescaped format.
func (pr *PropertiesReader) unescape(text string) (string, error) {
	var err error = nil
	unescaped := propertiesReaderReplace.ReplaceAllStringFunc(text, func(s string) string {
		ret := s
		switch s {
		case `\0`:
			ret = string(rune(0))
		case `\r`:
			ret = "\r"
		case `\n`:
			ret = "\n"
		case `\t`:
			ret = "\t"
		case `\\`:
			ret = "\\"
		case `\=`:
			ret = "="
		default:
			if strings.HasPrefix(s, `\u`) {
				if len(s) != 6 {
					err = errors.New(fmt.Sprintf("incorrect unicode format, text = %s, piece = %s", text, s))
				} else {
					ret, err = pr.unescapeUnicode(s[2:])
				}
			}
		}
		return ret
	})

	return unescaped, err
}

// unescapeUnicode() creates an unicode rune from specified string
func (pr *PropertiesReader) unescapeUnicode(text string) (string, error) {
	r, err := strconv.ParseInt(text, 16, 32)
	return string(rune(r)), err
}
