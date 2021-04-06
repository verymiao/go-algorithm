package LeetCode

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	inputs := [][]int{[]int{1, 1, 1, 2, 2, 3}, []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{}}
	outputs := []int{5, 7, 0}

	for i := 0; i < len(inputs); i++ {
		result := removeDuplicates(inputs[i])
		if !reflect.DeepEqual(result, outputs[i]) {
			t.Errorf("input: %v, expect: %v, but result: %v\n", inputs[i], outputs[i], result)
		}
	}
}
