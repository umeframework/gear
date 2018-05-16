/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"github.com/umeframework/gear/core/properties"
	"strings"
)

type logDataImpl struct {
	properties properties.Properties
}

func (logData *logDataImpl) normalizeKey(key properties.Key) properties.Key {
	key2 := key
	switch key.(type) {
	case string:
		text := key.(string)
		key2 = strings.ToLower(text)
	default:
	}
	return key2
}

func (logData *logDataImpl) GetProp(key properties.Key) (properties.Value, bool) {
	key = logData.normalizeKey(key)
	return logData.properties.GetProp(key)
}

func (logData *logDataImpl) SetProp(key properties.Key, value properties.Value) {
	key = logData.normalizeKey(key)
	logData.properties.SetProp(key, value)
}

func (logData *logDataImpl) RemoveProp(key properties.Key) {
	key = logData.normalizeKey(key)
	logData.properties.RemoveProp(key)
}

func (logData *logDataImpl) ClearProps() {
	logData.properties.ClearProps()
}

func (logData *logDataImpl) GetProps() []properties.Key {
	return logData.properties.GetProps()
}

func NewLogData() LogData {
	return &logDataImpl{
		properties: properties.New(),
	}
}

type logDataProvideMethod func(name string, logData LogData)

type logDataProviderImpl struct {
	logDataTypes      []string
	logDataMethodDict map[string]logDataProvideMethod
}

func (prov *logDataProviderImpl) Provide(logData LogData) {
	panic("implement me")
}
