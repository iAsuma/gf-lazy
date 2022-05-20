package qutil

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func InSlice(haystack []interface{}, needle interface{}, safe ...bool) bool {
	if g.IsEmpty(haystack) {
		return false
	}

	newArr := garray.NewFrom(haystack, safe...)
	if newArr.Contains(needle) {
		return true
	}

	return false
}

func InSliceString(haystack []string, needle string, safe ...bool) bool {
	if g.IsEmpty(haystack) {
		return false
	}

	newArr := garray.NewStrArray(safe...)
	newArr.SetArray(haystack)
	if newArr.Contains(needle) {
		return true
	}

	return false
}

func InSliceStringI(haystack []string, needle string, safe ...bool) bool {
	if g.IsEmpty(haystack) {
		return false
	}

	newArr := garray.NewStrArray()
	newArr.SetArray(haystack)
	if newArr.ContainsI(needle) {
		return true
	}

	return false
}
