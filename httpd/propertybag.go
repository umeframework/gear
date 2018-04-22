package httpd

import (
	"reflect"
	"fmt"
)

type PropertyBagBase struct {
	namedMap map[string]interface{}
	interfaceMap map[reflect.Type]interface{}
}

func (this *PropertyBagBase) GetValue(key string) (interface{}, bool) {
	value, found := this.namedMap[key]
	return value, found
}

func (this *PropertyBagBase) SetValue(key string, value interface{}) {
	this.namedMap[key] = value
}

func (this *PropertyBagBase) DeleteValue(key string) {
	delete(this.namedMap, key)
}

func (this *PropertyBagBase) GetAllKeys() []string {
	keys := make([]string, len(this.namedMap))
	i := 0
	for key := range this.namedMap {
		keys[i] = key
		i++
	}
	return keys
}

func (this *PropertyBagBase) GetInterface(t reflect.Type) (interface{}, bool) {
	value, found := this.interfaceMap[t]
	return value, found
}

func (this *PropertyBagBase) SetInterface(t reflect.Type, value interface{}) {
	fmt.Println("type = ", t)
	this.interfaceMap[t] = value
}

func NewPropertyBag() PropertyBag {
	propertyBag := PropertyBagBase{
		make(map[string]interface{}),
		make(map[reflect.Type]interface{}),
	}
	return &propertyBag
}