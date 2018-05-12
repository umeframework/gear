/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package config

import (
	"github.com/umeframework/gear/core/collection"
	"github.com/umeframework/gear/core/properties"
)

type configImpl struct {
	properties.KeyValuePair
	properties.Properties
	parent   Config
	children collection.List
}

func New() Config {
	obj := configImpl{
		KeyValuePair: properties.NewKeyValuePair(),
		Properties:   properties.New(),
		parent:       nil,
		children:     collection.NewLinkedList(),
	}
	return &obj
}

func (cfg *configImpl) GetParent() Config {
	return cfg.parent
}

func (cfg *configImpl) SetParent(config Config) {
	cfg.parent = config
}

func (cfg *configImpl) NumChildren() int {
	return cfg.children.Size()
}

func (cfg *configImpl) GetChild(index int) Config {
	child, _ := cfg.children.GetAt(index)
	return child.(Config)
}

func (cfg *configImpl) SetChild(index int, config Config) {
	cfg.children.SetAt(index, config)
}

func (cfg *configImpl) AddChild(index int, config Config) {
	cfg.children.AddAt(index, config)
}

func (cfg *configImpl) RemoveChild(index int) {
	cfg.children.RemoveAt(index)
}

func (cfg *configImpl) ClearChildren() {
	cfg.children.Clear()
}
