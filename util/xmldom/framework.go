/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

import "io"

// Node represents a node in XML document.
// It can be ProcInst, Comment, Element or Directive.
type Node interface {
	GetDocument() Document
	GetParent() Element
}

// ProcInst represents an XML processing instruction of the form <? target inst ?>.
// Usually, it goes like <?xml version="1.0" encoding="UTF-8"?>.
type ProcInst interface {
	Node
	GetTarget() string
	SetTarget(target string)
	GetInst() string
	SetInst(inst string)
}

// Comment represents a comment block in XML document.
type Comment interface {
	Node
	GetComment() string
	SetComment(comment string)
}

// Comment represents an XML directive of the form <!text>.
type Directive interface {
	Node
	GetDirective() string
	SetDirective(directive string)
}

// Attribute represents an attribute in XML document.
// It is implemented as more similar to golang's Attr,
// instead of Java's equivalent which is much more heavier.
type Attribute interface {
	GetNamespace() string
	SetNamespace(namespace string)
	GetName() string
	SetName(name string)
	GetValue() string
	SetValue(value string)
}

// AttributeMap represents a series of Attribute,
// while providing both sequenced and key-based accesses.
type AttributeMap interface {
	NumAttr() int
	GetAttrAt(index int) Attribute
	GetAttr(namespace string, name string) Attribute
	SetAttr(attr Attribute)
	RemoveAttr(namespace string, name string)
}

// NodeCollection represents a collection of Node.
type NodeCollection interface {
	NumChildren() int
	GetNode(index int) Node
	SetNode(index int, child Node)
	AddNode(child Node)
	AddNodeAt(index int, child Node)
	RemoveNode(child Node)
	RemoveNodeAt(index int)
	ClearNodes()
}

// Element represents an element in XML document.
type Element interface {
	Node
	Attribute
	AttributeMap
	NodeCollection
}

type OutputConfig struct {
	Prefix string
	Indent string
}

// Document represents a XML document, with Load & Save features.
type Document interface {
	NodeCollection
	GetRoot() Element

	CreateElement() Element
	CreateComment() Comment
	CreateAttribute() Attribute

	Load(r io.Reader) error
	Save(cfg OutputConfig, w io.Writer) error
}
