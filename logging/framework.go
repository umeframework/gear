/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import "github.com/umeframework/gear/core/properties"

type LogLevel int

type LogObject interface {
	GetName() string
	Initialize(properties properties.Properties)
	Terminate()
}

type Logger interface {
	LogObject
	Log(level LogLevel, format string, args ...interface{})
}

type LogFactory interface {
	LogObject
	GetLogger(name string) Logger
	SetLogger(logger Logger)
}

type LogData interface {
	properties.Properties
}

type LogDataProvider interface {
	LogObject
	Provide(logData LogData)
}

type LogAppender interface {
	LogObject
	Append(logData LogData)
}

type LogLayout interface {
	LogObject
	GetHeader() string
	Format(logData LogData) string
	GetFooter() string
}
