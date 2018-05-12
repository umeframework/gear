/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package config

import (
	"github.com/umeframework/gear/core/properties"
	"io"
)

// Config represents a hierarchical config item, which contains children
type Config interface {
	properties.KeyValuePair
	properties.Properties

	GetParent() Config
	SetParent(config Config)

	NumChildren() int
	GetChild(index int) Config
	SetChild(index int, config Config)
	AddChild(index int, config Config)
	RemoveChild(index int)
	ClearChildren()
}

// Reader reads config from io.reader
type Reader interface {
	Read(config Config, r io.Reader) error
}

// Writer writers config to io.writer
type Writer interface {
	Write(config Config, w io.Writer) error
}
