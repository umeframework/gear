/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

import (
	"encoding/xml"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/umeframework/gear/core/collection"
	"io"
	"strings"
)

var (
	DefaultOutputConfig = OutputConfig{
		Prefix: "",
		Indent: "\t",
	}
)

var (
	ErrorRootElementDuplicated = errors.New("root element already exists in xml document")
)

type documentImpl struct {
	NodeCollection
	root Element
}

func New() Document {
	doc := &documentImpl{}
	doc.NodeCollection = newNodeCollection(doc.onChildAdded)
	doc.root = doc.CreateElement()
	return doc
}

func Load(r io.Reader) (Document, error) {
	doc := &documentImpl{}
	doc.NodeCollection = newNodeCollection(doc.onChildAdded)
	err := doc.Load(r)
	return doc, err
}

func (doc *documentImpl) GetRoot() Element {
	return doc.root
}

func (doc *documentImpl) CreateElement() Element {
	element := newElement()
	element.SetDocument(doc)
	return element
}

func (doc *documentImpl) CreateComment() Comment {
	comment := newComment()
	comment.SetDocument(doc)
	return comment
}

func (doc *documentImpl) CreateAttribute() Attribute {
	attr := newAttribute()
	return attr
}

func (doc *documentImpl) onChildAdded(child Node) {
	if element, ok := child.(Element); ok {
		if doc.root != nil {
			panic(ErrorRootElementDuplicated)
		}
		doc.root = element
	}
}

func (doc *documentImpl) Load(r io.Reader) error {
	var err error = nil
	var token xml.Token = nil
	decoder := xml.NewDecoder(r)
	stack := collection.NewStack()

	// Reset current content
	doc.root = nil
	doc.ClearNodes()

	for {
		token, err = decoder.Token()
		if err != nil && err != io.EOF {
			break
		}

		var currentElemnt Element = nil
		if item, ok := stack.Peek(); ok {
			currentElemnt = item.(Element)
		}

		switch token.(type) {
		case xml.StartElement:
			se, _ := token.(xml.StartElement)
			element := doc.CreateElement()
			element.SetNamespace(se.Name.Space)
			element.SetName(se.Name.Local)
			for _, seAttr := range se.Attr {
				attr := doc.CreateAttribute()
				attr.SetNamespace(seAttr.Name.Space)
				attr.SetName(seAttr.Name.Local)
				attr.SetValue(seAttr.Value)
				element.SetAttr(attr)
			}
			stack.Push(element)
			if currentElemnt != nil {
				currentElemnt.AddNode(element)
			} else {
				if doc.root == nil {
					doc.AddNode(element)
				} else {
					err = ErrorRootElementDuplicated
					break
				}
			}
		case xml.EndElement:
			ee, _ := token.(xml.EndElement)
			if currentElemnt == nil || currentElemnt.GetNamespace() != ee.Name.Space || currentElemnt.GetName() != ee.Name.Local {
				err = errors.New(fmt.Sprintf("invlaid end element. namespace = %v, name = %v", ee.Name.Space, ee.Name.Local))
				break
			}
			stack.Pop()
		case xml.Comment:
			cmt, _ := token.(xml.Comment)
			text := string(cmt)
			comment := doc.CreateComment()
			comment.SetComment(text)
			if currentElemnt != nil {
				currentElemnt.AddNode(comment)
			} else {
				doc.AddNode(comment)
			}
		case xml.CharData:
			if currentElemnt != nil {
				cd, _ := token.(xml.CharData)
				text := string(cd)
				text = strings.TrimSpace(text)
				if len(text) > 0 {
					currentElemnt.SetValue(text)
				}
			}
		case xml.Directive:
			di, _ := token.(xml.Directive)
			directive := newDirective()
			directive.SetDirective(string(di))
			if currentElemnt != nil {
				currentElemnt.AddNode(directive)
			} else {
				doc.AddNode(directive)
			}
		case xml.ProcInst:
			pi, _ := token.(xml.ProcInst)
			procInst := newProcInst()
			procInst.SetDocument(doc)
			procInst.SetTarget(pi.Target)
			procInst.SetInst(string(pi.Inst))
			if currentElemnt != nil {
				currentElemnt.AddNode(procInst)
			} else {
				doc.AddNode(procInst)
			}
		default:
			//fmt.Println(token)
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}

	// Exit
	return err
}

func (doc *documentImpl) Save(cfg OutputConfig, w io.Writer) error {
	var err error = nil

	encoder := xml.NewEncoder(w)
	defer encoder.Flush()

	encoder.Indent(cfg.Prefix, cfg.Indent)
	err = doc.encodeNodeCollection(encoder, doc)

	return err
}

func (doc *documentImpl) encodeNode(encoder *xml.Encoder, node Node) error {
	var err error = nil
	switch node.(type) {
	case Element:
		element, _ := node.(Element)
		err = doc.encodeElement(encoder, element)
	case ProcInst:
		procinst, _ := node.(ProcInst)
		err = doc.encodeProcInst(encoder, procinst)
	case Comment:
		comment, _ := node.(Comment)
		err = doc.encodeComment(encoder, comment)
	case Directive:
		directive, _ := node.(Directive)
		err = doc.encodeDirective(encoder, directive)
	default:
	}
	return err
}

func (doc *documentImpl) encodeElement(encoder *xml.Encoder, element Element) error {
	// Write start element
	se := xml.StartElement{}
	se.Name.Space = element.GetNamespace()
	se.Name.Local = element.GetName()
	attrCount := element.NumAttr()
	se.Attr = make([]xml.Attr, attrCount)
	for i := 0; i < attrCount; i++ {
		attr := element.GetAttrAt(i)
		seAttr := xml.Attr{}
		seAttr.Name.Space = attr.GetNamespace()
		seAttr.Name.Local = attr.GetName()
		seAttr.Value = attr.GetValue()
		se.Attr[i] = seAttr
	}
	if err := encoder.EncodeToken(se); err != nil {
		return err
	}

	// Write value if exists
	value := element.GetValue()
	if len(value) > 0 {
		bytes := []byte(value)
		cd := xml.CharData(bytes)
		if err := encoder.EncodeToken(cd); err != nil {
			return err
		}
	}

	// Write children
	doc.encodeNodeCollection(encoder, element)

	// Write end element
	ee := se.End()
	if err := encoder.EncodeToken(ee); err != nil {
		return err
	}

	// Exit
	return nil
}

func (doc *documentImpl) encodeNodeCollection(encoder *xml.Encoder, nodes NodeCollection) error {
	count := nodes.NumChildren()
	for i := 0; i < count; i++ {
		child := nodes.GetNode(i)
		if err := doc.encodeNode(encoder, child); err != nil {
			return err
		}
	}
	return nil
}

func (doc *documentImpl) encodeComment(encoder *xml.Encoder, comment Comment) error {
	bytes := []byte(comment.GetComment())
	cmt := xml.Comment(bytes)
	return encoder.EncodeToken(cmt)
}

func (doc *documentImpl) encodeProcInst(encoder *xml.Encoder, procInst ProcInst) error {
	pi := xml.ProcInst{}
	pi.Target = procInst.GetTarget()
	pi.Inst = []byte(procInst.GetInst())
	return encoder.EncodeToken(pi)
}

func (doc *documentImpl) encodeDirective(encoder *xml.Encoder, directive Directive) error {
	bytes := []byte(directive.GetDirective())
	di := xml.Directive(bytes)
	return encoder.EncodeToken(di)
}
