/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package xmldom

type commentImpl struct {
	nodeEx
	comment string
}

func (cmt *commentImpl) GetComment() string {
	return cmt.comment
}

func (cmt *commentImpl) SetComment(comment string) {
	cmt.comment = comment
}

func newComment() *commentImpl {
	return &commentImpl{
		nodeEx: newNode(),
	}
}
