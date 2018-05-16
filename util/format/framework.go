/*
 * Copyright (c) 2018. All rights reserved.
 * Use of this source code is governed by a Apache
 * license that can be found in the LICENSE file.
 */

package format

type FormatContext interface {
}

type Formatable interface {
	Format(format string) (string, error)
}

type Parsable interface {
	Parse(text string) (interface{}, error)
}

type Formatter interface {
	Format(object interface{}, context FormatContext) (string, error, bool)
}

type Parser interface {
	Parse(text string, context FormatContext) (interface{}, error, bool)
}
