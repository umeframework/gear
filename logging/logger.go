/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"fmt"
	"github.com/umeframework/gear/util/collection"
	"runtime"
)

const (
	DataLevel   = "log.level"
	DataFormat  = "log.fmt"
	DataArgs    = "log.args"
	DataMessage = "log.message"
	DataNewLine = "log.newline"
)

var (
	loggerNewLine string
)

type AppenderCollection interface {
	NumAppenders() int
	GetAppender(index int) LogAppender
	SetAppender(index int, appender LogAppender)
	AddAppender(appender LogAppender)
	AddAppenderAt(index int, appender LogAppender)
	RemoveAppender(appender LogAppender)
	RemoveAppenderAt(index int)
	ClearAppenders()
}

type ProviderCollection interface {
	NumProviders() int
	GetProvider(index int) LogDataProvider
	SetProvider(index int, Provider LogDataProvider)
	AddProvider(provider LogDataProvider)
	AddProviderAt(index int, Provider LogDataProvider)
	RemoveProvider(provider LogDataProvider)
	RemoveProviderAt(index int)
	ClearProviders()
}

type LoggerImpl struct {
	LogObjectImpl
	level     LogLevel
	appenders collection.List
	providers collection.List
}

func NewLogger() *LoggerImpl {
	logger := &LoggerImpl{
		appenders: collection.NewLinkedList(),
		providers: collection.NewLinkedList(),
	}
	return logger
}

func (logger *LoggerImpl) GetLevel() LogLevel {
	return logger.level
}

func (logger *LoggerImpl) SetLevel(level LogLevel) {
	logger.level = level
}

func (logger *LoggerImpl) NumAppenders() int {
	return logger.appenders.Size()
}

func (logger *LoggerImpl) GetAppender(index int) LogAppender {
	appender, _ := logger.appenders.GetAt(index)
	return appender.(LogAppender)
}

func (logger *LoggerImpl) SetAppender(index int, appender LogAppender) {
	logger.appenders.SetAt(index, appender)
}

func (logger *LoggerImpl) AddAppender(appender LogAppender) {
	logger.appenders.Add(appender)
}

func (logger *LoggerImpl) AddAppenderAt(index int, appender LogAppender) {
	logger.appenders.AddAt(index, appender)
}

func (logger *LoggerImpl) RemoveAppender(appender LogAppender) {
	logger.appenders.Remove(appender)
}

func (logger *LoggerImpl) RemoveAppenderAt(index int) {
	logger.appenders.RemoveAt(index)
}

func (logger *LoggerImpl) ClearAppenders() {
	logger.appenders.Clear()
}

func (logger *LoggerImpl) NumProviders() int {
	return logger.providers.Size()
}

func (logger *LoggerImpl) GetProvider(index int) LogDataProvider {
	Provider, _ := logger.providers.GetAt(index)
	return Provider.(LogDataProvider)
}

func (logger *LoggerImpl) SetProvider(index int, Provider LogDataProvider) {
	logger.providers.SetAt(index, Provider)
}

func (logger *LoggerImpl) AddProvider(provider LogDataProvider) {
	logger.providers.Add(provider)
}

func (logger *LoggerImpl) AddProviderAt(index int, Provider LogDataProvider) {
	logger.providers.AddAt(index, Provider)
}

func (logger *LoggerImpl) RemoveProvider(provider LogDataProvider) {
	logger.providers.Remove(provider)
}

func (logger *LoggerImpl) RemoveProviderAt(index int) {
	logger.providers.RemoveAt(index)
}

func (logger *LoggerImpl) ClearProviders() {
	logger.providers.Clear()
}

func (logger *LoggerImpl) Log(level LogLevel, format string, args ...interface{}) {
	if logger.levelEnabled(level) {
		logger.forceLog(level, format, args...)
	}
}

func (logger *LoggerImpl) levelEnabled(level LogLevel) bool {
	return logger.level <= level
}

func (logger *LoggerImpl) forceLog(level LogLevel, format string, args ...interface{}) {
	logData := logger.createLogData(level, format, args)
	logger.callAppenders(logData)
}

func (logger *LoggerImpl) createLogData(level LogLevel, format string, args []interface{}) LogData {
	logData := NewLogData()
	// Initialize log data
	levelName, _ := GetLogLevelName(level)
	logData.SetProp(DataLevel, levelName)
	logData.SetProp(DataFormat, format)
	logData.SetProp(DataArgs, args)
	logData.SetProp(DataMessage, fmt.Sprintf(format, args...))
	logData.SetProp(DataNewLine, loggerNewLine)

	// Call succeeding log data providers
	logger.callDataProviders(logData)

	// Exit
	return logData
}

func (logger *LoggerImpl) callAppenders(logData LogData) {
	for it := logger.appenders.GetIterator(); it.HasNext(); {
		if appender, _ := it.Next().(LogAppender); appender != nil {
			appender.Append(logData)
		}
	}
}

func (logger *LoggerImpl) callDataProviders(logData LogData) {
	for it := logger.providers.GetIterator(); it.HasNext(); {
		if provider, _ := it.Next().(LogDataProvider); provider != nil {
			provider.Provide(logData)
		}
	}
}

func init() {
	switch runtime.GOOS {
	case "windows":
		loggerNewLine = "\r\n"
	default:
		loggerNewLine = "\n"
	}
}
