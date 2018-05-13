/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

import (
	"bytes"
	"fmt"
	"github.com/umeframework/gear/core/collection"
)

type attributeImpl struct {
	namespace string
	name      string
	value     string
}

func newAttribute() *attributeImpl {
	return &attributeImpl{}
}

func (attr *attributeImpl) GetNamespace() string {
	return attr.namespace
}

func (attr *attributeImpl) SetNamespace(namespace string) {
	attr.namespace = namespace
}

func (attr *attributeImpl) GetName() string {
	return attr.name
}

func (attr *attributeImpl) SetName(name string) {
	attr.name = name
}

func (attr *attributeImpl) GetValue() string {
	return attr.value
}

func (attr *attributeImpl) SetValue(value string) {
	attr.value = value
}

func (attr *attributeImpl) String() string {
	if len(attr.namespace) > 0 {
		return fmt.Sprintf("%v::%v = %v", attr.namespace, attr.name, attr.value)
	} else {
		return fmt.Sprintf("%v = %v", attr.namespace, attr.value)
	}
}

type attributeMapImpl struct {
	list collection.List
}

func newAttributeMap() *attributeMapImpl {
	return &attributeMapImpl{
		list: collection.NewLinkedList(),
	}
}

func (attrmap *attributeMapImpl) NumAttr() int {
	return attrmap.list.Size()
}

func (attrmap *attributeMapImpl) GetAttrAt(index int) Attribute {
	if attr, ok := attrmap.list.GetAt(index); ok {
		return attr.(Attribute)
	}
	return nil
}

func (attrmap *attributeMapImpl) GetAttr(namespace string, name string) Attribute {
	var attr Attribute = nil
	for it := attrmap.list.GetIterator(); it.HasNext(); {
		item := it.Next().(Attribute)
		if item.GetNamespace() == namespace && item.GetName() == name {
			attr = item
			break
		}
	}
	return attr
}

func (attrmap *attributeMapImpl) matchNamespaceName(element collection.Element, param interface{}) bool {
	attr := element.(Attribute)
	if attr == nil {
		return false
	}
	args, ok := param.([]string)
	if (!ok) || len(args) < 2 {
		return false
	}
	namespace := args[0]
	name := args[1]

	return attr.GetNamespace() == namespace && attr.GetName() == name
}

func (attrmap *attributeMapImpl) SetAttr(attr Attribute) {
	if attr != nil {
		namespace := attr.GetNamespace()
		name := attr.GetName()
		index := attrmap.list.IndexIf(attrmap.matchNamespaceName, []string{namespace, name})
		if index >= 0 {
			// Update existing attribute
			attrmap.list.SetAt(index, attr)
		} else {
			// Add new attribute
			attrmap.list.Add(attr)
		}
	}
}

func (attrmap *attributeMapImpl) RemoveAttr(namespace string, name string) {
	attrmap.list.RemoveIf(attrmap.matchNamespaceName, []string{namespace, name})
}

func (attrmap *attributeMapImpl) String() string {
	count := attrmap.list.Size()
	buffer := bytes.NewBuffer(nil)
	for index, it := 0, attrmap.list.GetIterator(); it.HasNext(); {
		attr := it.Next().(Attribute)
		buffer.WriteString(fmt.Sprintf("%v", attr))
		if index < count-1 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
