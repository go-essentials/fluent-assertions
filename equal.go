// =====================================================================================================================
// = LICENSE:       Copyright (c) 2021 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package assert implements functions for making assertions.
package assert

import (
	"reflect"
)

// EqualTypes fails t when reflect.TypeOf(expected) is NOT equal to reflect.TypeOf(actual).
func EqualTypes(t utCase, tName string, expected, actual interface{}) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Fatalf("\nUT Name:         %v\nExpected (type): %v\nActual (type):   %v", tName, reflect.TypeOf(expected),
			reflect.TypeOf(actual))
	}
}

// Equal fails t when expected is NOT equal to actual.
func Equal(t utCase, tName string, expected, actual interface{}) {
	EqualTypes(t, tName, expected, actual)

	if expected != actual {
		t.Fatalf("\nUT Name:  %v\nExpected: %v\nActual:   %v", tName, expected, actual)
	}
}
