/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	DataTimestamp  = "log.timestamp"
	DataElapsed    = "log.elapsed"
	DataType       = "log.type"
	DataClass      = DataType
	DataFile       = "log.file"
	DataLine       = "log.line"
	DataMethod     = "log.method"
	DataFileFull   = "log.fullfile"
	DataMethodFull = "log.fullmethod"

	logDataProviderCallDepth = 5
)

var (
	logDataProviderStartTime = time.Now()
)

type LogDataProviderImpl struct {
	LogObjectImpl
	CallDepth int
}

func NewLogDataProvider() *LogDataProviderImpl {
	return &LogDataProviderImpl{
		CallDepth: logDataProviderCallDepth,
	}
}

func (provider *LogDataProviderImpl) Provide(logData LogData) {
	// Timestamp
	logData.SetProp(DataTimestamp, time.Now())

	// Elapsed
	logData.SetProp(DataElapsed, time.Since(logDataProviderStartTime))

	// Type, File, Line & Method
	if ptr, fullFile, line, ok := runtime.Caller(logDataProviderCallDepth); ok {
		if fc := runtime.FuncForPC(ptr); fc != nil {
			fullMethod := fc.Name()
			logData.SetProp(DataMethodFull, fullMethod)
			_, method := path.Split(fullMethod)
			logData.SetProp(DataMethod, method)
			pos := strings.LastIndex(method, ".")
			typeName := method[:pos]
			logData.SetProp(DataType, typeName)
		}
		logData.SetProp(DataFileFull, fullFile)
		_, file := path.Split(fullFile)
		logData.SetProp(DataFile, file)
		logData.SetProp(DataLine, line)
	}
}
