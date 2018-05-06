/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

import "io"

// Key stands for types of key used in properties.
type Key interface{}

// Value stands for types of value used in properties.
type Value interface{}

type Properties interface {
	GetProp(key Key) (Value, bool)
	SetProp(key Key, value Value)
	RemoveProp(key Key)
	ClearProps()
	GetProps() []Key
}

type Reader interface {
	Read(p Properties, r io.Reader) error
}

type Writer interface {
	Write(p Properties, w io.Writer) error
}
