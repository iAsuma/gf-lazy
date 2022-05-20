package qutil

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

// InSlice 判断元素是否在数组中
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

// InSliceString 判断字符串是否在数组中
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

// InSliceStringI 判断字符串是否在数组中，不区分大小写
func InSliceStringI(haystack []string, needle string, safe ...bool) bool {
	if g.IsEmpty(haystack) {
		return false
	}

	newArr := garray.NewStrArray(safe...)
	newArr.SetArray(haystack)
	if newArr.ContainsI(needle) {
		return true
	}

	return false
}

// InSliceInt 判断int是否在数组中，不区分大小写
func InSliceInt(haystack []int, needle int, safe ...bool) bool {
	if g.IsEmpty(haystack) {
		return false
	}

	newArr := garray.NewIntArray(safe...)
	newArr.SetArray(haystack)
	if newArr.Contains(needle) {
		return true
	}

	return false
}
