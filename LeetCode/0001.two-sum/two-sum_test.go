package LeetCode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type testStruct struct {
		target int
		input  []int
		output []int
	}

	testdatas := []testStruct{
		// 正常数据
		testStruct{
			target: 9,
			input:  []int{2, 7, 11, 15},
			output: []int{0, 1}},
		testStruct{
			target: 6,
			input:  []int{3, 2, 4},
			output: []int{1, 2}},
		// 异常数据
		testStruct{
			target: 6,
			input:  []int{10, 11, 12},
			output: []int{}},
	}

	for _, testdata := range testdatas {
		result := twoSum(testdata.input, testdata.target)
		if !reflect.DeepEqual(result, testdata.output) {
			t.Errorf("input: %v,target: %v, expect: %v, but result: %v\n", testdata.input, testdata.target, testdata.output, result)
		}
	}
}
