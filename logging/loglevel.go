/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import "strings"

var (
	logLevelMap         = make(map[string]LogLevel)
	logLevelMapReversed = make(map[LogLevel]string)
)

func AddLogLevel(name string, logLevel LogLevel) {
	logLevelMap[strings.ToLower(name)] = logLevel
	logLevelMapReversed[logLevel] = name
}

func GetLogLevel(name string) (LogLevel, bool) {
	logLevel, ok := logLevelMap[strings.ToLower(name)]
	return logLevel, ok
}

func GetLogLevelName(level LogLevel) (string, bool) {
	name, ok := logLevelMapReversed[level]
	return name, ok
}

const (
	LevelVerbose LogLevel = 1000 * (iota + 1)
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal

	LevelTrace LogLevel = LevelVerbose
	LevelNone  LogLevel = LogLevel(^uint(0) >> 1)
	LevelAll   LogLevel = ^LevelNone
)

func init() {
	AddLogLevel("Verbose", LevelVerbose)
	AddLogLevel("Debug", LevelDebug)
	AddLogLevel("Info", LevelInfo)
	AddLogLevel("Warn", LevelWarn)
	AddLogLevel("Error", LevelError)
	AddLogLevel("Fatal", LevelFatal)
	AddLogLevel("Trace", LevelTrace)
	AddLogLevel("None", LevelNone)
	AddLogLevel("All", LevelAll)
}
