/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package typing

import (
	"github.com/umeframework/gear/core/properties"
	"reflect"
	"sync"
)

var (
	typeDefDict      = make(map[string]TypeDefinition)
	typeDefDictMutex = sync.RWMutex{}
)

type typeDefinition struct {
	properties.Properties
	t            reflect.Type
	createMethod CreateInstanceMethod
}

func (td *typeDefinition) GetType() reflect.Type {
	return td.t
}

func (td *typeDefinition) CreateInstance(args ...interface{}) interface{} {
	var instance interface{} = nil
	if td.createMethod != nil {
		instance = td.createMethod(args)
	} else {
		instance = reflect.New(td.GetType()).Interface()
	}
	return instance
}

func Register(name string, t reflect.Type) TypeDefinition {
	return RegisterEx(name, t, nil, nil)
}

func RegisterEx(name string, t reflect.Type, createMethod CreateInstanceMethod, param TypeDefinitionParam) TypeDefinition {
	typeDef := typeDefinition{
		Properties:   properties.New(),
		t:            t,
		createMethod: createMethod,
	}
	properties.ToProperties(typeDef.Properties, param)

	typeDefDictMutex.Lock()
	defer typeDefDictMutex.Unlock()

	typeDefDict[name] = &typeDef
	return &typeDef
}

func FromName(name string) (TypeDefinition, bool) {
	typeDefDictMutex.RLock()
	defer typeDefDictMutex.RUnlock()

	typeDef, ok := typeDefDict[name]
	return typeDef, ok
}

func GetTypeNames() []string {
	typeDefDictMutex.RLock()
	defer typeDefDictMutex.RUnlock()

	names := make([]string, 0, len(typeDefDict))
	for key, _ := range typeDefDict {
		names = append(names, key)
	}
	return names
}

func init() {
	// Register built-in types
	Register("bool", reflect.TypeOf(true))
	Register("byte", reflect.TypeOf(byte(0)))

	Register("int8", reflect.TypeOf(int8(0)))
	Register("int16", reflect.TypeOf(int16(0)))
	Register("int32", reflect.TypeOf(int32(0)))
	Register("int64", reflect.TypeOf(int64(0)))

	Register("uint", reflect.TypeOf(int(0)))
	Register("uint8", reflect.TypeOf(int8(0)))
	Register("uint16", reflect.TypeOf(int16(0)))
	Register("uint32", reflect.TypeOf(int32(0)))
	Register("uint64", reflect.TypeOf(int64(0)))
	Register("uint", reflect.TypeOf(int(0)))

	Register("float32", reflect.TypeOf(float32(0)))
	Register("float64", reflect.TypeOf(float64(0)))

	Register("rune", reflect.TypeOf(rune(0)))
	Register("string", reflect.TypeOf(""))
}
