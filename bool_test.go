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

// UT: TestTrueFails ensures that assert.True is correctly implemented.
func TestTrueFails(t *testing.T) {
	// CONSTANTS.
	utName := "UT should fail with the correct message when val is false."
	val := false
	expFailureMsg := "\n" +
		"UT Name:  UT Name\n" +
		"Expected: true\n" +
		"Actual:   false"

	// SETUP.
	utCaseWrapper := &utCaseWrapper{}

	// ACT.
	True(utCaseWrapper, "UT Name", val)

	// ASSERT.
	if utCaseWrapper.failedWithMsg != expFailureMsg {
		t.Fatalf("\n"+
			"UT Name:            %v\n"+
			"Expected (message): \n    %v\n\n"+
			"Actual (message):   \n    %v\n\n",
			utName, fmtFailureMsg(expFailureMsg), fmtFailureMsg(utCaseWrapper.failedWithMsg))
	}
}

// UT: TestTrue ensures that assert.True is correctly implemented.
func TestTrue(t *testing.T) {
	// CONSTANTS.
	utName := "UT should NOT fail when val is true."
	val := true

	// SETUP.
	utCaseWrapper := &utCaseWrapper{}

	// ACT.
	True(utCaseWrapper, "UT Name", val)

	// ASSERT.
	if utCaseWrapper.failedWithMsg != "" {
		t.Fatalf("\n"+
			"UT Name:            %v\n"+
			"Expected (message): <nil>\n"+
			"Actual (message):   \n    %v\n\n",
			utName, fmtFailureMsg(utCaseWrapper.failedWithMsg))
	}
}

// UT: TestFalseFails ensures that assert.False is correctly implemented.
func TestFalseFails(t *testing.T) {
	// CONSTANTS.
	utName := "UT should fail with the correct message when val is true."
	val := true
	expFailureMsg := "\n" +
		"UT Name:  UT Name\n" +
		"Expected: false\n" +
		"Actual:   true"

	// SETUP.
	utCaseWrapper := &utCaseWrapper{}

	// ACT.
	False(utCaseWrapper, "UT Name", val)

	// ASSERT.
	if utCaseWrapper.failedWithMsg != expFailureMsg {
		t.Fatalf("\n"+
			"UT Name:            %v\n"+
			"Expected (message): \n    %v\n\n"+
			"Actual (message):   \n    %v\n\n",
			utName, fmtFailureMsg(expFailureMsg), fmtFailureMsg(utCaseWrapper.failedWithMsg))
	}
}

// UT: TestFalse ensures that assert.False is correctly implemented.
func TestFalse(t *testing.T) {
	// CONSTANTS.
	utName := "UT should NOT fail when val is false."
	val := false

	// SETUP.
	utCaseWrapper := &utCaseWrapper{}

	// ACT.
	False(utCaseWrapper, "UT Name", val)

	// ASSERT.
	if utCaseWrapper.failedWithMsg != "" {
		t.Fatalf("\n"+
			"UT Name:            %v\n"+
			"Expected (message): <nil>\n"+
			"Actual (message):   \n    %v\n\n",
			utName, fmtFailureMsg(utCaseWrapper.failedWithMsg))
	}
}
