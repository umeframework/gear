/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"sort"
	"strings"
)

var (
	defaultLogFactory = &LogFactoryImpl{
		loggerMap:   make(map[string]Logger),
		loggerNames: make([]string, 0, 0x10),
	}
)

func GetLogFactory() LogFactory {
	return defaultLogFactory
}

func GetLogger(name string) Logger {
	return GetLogFactory().GetLogger(name)
}

func SetLogger(logger Logger) {
	GetLogFactory().SetLogger(logger)
}

type LogFactoryImpl struct {
	LogObjectImpl
	loggerMap   map[string]Logger
	loggerNames []string
}

func (lf *LogFactoryImpl) SetLogger(logger Logger) {
	if logger == nil {
		return
	}
	name := logger.GetName()
	_, exists := lf.loggerMap[name]
	lf.loggerMap[name] = logger
	if !exists {
		lf.loggerNames = append(lf.loggerNames, name)
		sort.Slice(lf.loggerNames, func(i, j int) bool {
			return lf.loggerNames[i] > lf.loggerNames[j]
		})
	}
}

func (lf *LogFactoryImpl) GetLogger(name string) Logger {
	loggerName, ok := lf.findLoggerName(name)
	if !ok {
		loggerName = ""
	}
	logger, _ := lf.loggerMap[loggerName]
	return logger
}

func (lf *LogFactoryImpl) findLoggerName(name string) (string, bool) {
	var ret string
	found := false

	for _, loggerName := range lf.loggerNames {
		if loggerName == name || strings.HasPrefix(loggerName, name+".") {
			ret = loggerName
			found = true
			break
		}
	}

	return ret, found
}
