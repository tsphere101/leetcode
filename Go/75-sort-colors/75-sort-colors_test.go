package main

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{2, 1, 0, 0, 1, 2},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0},
			},
			want: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{2, 2, 2, 2, 2, 2},
			},
			want: []int{2, 2, 2, 2, 2, 2},
		},
		{
			name: "Test 6",
			args: args{
				nums: []int{0, 1, 2, 0, 1, 2},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Test 7",
			args: args{
				nums: []int{2, 1, 0, 2, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Test 8",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "Test 9",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "Test 10",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "Test 11",
			args: args{
				nums: []int{2},
			},
			want: []int{2},
		},
		{
			name: "Test 12",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			for i := 0; i < len(tt.args.nums); i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
				}
			}
		})
	}
}
