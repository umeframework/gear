/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package logging

import "github.com/umeframework/gear/core/properties"

type LogObjectImpl struct {
	name       string
	properties properties.Properties

	OnNameModified func(name string)
	OnInitialized  func(properties properties.Properties)
	OnTerminated   func()
}

func (obj *LogObjectImpl) GetName() string {
	return obj.name
}
func (obj *LogObjectImpl) SetName(name string) {
	obj.name = name
	if obj.OnNameModified != nil {
		obj.OnNameModified(name)
	}
}

func (obj *LogObjectImpl) Initialize(properties properties.Properties) {
	obj.properties = properties
	if obj.OnInitialized != nil {
		obj.OnInitialized(properties)
	}
}

func (obj *LogObjectImpl) Terminate() {
	if obj.OnTerminated != nil {
		obj.OnTerminated()
	}
}
