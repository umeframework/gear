/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package typing

import (
	"github.com/wingsweaver/gear/core/properties"
	"reflect"
)

type TypeDefinition interface {
	properties.PropertyBag
	GetType() reflect.Type
	CreateInstance(args ...interface{}) interface{}
}

type CreateInstanceMethod func(args ...interface{}) interface{}

type TypeDefinitionParam map[interface{}]interface{}
