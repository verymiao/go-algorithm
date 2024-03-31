package LeetCode

import "testing"

func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{1, 3, 4, 2, 2}, want: 2},
		{name: "case2", args: []int{3, 1, 3, 4, 2}, want: 3},
		{name: "case3", args: []int{3, 3, 3, 3, 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
