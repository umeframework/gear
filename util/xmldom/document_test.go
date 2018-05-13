/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

import (
	"os"
	"path"
	"strings"
	"testing"
)

func walkNode(t *testing.T, node Node, depth int) {
	prefix := strings.Repeat("  ", depth)
	switch node.(type) {
	case Element:
		element := node.(Element)
		walkElement(t, element, depth)
	case Comment:
		comment := node.(Comment)
		t.Logf("%v%v", prefix, comment)
	default:
		t.Errorf("%vinvalid node found. %v", prefix, node)
	}
}

func walkElement(t *testing.T, element Element, depth int) {
	//prefix := strings.Repeat("  ", depth)
	//t.Logf("%v%v::%v, %v", prefix, element.GetNamespace(), element.GetName(), element.GetValue(), element.(AttributeMap))
	count := element.NumChildren()
	for i := 0; i < count; i++ {
		node := element.GetNode(i)
		walkNode(t, node, depth+1)
	}
}

func TestDocument(t *testing.T) {
	fileName := "./testdata/log.xml"
	file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("failed to open file %v. %v", fileName, err)
	}
	defer file.Close()

	doc, err := Load(file)
	if err != nil {
		t.Fatalf("failed to load xml document. %v", fileName)
	}

	root := doc.GetRoot()
	if root == nil {
		t.Fatalf("failed to get root element from xml document")
	}
	//walkElement(t, root, 0)

	// Write to file
	folder := "_output"
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		t.Fatalf("failed to create folder %v. %v", folder, err)
	}

	outputFileName := path.Join(folder, "log.xml")
	outputFile, err := os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create output file. %v", outputFileName)
	}
	defer outputFile.Close()
	outputFile.Truncate(0)

	if err := doc.Save(DefaultOutputConfig, outputFile); err != nil {
		t.Errorf("failed to save xml dom to file. %v", err)
	}
}
