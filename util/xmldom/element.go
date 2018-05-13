/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

type elementImpl struct {
	nodeEx
	Attribute
	AttributeMap
	NodeCollection
}

func newElement() *elementImpl {
	element := &elementImpl{
		nodeEx:       newNode(),
		Attribute:    newAttribute(),
		AttributeMap: newAttributeMap(),
	}
	element.NodeCollection = newNodeCollection(element.onChildAdded)
	return element
}

func (elem *elementImpl) onChildAdded(child Node) {
	if setter, _ := child.(nodeSetter); setter != nil {
		setter.SetParent(elem)
	}
}
