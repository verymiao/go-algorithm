package LeetCode

import (
	"reflect"
	"testing"
)

func Test_findMin(t *testing.T) {
	inputs := [][]int{[]int{1, 3, 5}, []int{2, 2, 2, 0, 1}, []int{}}
	outputs := []int{1, 0, 0}

	for i := 0; i < len(inputs); i++ {
		result := findMin(inputs[i])
		if !reflect.DeepEqual(result, outputs[i]) {
			t.Errorf("input: %v, expect: %v, but result: %v\n", inputs[i], outputs[i], result)
		}
	}
}
