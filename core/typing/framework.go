/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package typing

import (
	"github.com/umeframework/gear/core/properties"
	"reflect"
)

// TypeDefinition interface provides a simple method to represent type definition.
type TypeDefinition interface {
	properties.Properties
	GetType() reflect.Type
	CreateInstance(args ...interface{}) interface{}
}

// CreateInstanceMethod() stands for methods to create a new instance.
// IMPORTANT: You should always return pointer to an instance.
type CreateInstanceMethod func(args ...interface{}) interface{}

// TypeDefinitionParam stands for user-defined extra param in type definition.
// It will be transformed to a read-only format in TypeDefinition interface
type TypeDefinitionParam map[properties.Key]properties.Value
