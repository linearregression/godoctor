package doctor

// This file defines miscellaneous utility methods and structs that are used
// throughout the system.

import (
	"strconv"
)

// An OffsetLength consists of two integers: a 0-based offset and a nonnegative
// length.  An OffsetLength is used to specify a region of a string or file.
// For example, given the string "ABCDEFG", the substring CDE could be
// specified by
//     OffsetLength{offset: 2, length: 3}
type OffsetLength struct {
	offset int
	length int
}

func (o *OffsetLength) String() string {
	return "offset " + strconv.Itoa(o.offset) +
		", length " + strconv.Itoa(o.length)
}

// A TextSelection represents a selection in a text editor.  It consists of a
// filename, the line/column where the selected text begins, and the
// line/column where the text selection ends.  The end line and column must be
// greater than or equal to the start line and column, respectively.
type TextSelection struct {
	filename  string
	startLine int
	startCol  int
	endLine   int
	endCol    int
}