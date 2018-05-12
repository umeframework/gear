/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

var (
	iniWriterRunesNoUseEscape = []rune{
		0, '\a', '\b', '\t', '\r', '\n', '#', '=',
	}
	iniWriterMatchesNoUseEscape  = string(iniWriterRunesNoUseEscape)
	iniWriterReplacesNoUseEscape = []string{
		`\0`, `\a`, `\b`, `\t`, `\r`, `\n`, `\#`, `\=`,
	}

	iniWriterRunesUseEscape    = append(iniWriterRunesNoUseEscape, '\\')
	iniWriterMatchesUseEscape  = string(iniWriterRunesUseEscape)
	iniWriterReplacesUseEscape = append(iniWriterReplacesNoUseEscape, `\\`)
)

type IniWriterConfig struct {
	UseEscape     bool
	IdentKeyValue bool
}

type iniWriter struct {
	config IniWriterConfig
}

func NewIniWriter(config IniWriterConfig) Writer {
	return &iniWriter{
		config: config,
	}
}

func (iw *iniWriter) Write(config Config, w io.Writer) error {
	var err error = nil
	writer := bufio.NewWriter(w)
	defer writer.Flush()
	count := config.NumChildren()
	for i := 0; i < count; i++ {
		child := config.GetChild(i)
		if err = iw.writeSection(child, writer); err != nil {
			break
		}
		fmt.Fprintln(writer)
	}

	return err
}

func (iw *iniWriter) writeSection(config Config, w *bufio.Writer) error {
	// Write section
	sectionName := iw.escapef("[%v]", config.GetKey())
	fmt.Fprintln(w, sectionName)

	// Write each key
	keys := config.GetProps()
	sort.Slice(keys, func(i, j int) bool {
		text := fmt.Sprintf("%v", keys[i])
		text2 := fmt.Sprintf("%v", keys[j])
		return strings.Compare(text, text2) < 0
	})
	for _, key := range keys {
		if value, ok := config.GetProp(key); ok {
			keyText := iw.escapef("%v", key)
			valueText := iw.escapef("%v", value)
			fmt.Fprintln(w, fmt.Sprintf("%s=%s", keyText, valueText))
		}
	}

	return nil
}

func (iw *iniWriter) escape(text string) string {
	matchString := iniWriterMatchesNoUseEscape
	replaceStrings := iniWriterReplacesNoUseEscape

	if iw.config.UseEscape {
		matchString = iniWriterMatchesUseEscape
		replaceStrings = iniWriterReplacesUseEscape
	}

	// Check whether escapement is needed
	if !strings.ContainsAny(text, matchString) {
		return text
	}

	// Escape
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	for _, r := range text {
		if index := strings.IndexRune(matchString, r); index >= 0 {
			buffer.WriteString(replaceStrings[index])
		} else {
			buffer.WriteRune(r)
		}
	}

	// Exit
	writer.Flush()
	escaped := string(buffer.Bytes())
	return escaped
}

func (iw *iniWriter) escapef(format string, args ...interface{}) string {
	text := fmt.Sprintf(format, args...)
	escaped := iw.escape(text)
	return escaped
}
