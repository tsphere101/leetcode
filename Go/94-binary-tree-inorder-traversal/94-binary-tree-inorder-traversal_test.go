package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				root: NewNode(1, nil, NewNode(2, NewNode(3, nil, nil), nil)), // 1, 3, 2
			},
			want: []int{1, 3, 2},
		},
		{
			name: "Test 2",
			args: args{
				root: NewNode(2, NewNode(1, nil, nil), nil), // 1, 2
			},
			want: []int{1, 2},
		},
		{
			name: "Test 3",
			args: args{
				root: NewNode(1, nil, nil),
			},
			want: []int{1},
		},
		{
			name: "Test 4",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "Test 5",
			args: args{
				root: NewNode(1, NewNode(2, NewNode(3, NewNode(4, NewNode(5, nil, nil), nil), nil), nil), nil), // 5, 4, 3, 2, 1
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
