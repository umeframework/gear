/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPropertiesWriter(t *testing.T) {
	// Create properties
	data := map[Key]Value{
		"boot.timeout": 30,
		"boot.default": `multi(0)disk(0)rdisk(0)partition(1)\WINNT`,

		`编号`:    12345678,
		`名称`:    `第12345678号设置项目`,
		`複雑な設定`: "a\\b\u0000 A\t\r\n ;#=:",
	}
	p := New()
	ToProperties(p, data)

	// Test utility functions
	data2 := make(map[Key]Value)
	ToMap(data2, p)
	if len(data) != len(data2) {
		t.Errorf("lengths are different. original = %v, new = %v", len(data), len(data2))
	} else {
		for key, value := range data {
			if value2, ok := data2[key]; value != value2 || !ok {
				t.Errorf("values are different. key = %v, original = %v, new = %v, found = %v", key, value, value2, ok)
			}
		}
	}

	// Write to file
	folder := "_output"
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		t.Fatalf("failed to create folder %v. %v", folder, err)
	}

	filePath := filepath.Join(folder, fmt.Sprintf("%v-%v.properties", os.Getpid(), time.Now().Nanosecond()))

	func(filePath string) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for writing. %v", filePath, err)
		}
		defer file.Close()

		writer := NewWriter()
		if err := writer.Write(p, file); err != nil {
			t.Fatalf("failed to save to ini file (%s). %v", filePath, err)
		}
	}(filePath)

	// Read from file
	func(filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
		}
		defer file.Close()

		reader := NewReader()
		p2 := New()
		if err := reader.Read(p2, file); err != nil {
			t.Errorf("failed to read file %v. %v", filePath, err)
		}
	}(filePath)
}

func TestPropertiesReader_Invalid(t *testing.T) {
	filePaths := []string{
		`testdata/invalid-1.properties`,
		`testdata/invalid-2.properties`,
		`testdata/invalid-3.properties`,
	}
	for _, filePath := range filePaths {
		func(filePath string) {
			file, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
			}
			defer file.Close()

			reader := NewReader()
			p2 := New()
			if err := reader.Read(p2, file); err == nil {
				t.Errorf("incorrectly read invalid file %v. %v", filePath, p2)
			}
		}(filePath)
	}
}

func TestPropertiesReader_Extended(t *testing.T) {
	var testFunc = func(filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
		}
		defer file.Close()

		reader := NewReader()
		p2 := New()
		if err := reader.Read(p2, file); err != nil {
			t.Errorf("failed to read invalid file %v. %v", filePath, err)
		}
	}
	testFunc(`testdata/test-1.properties`)
}
