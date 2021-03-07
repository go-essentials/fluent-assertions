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
	"testing"
)

// UT: TestEqualTypesFails ensures that assert.EqualTypes is correctly implemented.
func TestEqualTypesFails(t *testing.T) {
	// CONSTANTS.
	utName := "UT should fail with the correct message when relfect.TypeOf(expected) is NOT equal to " +
		"reflect.TypeOf(actual)."
	expected := ""
	actual := true
	expFailureMsg := "\n" +
		"UT Name:         UT Name\n" +
		"Expected (type): string\n" +
		"Actual (type):   bool"

	// SETUP.
	utCaseWrapper := &utCaseWrapper{}

	// ACT.
	EqualTypes(utCaseWrapper, "UT Name", expected, actual)

	// ASSERT.
	if utCaseWrapper.failedWithMsg != expFailureMsg {
		t.Fatalf("\n"+
			"UT Name:            %v\n"+
			"Expected (message): \n    %v\n\n"+
			"Actual (message):   \n    %v\n\n",
			utName, fmtFailureMsg(expFailureMsg), fmtFailureMsg(utCaseWrapper.failedWithMsg))
	}
}

// UT: TestEqualTypes ensures that assert.EqualTypes is correctly implemented.
func TestEqualTypes(t *testing.T) {
	// CONSTANTS.
	utName := "UT should NOT fail when relfect.TypeOf(expected) is equal to reflect.TypeOf(actual)."
	expected := true
	actual := false

	// SETUP.
	utCaseWrapper := &utCaseWrapper{}

	// ACT.
	EqualTypes(utCaseWrapper, "UT Name", expected, actual)

	// ASSERT.
	if utCaseWrapper.failedWithMsg != "" {
		t.Fatalf("\n"+
			"UT Name:            %v\n"+
			"Expected (message): <nil>\n"+
			"Actual (message):   \n    %v\n\n",
			utName, fmtFailureMsg(utCaseWrapper.failedWithMsg))
	}
}
