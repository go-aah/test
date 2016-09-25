// Copyright (c) Jeevanandam M (https://github.com/jeevatkm)
// go-aah/test source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package assert

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// Equal asserts given two values are equal.
// If it's not equal, it logs the error trace.
// It supports all the values supported by `reflect.DeepEqual`
func Equal(t *testing.T, expected, got interface{}) {
	if !equal(expected, got) {
		fail(t, 3, "Expected [%v], got [%v]", expected, got)
	}
}

// Equalf asserts given two values are equal.
// If it's not equal, it logs the error trace.
// It supports all the values supported by `reflect.DeepEqual`
func Equalf(t *testing.T, expected, got interface{}, msg string, args ...interface{}) {
	if !equal(expected, got) {
		fail(t, 3, msg, args...)
	}
}

// NotEqual asserts given two values are not equal.
// If it's equal, it logs the error trace.
// It supports all the values supported by `reflect.DeepEqual`
func NotEqual(t *testing.T, expected, got interface{}) {
	if equal(expected, got) {
		fail(t, 3, "Expected [%v], got [%v]", expected, got)
	}
}

// NotEqualf asserts given two values are not equal.
// If it's equal, it logs the error trace.
// It supports all the values supported by `reflect.DeepEqual`
func NotEqualf(t *testing.T, expected, got interface{}, msg string, args ...interface{}) {
	if equal(expected, got) {
		fail(t, 3, msg, args...)
	}
}

// Nil asserts the given value is `nil`. If it's not nil,
// it log the error trace
func Nil(t *testing.T, v interface{}) {
	if !isNil(v) {
		fail(t, 3, "Expected [nil], got [%v]", v)
	}
}

// Nilf asserts the given value is `nil`. If it's not nil,
// it log the error trace
func Nilf(t *testing.T, v interface{}, msg string, args ...interface{}) {
	if !isNil(v) {
		fail(t, 3, msg, args...)
	}
}

// NotNil asserts the given value is not `nil`. If it's nil,
// it log the error trace
func NotNil(t *testing.T, v interface{}) {
	if isNil(v) {
		fail(t, 3, "Expected [%v], got [nil]", v)
	}
}

// NotNilf asserts the given value is not `nil`. If it's nil,
// it log the error trace
func NotNilf(t *testing.T, v interface{}, msg string, args ...interface{}) {
	if isNil(v) {
		fail(t, 3, msg, args...)
	}
}

// True asserts the given value is true, if not true
// it log the error trace
func True(t *testing.T, v interface{}) {
	if !equal(true, v) {
		fail(t, 3, "Expected [true], got [%v]", v)
	}
}

// Truef asserts the given value is true, if not true
// it log the error trace
func Truef(t *testing.T, v interface{}, msg string, args ...interface{}) {
	if !equal(true, v) {
		fail(t, 3, msg, args...)
	}
}

// False asserts the given value is false, if not false
// it log the error trace
func False(t *testing.T, v interface{}) {
	if !equal(false, v) {
		fail(t, 3, "Expected [false], got [%v]", v)
	}
}

// Falsef asserts the given value is false, if not false
// it log the error trace
func Falsef(t *testing.T, v interface{}, msg string, args ...interface{}) {
	if !equal(false, v) {
		fail(t, 3, msg, args...)
	}
}

// Fail reports fail through and logs the error trace
func Fail(t *testing.T, msg string, args ...interface{}) {
	fail(t, 3, msg, args...)
}

// FailOnError asserts given `error` if it's not nil. It reports
// the error trace
func FailOnError(t *testing.T, err error, msg string) {
	if err != nil {
		fail(t, 3, msg+": %v", err)
	}
}

// FailOnErrorf asserts given `error` if it's not nil. It reports
// the error trace
func FailOnErrorf(t *testing.T, err error, msg string, args ...interface{}) {
	if err != nil {
		fail(t, 3, msg, args...)
	}
}

// FailNowOnError asserts given `error` if it's not nil. It reports
// the error trace and fails the test
func FailNowOnError(t *testing.T, err error, msg string) {
	if err != nil {
		fail(t, 3, msg+": %v", err)
		t.FailNow()
	}
}

// FailNowOnErrorf asserts given `error` if it's not nil. It reports
// the error trace and fails the test
func FailNowOnErrorf(t *testing.T, err error, msg string, args ...interface{}) {
	if err != nil {
		fail(t, 3, msg, args...)
		t.FailNow()
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Unexported methods
//___________________________________

func fail(t *testing.T, calldepth int, msg string, args ...interface{}) {
	if len(args) > 0 {
		t.Errorf("\nError Trace: \n%v: %v", fetchCallerInfo(calldepth), fmt.Sprintf(msg, args...))
	} else {
		t.Errorf("\nError Trace: \n%v: %v", fetchCallerInfo(calldepth), msg)
	}
}

func equal(expected, got interface{}) bool {
	return reflect.DeepEqual(expected, got)
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && rv.IsNil() {
		return true
	}

	return false
}

func fetchCallerInfo(calldepth int) string {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%v:%v", file[len(goSrcPath())+1:], line)
}

func goSrcPath() string {
	gopath := build.Default.GOPATH
	if len(gopath) == 0 {
		return ""
	}

	var currentGoPath string
	workingDir, _ := os.Getwd()
	goPathList := filepath.SplitList(gopath)
	for _, path := range goPathList {
		if strings.HasPrefix(strings.ToLower(workingDir), strings.ToLower(path)) {
			currentGoPath = path
			break
		}

		path, _ = filepath.EvalSymlinks(path)
		if len(path) > 0 && strings.HasPrefix(strings.ToLower(workingDir), strings.ToLower(path)) {
			currentGoPath = path
			break
		}
	}

	if len(currentGoPath) == 0 {
		currentGoPath = goPathList[0]
	}

	return filepath.Join(currentGoPath, "src")
}
