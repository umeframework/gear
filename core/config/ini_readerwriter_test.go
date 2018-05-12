/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func walkChildren(t *testing.T, cfg Config) {
	num := cfg.NumChildren()
	for i := 0; i < num; i++ {
		child := cfg.GetChild(i)
		t.Logf("No. %v = %v", i, child.GetKey())
	}
}

func TestIniWriter(t *testing.T) {
	// Create properties
	data := map[string]map[string]interface{}{
		"boot loader": {
			"timeout": 30,
			"default": `multi(0)disk(0)rdisk(0)partition(1)\WINNT`,
		},
		"operation systems": {
			`multi(0)disk(0)rdisk(0)partition(1)\WINNT`: `"Windows 2000 Server" /fastdetect`,
			`multi(0)disk(0)rdisk(0)partition(2)\WINNT`: `"Windows XP Professional" /fastdetect`,
		},
		"中文配置项": {
			`编号`:    12345678,
			`名称`:    `第12345678号设置项目`,
			`複雑な設定`: "a\a\b\u0000 A\t\r\n ;#=:",
		},
	}
	cfg := New()
	for key, value := range data {
		child := New()
		child.SetKey(key)
		for subKey, subValue := range value {
			child.SetProp(subKey, subValue)
		}
		cfg.AddChild(cfg.NumChildren(), child)
	}

	// Write to file
	folder := "_output"
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		t.Fatalf("failed to create folder %v. %v", folder, err)
	}

	filePathBase := filepath.Join(folder, fmt.Sprintf("%v-%v", os.Getpid(), time.Now().Nanosecond()))
	escapedFilePath := filePathBase + "_escaped.ini"
	unescapedFilePath := filePathBase + "_unescaped.ini"

	func(filePath string) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for writing. %v", filePath, err)
		}
		defer file.Close()

		config := IniWriterConfig{
			UseEscape:     true,
			IdentKeyValue: true,
		}
		writer := NewIniWriter(config)
		if err := writer.Write(cfg, file); err != nil {
			t.Fatalf("failed to save to ini file (%s). %v", filePath, err)
		}
	}(escapedFilePath)
	func(filePath string) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for writing. %v", filePath, err)
		}
		defer file.Close()

		config := IniWriterConfig{
			UseEscape:     false,
			IdentKeyValue: false,
		}
		writer := NewIniWriter(config)
		if err := writer.Write(cfg, file); err != nil {
			t.Fatalf("failed to save to ini file (%s). %v", filePath, err)
		}
	}(unescapedFilePath)

	// Read from file
	func(filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
		}
		defer file.Close()

		config := IniReaderConfig{
			UseEscape: true,
		}
		reader := NewIniReader(config)
		cfg2 := New()
		if err := reader.Read(cfg2, file); err != nil {
			t.Errorf("failed to read file %v. %v", filePath, err)
		} else {
			walkChildren(t, cfg)
		}
	}(escapedFilePath)

	func(filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
		}
		defer file.Close()

		config := IniReaderConfig{
			UseEscape: false,
		}
		reader := NewIniReader(config)
		cfg2 := New()
		if err := reader.Read(cfg2, file); err != nil {
			t.Errorf("failed to read file %v. %v", filePath, err)
		} else {
			walkChildren(t, cfg)
		}
	}(unescapedFilePath)
}

func TestIniReader_Invalid(t *testing.T) {
	filePaths := []string{
		`testdata/invalid-1.ini`,
		`testdata/invalid-2.ini`,
		`testdata/invalid-3.ini`,
	}
	for _, filePath := range filePaths {
		func(filePath string) {
			file, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
			}
			defer file.Close()

			config := IniReaderConfig{
				UseEscape: false,
			}
			reader := NewIniReader(config)
			cfg2 := New()
			if err := reader.Read(cfg2, file); err == nil {
				t.Errorf("incorrectly read invalid file %v. %v", filePath, cfg2)
			}
		}(filePath)
	}
}

func TestIniReader_Extended(t *testing.T) {
	var testFunc = func(filePath string) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("failed to open ini file (%s) for reading. %v", filePath, err)
		}
		defer file.Close()

		config := IniReaderConfig{
			UseEscape: false,
		}
		reader := NewIniReader(config)
		cfg2 := New()
		if err := reader.Read(cfg2, file); err != nil {
			t.Errorf("failed to read invalid file %v. %v", filePath, err)
		}
	}
	testFunc(`testdata/test-1.ini`)
}
