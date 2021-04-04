package LeetCode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type testStruct struct {
		target []int
		nums1  []int
		nums2  []int
		m      int
		n      int
	}

	testdatas := []testStruct{
		// 正常数据
		testStruct{
			target: []int{1, 2, 2, 3, 5, 6},
			nums1:  []int{1, 2, 3, 0, 0, 0},
			nums2:  []int{2, 5, 6},
			m:      3,
			n:      3},
		testStruct{
			target: []int{1},
			nums1:  []int{1},
			nums2:  []int{},
			m:      1,
			n:      0},
	}

	for _, testdata := range testdatas {
		result := merge(testdata.intput, testdata.target)
		if !reflect.DeepEqual(result, testdata.output) {
			t.Errorf("input: %v,target: %v, expect: %v, but result: %v\n", testdata.input, testdata.target, testdata.output, result)
		}
	}
}
