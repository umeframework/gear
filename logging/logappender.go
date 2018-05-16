/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"fmt"
	"github.com/umeframework/gear/core/properties"
	"io"
)

const (
	LogWriterAppenderDefaultBatchSize = 0x10
)

type LogWriterAppenderImpl struct {
	LogObjectImpl
	Layout         LogLayout
	FlushBatchSize int
	Writer         io.Writer
}

func NewLogWriterAppender(writer io.Writer) *LogWriterAppenderImpl {
	appender := &LogWriterAppenderImpl{
		Writer:         writer,
		FlushBatchSize: LogWriterAppenderDefaultBatchSize,
	}
	appender.OnInitialized = appender.onInitialized
	appender.OnTerminated = appender.onTerminated
	return appender
}

func (appender *LogWriterAppenderImpl) onInitialized(properties properties.Properties) {
	if appender.Writer != nil && appender.Layout != nil {
		header := appender.Layout.GetHeader()
		appender.writeText(header)
	}
}

func (appender *LogWriterAppenderImpl) onTerminated() {
	if appender.Writer != nil && appender.Layout != nil {
		footer := appender.Layout.GetFooter()
		appender.writeText(footer)
	}
}

func (appender *LogWriterAppenderImpl) Append(logData LogData) {
	if appender.Writer != nil && appender.Layout != nil {
		text := appender.Layout.Format(logData)
		appender.writeText(text)
	}
}

func (appender *LogWriterAppenderImpl) writeText(text string) {
	fmt.Fprint(appender.Writer, text)
}
