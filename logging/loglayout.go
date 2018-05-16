/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"fmt"
	"github.com/umeframework/gear/util/collection"
	gformat "github.com/umeframework/gear/util/format"
	"regexp"
	"strings"
	"time"
)

const (
	indexFullText = iota
	indexExpression

	DefaultLayoutPattern = `[${log.timestamp : yyyy-MM-dd hh:mm:ss.SSS}][${log.level}][${log.method}] ${log.message} ${log.newline}`
)

var (
	//loglayoutRegex = regexp.MustCompile(`([%]+)([-]?[0-9.]+)*([a-zA-Z]+)((\s)*{[\w\W]+})?`)
	loglayoutRegex = regexp.MustCompile(`\${([^\}]+)\}`)
)

type FormatterCollection interface {
	NumFormatters() int
	GetFormatter(index int) gformat.Formatter
	SetFormatter(index int, Formatter gformat.Formatter)
	AddFormatter(Formatter gformat.Formatter)
	AddFormatterAt(index int, Formatter gformat.Formatter)
	RemoveFormatter(Formatter gformat.Formatter)
	RemoveFormatterAt(index int)
	ClearFormatters()
}

type PatternLayout struct {
	LogObjectImpl
	header     string
	pattern    string
	footer     string
	formatters collection.List
}

func NewPatternLayout() *PatternLayout {
	layout := &PatternLayout{
		pattern:    DefaultLayoutPattern,
		formatters: collection.NewLinkedList(),
	}
	return layout
}

func (layout *PatternLayout) GetHeader() string {
	return layout.header
}

func (layout *PatternLayout) SetHeader(header string) {
	layout.header = header
}

func (layout *PatternLayout) Format(logData LogData) string {
	text := loglayoutRegex.ReplaceAllStringFunc(layout.pattern, func(s string) string {
		var r string = s
		if submatches := loglayoutRegex.FindAllStringSubmatch(s, -1); len(submatches) == 1 {
			if submatch := submatches[0]; len(submatch) > indexExpression {
				expression := submatch[indexExpression]
				var argName, format string
				if pos := strings.IndexRune(expression, ':'); pos > 0 {
					argName = expression[:pos]
					format = expression[pos+1:]
				} else {
					argName = expression
				}
				argName = strings.TrimSpace(argName)
				format = strings.TrimSpace(format)
				if arg, ok := logData.GetProp(argName); ok {
					r = layout.format(arg, format)
				}
			}
		}

		return r
	})

	return text
}

func (layout *PatternLayout) format(arg interface{}, format string) string {
	var text string
	if formatable, _ := arg.(gformat.Formatable); formatable != nil {
		text, _ = formatable.Format(format)
	} else if formatter, _ := arg.(gformat.Formatter); formatter != nil {
		text, _, _ = formatter.Format(arg, format)
	} else {
		if text2, _, handled := layout.callFormatters(arg, format); handled {
			text = text2
		} else {
			text = layout.defaultFormat(arg, format)
		}
	}
	return text
}

func (layout *PatternLayout) callFormatters(arg interface{}, format string) (string, error, bool) {
	var text string
	var err error
	var handled bool

	for it := layout.formatters.GetIterator(); it.HasNext(); {
		if formatter, _ := it.Next().(gformat.Formatter); formatter != nil {
			if text, err, handled = formatter.Format(arg, format); handled && err == nil {
				break
			}
		}
	}

	return text, err, handled
}

func (layout *PatternLayout) defaultFormat(arg interface{}, format string) string {
	var text string

	switch arg.(type) {
	case time.Time:
		formatter := gformat.TimeFormat{
			Pattern: format,
		}
		text, _ = formatter.Format(arg)
	case time.Duration:
		duration, _ := arg.(time.Duration)
		text = duration.String()
	default:
		if len(format) < 1 {
			text = fmt.Sprintf("%v", arg)
		} else {
			text = fmt.Sprintf(format, arg)
		}
	}

	return text
}

func (layout *PatternLayout) GetFooter() string {
	return layout.footer
}

func (layout *PatternLayout) SetFooter(footer string) {
	layout.footer = footer
}

func (layout *PatternLayout) GetPattern() string {
	return layout.pattern
}

func (layout *PatternLayout) SetPattern(pattern string) {
	layout.pattern = pattern
}

func (layout *PatternLayout) NumFormatters() int {
	return layout.formatters.Size()
}

func (layout *PatternLayout) GetFormatter(index int) gformat.Formatter {
	formatter, _ := layout.formatters.GetAt(index)
	return formatter.(gformat.Formatter)
}

func (layout *PatternLayout) SetFormatter(index int, formatter gformat.Formatter) {
	layout.formatters.SetAt(index, formatter)
}

func (layout *PatternLayout) AddFormatter(formatter gformat.Formatter) {
	layout.formatters.Add(formatter)
}

func (layout *PatternLayout) AddFormatterAt(index int, formatter gformat.Formatter) {
	layout.formatters.AddAt(index, formatter)
}

func (layout *PatternLayout) RemoveFormatter(formatter gformat.Formatter) {
	layout.formatters.Remove(formatter)
}

func (layout *PatternLayout) RemoveFormatterAt(index int) {
	layout.formatters.RemoveAt(index)
}

func (layout *PatternLayout) ClearFormatters() {
	layout.formatters.Clear()
}
