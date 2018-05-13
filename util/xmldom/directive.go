/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

type directiveImpl struct {
	nodeEx
	directive string
}

func (di *directiveImpl) GetDirective() string {
	return di.directive
}

func (di *directiveImpl) SetDirective(directive string) {
	di.directive = directive
}

func newDirective() *directiveImpl {
	return &directiveImpl{
		nodeEx: newNode(),
	}
}
