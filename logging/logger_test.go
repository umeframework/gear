/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(LevelInfo)

	provider := NewLogDataProvider()
	logger.AddProvider(provider)

	provider2 := &EnvDataProvider{
		Prefix: "env.",
	}
	logger.AddProvider(provider2)

	layout := NewPatternLayout()
	writer := os.Stdout
	appender := NewLogWriterAppender(writer)
	appender.Layout = layout
	logger.AddAppender(appender)

	logger.Log(LevelDebug, "this is a test on %v", time.Now())
	logger.Log(LevelWarn, "warning: something happended")
}
