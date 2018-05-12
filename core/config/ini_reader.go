/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package config

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
	iniReaderRegexComment = regexp.MustCompile(`^[\s]*[#]+`)
	iniReaderRegexSection = regexp.MustCompile(`^[\s]*\[[^\]]+\]`)
	iniReaderRegexSept    = regexp.MustCompile(`([\\]*)=`)
	iniReaderReplace      = regexp.MustCompile(`\\0|\\\\|\\=|\\[a-tv-z]|\\u[\w\W]{0,4}`)
)

type IniReaderConfig struct {
	UseEscape bool
}

type iniReader struct {
	config IniReaderConfig
}

func NewIniReader(config IniReaderConfig) Reader {
	return &iniReader{
		config: config,
	}
}

func (ir *iniReader) Read(config Config, r io.Reader) error {
	reader := bufio.NewReader(r)

	var sectionName string
	var lines []string
	lastSectionName := ""
	var err error = nil

	for loop := true; loop; {
		if sectionName, lines, lastSectionName, err = ir.readBlock(reader, lastSectionName); err != nil {
			loop = false
			if err != io.EOF {
				break
			}
			err = nil
		}

		if child, e := ir.createProperties(sectionName, lines); e != nil {
			err = e
			break
		} else {
			config.AddChild(config.NumChildren(), child)
		}
	}
	return err
}

// readBlock() reads a block (section) from reader
func (ir *iniReader) readBlock(r *bufio.Reader, lastSectionName string) (string, []string, string, error) {
	var sectionName string
	lines := make([]string, 0)
	var err error = nil
	var line string

	if len(lastSectionName) > 0 {
		sectionName = lastSectionName
	}

	for loop := true; loop; {
		// Read a line
		if line, err = r.ReadString('\n'); err != nil {
			loop = false
			if err != io.EOF {
				break
			}
		}

		// Normalize line
		line = ir.normalizeLine(line)

		// Continue for blank & comment lines
		if len(line) < 1 || iniReaderRegexComment.MatchString(line) {
			continue
		}

		// Check whether the line is section
		if iniReaderRegexSection.MatchString(line) {
			if len(sectionName) > 0 { // 2nd section
				lastSectionName = line
				break
			} else {
				sectionName = line
				continue
			}
		}

		// Add other lines to array
		lines = append(lines, line)
	}

	// Exit
	return sectionName, lines, lastSectionName, err
}

// normalizeLine() normalizes a line (trimming spaces & line carriers)
func (ir *iniReader) normalizeLine(line string) string {
	normalized := strings.TrimSpace(line)
	normalized = strings.TrimRight(normalized, "\r\n")
	return normalized
}

// createProperties() creates a new Properties instance from specified section.
// It sets the section name to the key, and key-value settings to the props.
func (ir *iniReader) createProperties(sectionName string, lines []string) (Config, error) {
	config := New()
	var err error = nil

	// Extract section header and set to key
	config.SetKey(ir.normalizeSectionName(sectionName))

	// Extract key and value
	for _, line := range lines {
		if key, value, e := ir.parseKeyValue(line); e != nil {
			err = e
			break
		} else {
			if value != nil {
				config.SetProp(key, *value)
			} else {
				config.SetProp(key, nil)
			}
		}
	}

	// Exit
	return config, err
}

// normalizeSectionName() normalizes the section name (trimming `[]` in the two sides)
func (ir *iniReader) normalizeSectionName(sectionName string) string {
	normalized := strings.TrimLeft(sectionName, "[")
	normalized = strings.TrimRight(normalized, "]")
	normalized = strings.TrimSpace(normalized)
	return normalized
}

// parseKeyValue() extracts key and value settings from a line
func (ir *iniReader) parseKeyValue(line string) (string, *string, error) {
	septPos := -1

	// Find key value separator
	indexes := iniReaderRegexSept.FindStringIndex(line)
	if count := len(indexes); count > 0 && count%2 == 0 {
		for i := 0; i < count; i += 2 {
			sept := line[indexes[i]:indexes[i+1]]
			if ir.isValidSept(sept) {
				septPos = indexes[i+1]
				break
			}
		}
	}

	// Parse key & value
	var key string
	var value *string = nil
	var err error = nil
	if septPos < 0 {
		key = strings.TrimSpace(line)
		key, err = ir.unescape(key)
	} else if septPos > 0 {
		key = strings.TrimSpace(line[:septPos-1])
		valueText := strings.TrimSpace(line[septPos:])
		if key, err = ir.unescape(key); err == nil {
			valueText, err = ir.unescape(valueText)
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
func (ir *iniReader) isValidSept(sept string) bool {
	valid := false
	slackCount := strings.Count(sept, "\\")
	if ir.config.UseEscape {
		valid = slackCount%2 == 0
	} else {
		valid = slackCount != 1
	}
	return valid
}

// unescape() converts specified text to unescaped format.
func (ir *iniReader) unescape(text string) (string, error) {
	var err error = nil
	unescaped := iniReaderReplace.ReplaceAllStringFunc(text, func(s string) string {
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
		case `\:`:
			ret = ":"
		case `\#`:
			ret = "#"
		default:
			if strings.HasPrefix(s, `\u`) {
				if len(s) != 6 {
					err = errors.New(fmt.Sprintf("incorrect unicode format, text = %s, piece = %s", text, s))
				} else {
					ret, err = ir.unescapeUnicode(s[2:])
				}
			}
		}
		return ret
	})

	return unescaped, err
}

// unescapeUnicode() creates an unicode rune from specified string
func (ir *iniReader) unescapeUnicode(text string) (string, error) {
	r, err := strconv.ParseInt(text, 16, 32)
	return string(rune(r)), err
}
