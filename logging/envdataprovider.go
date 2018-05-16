/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import (
	"os"
	"strings"
)

type EnvDataProvider struct {
	LogObjectImpl
	Prefix string
}

func (provider *EnvDataProvider) Provide(logData LogData) {
	envs := os.Environ()
	var key, value string
	for _, env := range envs {
		pos := strings.IndexRune(env, '=')
		key = env[:pos]
		value = env[pos+1:]
		logData.SetProp(provider.Prefix+key, value)
	}
}
