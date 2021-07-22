package LeetCode

import (
	"reflect"
	"testing"
)

func Test_maxArea(t *testing.T) {
	inputs := [][]int{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, []int{1, 1}, []int{4, 3, 2, 1, 4}, []int{1, 2, 1}}
	outputs := []int{49, 1, 16}

	for i := 0; i < len(inputs); i++ {
		result := maxArea(inputs[i])
		if !reflect.DeepEqual(result, outputs[i]) {
			t.Errorf("input: %v, expect: %v, but result: %v\n", inputs[i], outputs[i], result)
		}
	}
}
