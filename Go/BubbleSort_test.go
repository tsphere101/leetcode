package main

import "testing"

func Test_bubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		// Test case 1
		{
			name: "Test case 1",
			args: args{
				arr: []int{3, 1, 8, 2, 5},
			},
		},

		// Test case 2
		{
			name: "Test case 2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
		},

		// Test case 3

		{name: "Test case 3",

			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
		},

		// Test case 4
		{
			name: "Test case 4",
			args: args{
				arr: []int{1, 1, 1, 1, 1},
			},
		},

		// Test case 5

		{name: "Test case 5",

			args: args{
				arr: []int{},
			},
		},

		// Test case 6
		{
			name: "Test case 6",
			args: args{
				arr: generateRandomArray(10),
			},
		},

		// Test case 7
		{
			name: "Test case 7",
			args: args{
				arr: generateRandomArray(100),
			},
		},

		// Test case 8
		{
			name: "Test case 8",
			args: args{
				arr: generateRandomArray(1000),
			},
		},

		// Test case 9
		{
			name: "Test case 9",
			args: args{
				arr: generateRandomArray(10000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.arr)
		})
	}
}
