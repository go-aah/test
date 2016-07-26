// Copyright (c) Jeevanandam M (https://github.com/jeevatkm)
// go-aah/test source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package assert

import "testing"

func TestGoPath(t *testing.T) {
	gopath := goSrcPath()
	t.Logf("gopath: %v", gopath)
}

func TestFetchCallerInfo(t *testing.T) {
	callerInfo := fetchCallerInfo(1)
	Equal(t, "github.com/go-aah/test/assert/assert_test.go:15", callerInfo)
}

func TestIsNil(t *testing.T) {
	res1 := isNil(nil)
	NotNil(t, res1)
	Nil(t, nil)
	True(t, res1)

	res2 := isNil("isNil test")
	NotNil(t, res2)
	False(t, res2)
}

func TestInnerEqual(t *testing.T) {
	res1 := equal("expected", "not expected")
	False(t, res1)
	// Just invert of above
	NotEqual(t, "expected", "not expected")

	res2 := equal("expected", "expected")
	True(t, res2)
	// Just invert of above
	Equal(t, "expected", "expected")
}
