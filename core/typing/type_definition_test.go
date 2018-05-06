/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package typing

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestBuiltInTypes(t *testing.T) {
	typeNames := []string{
		"bool", "byte",
	}
	boolValue := true
	byteValue := byte(0)
	typesExpected := []reflect.Type{
		reflect.TypeOf(&boolValue), reflect.TypeOf(&byteValue),
	}

	for index, typeName := range typeNames {
		if typeDef, ok := FromName(typeName); !ok {
			t.Errorf("failed to find type definition for type: %s", typeName)
		} else {
			instance := typeDef.CreateInstance()
			instanceType := reflect.TypeOf(instance)
			typeExpected := typesExpected[index]
			if instanceType != typeExpected {
				t.Errorf("failed to create instance for type: %s. expected type: %v, real type: %v",
					typeName, typeExpected, instanceType)
			}
		}
	}
}

type TestPublicStruct struct {
	Id   int
	Name string
}

func TestPublicTypes(t *testing.T) {
	oldTypeNames := sort.StringSlice(GetTypeNames())
	typeName := "types.test.TestPublicStruct"
	if _, ok := FromName(typeName); ok {
		t.Errorf("type name %s should not exist", typeName)
		return
	}

	typeExpected := reflect.TypeOf(TestPublicStruct{})
	Register(typeName, typeExpected)
	if typeDef, ok := FromName(typeName); !ok {
		t.Errorf("type name %s should exist", typeName)
		return
	} else {
		if typeReal := typeDef.GetType(); typeReal != typeExpected {
			t.Errorf("type retrieved is wrong for %s, expected: %v, real: %v", typeName, typeExpected, typeReal)
			return
		}
		instance := typeDef.CreateInstance()
		if object, ok := instance.(*TestPublicStruct); !ok {
			t.Errorf("failed to create instance for %s", typeName)
		} else {
			object.Id = 100
			object.Name = "hello"
			fmt.Println(object)
		}
	}

	newTypeNames := sort.StringSlice(GetTypeNames())
	oldTypeNames = append(oldTypeNames, typeName)
	oldTypeNames.Sort()
	newTypeNames.Sort()
	for index, oldTypeName := range oldTypeNames {
		newTypeName := newTypeNames[index]
		if oldTypeName != newTypeName {
			t.Errorf("type names retrieved are difference: expected: %v, real: %v", oldTypeName, newTypeName)
		}
	}
}

type areable interface {
	Area() int
}

type rect struct {
	width, height int
}

func (r *rect) Area() int {
	return r.width * r.height
}

func TestPrivateStructWithCreation(t *testing.T) {
	var typeName = "types.test.TestPrivateStruct"
	var width = 200
	var height = 100
	RegisterEx(typeName, reflect.TypeOf(rect{}), func(args ...interface{}) interface{} {
		r := rect{
			width:  width,
			height: height,
		}
		return &r
	}, TypeDefinitionParam{
		"package":    "test",
		"registered": time.Now(),
	})

	if typeDef, ok := FromName(typeName); !ok {
		t.Errorf("failed to get type for %s", typeName)
	} else {
		instance := typeDef.CreateInstance()
		if object, ok := instance.(*rect); !ok {
			t.Errorf("failed to create rect")
		} else {
			if object.width != width || object.height != height {
				t.Errorf("content of rect created is not correct")
			}
		}
	}

}
