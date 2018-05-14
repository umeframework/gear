/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

import "github.com/umeframework/gear/util/collection"

type nodeSetter interface {
	SetDocument(document Document)
	SetParent(parent Element)
}

type nodeEx interface {
	Node
	nodeSetter
}

type nodeImpl struct {
	document Document
	parent   Element
}

func newNode() *nodeImpl {
	return &nodeImpl{}
}

func (node *nodeImpl) GetDocument() Document {
	return node.document
}

func (node *nodeImpl) GetParent() Element {
	return node.parent
}

func (node *nodeImpl) SetDocument(document Document) {
	node.document = document
}

func (node *nodeImpl) SetParent(parent Element) {
	node.parent = parent
}

type childEventCallback func(child Node)

type nodeCollectionImpl struct {
	list       collection.List
	childAdded childEventCallback
	//childRemoved childEventCallback
}

func newNodeCollection(childAdded childEventCallback) NodeCollection {
	return &nodeCollectionImpl{
		list:       collection.NewLinkedList(),
		childAdded: childAdded,
	}
}

func (nodes *nodeCollectionImpl) NumChildren() int {
	return nodes.list.Size()
}

func (nodes *nodeCollectionImpl) GetNode(index int) Node {
	if node, ok := nodes.list.GetAt(index); ok {
		return node.(Node)
	}
	return nil
}

func (nodes *nodeCollectionImpl) SetNode(index int, child Node) {
	nodes.list.SetAt(index, child)
	if nodes.childAdded != nil {
		nodes.childAdded(child)
	}
}

func (nodes *nodeCollectionImpl) AddNode(child Node) {
	nodes.list.Add(child)
	if nodes.childAdded != nil {
		nodes.childAdded(child)
	}
}

func (nodes *nodeCollectionImpl) AddNodeAt(index int, child Node) {
	nodes.list.AddAt(index, child)
	if nodes.childAdded != nil {
		nodes.childAdded(child)
	}
}

func (nodes *nodeCollectionImpl) RemoveNode(child Node) {
	nodes.list.Remove(child)
}

func (nodes *nodeCollectionImpl) RemoveNodeAt(index int) {
	nodes.list.RemoveAt(index)
}

func (nodes *nodeCollectionImpl) ClearNodes() {
	nodes.list.Clear()
}
