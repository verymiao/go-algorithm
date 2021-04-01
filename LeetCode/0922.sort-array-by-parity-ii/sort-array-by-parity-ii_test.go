package LeetCode

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	inputs := [][]int{[]int{4, 2, 5, 7}, []int{}}
	outputs := [][]int{[]int{4, 5, 2, 7}, []int{}}

	for i := 0; i < len(inputs); i++ {
		result := sortArrayByParityII(inputs[i])
		if !reflect.DeepEqual(result, outputs[i]) {
			t.Errorf("input: %v, expect: %v, but result: %v\n", inputs[i], outputs[i], result)
		}
	}
}
