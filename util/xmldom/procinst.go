/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

type procInstImpl struct {
	nodeEx
	target string
	inst   string
}

func (pi *procInstImpl) GetTarget() string {
	return pi.target
}

func (pi *procInstImpl) SetTarget(target string) {
	pi.target = target
}

func (pi *procInstImpl) GetInst() string {
	return pi.inst
}

func (pi *procInstImpl) SetInst(inst string) {
	pi.inst = inst
}

func newProcInst() *procInstImpl {
	return &procInstImpl{
		nodeEx: newNode(),
	}
}
