/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package properties

type propertiesImpl struct {
	propDict map[Key]Value
}

// GetProp() retrieves the value of specified key
func (p *propertiesImpl) GetProp(key Key) (Value, bool) {
	value, ok := p.propDict[key]
	return value, ok
}

// SetProp() updates the value of specified key
func (p *propertiesImpl) SetProp(key Key, value Value) {
	p.propDict[key] = value
}

// RemoveProp() removes the value of specified key
func (p *propertiesImpl) RemoveProp(key Key) {
	delete(p.propDict, key)
}

// ClearProps() clears all the values
func (p *propertiesImpl) ClearProps() {
	p.propDict = map[Key]Value{}
}

// GetProps() retrieves all the keys
func (p *propertiesImpl) GetProps() []Key {
	count := len(p.propDict)
	keys := make([]Key, 0, count)
	for key := range p.propDict {
		keys = append(keys, key)
	}

	return keys
}

// New() create a new instance of Properties, which is NOT concurrency safe.
func New() Properties {
	obj := &propertiesImpl{
		propDict: make(map[Key]Value),
	}
	return obj
}
