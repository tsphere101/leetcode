package main

import (
	"reflect"
	"testing"
)

func Test_product(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.

		// Test case 1
		{
			name: "Test case 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{120, 60, 40, 30, 24},
		},

		// Test case 2
		{
			name: "Test case 2",
			args: args{
				arr: []int{3, 2, 1},
			},
			want: []int{2, 3, 6},
		},

		// Test case 3
		{
			name: "Test case 3, Negative numbers",

			args: args{
				arr: []int{-1, -2, -3, 4, -5},
			},
			want: []int{-120, -60, -40, 30, -24},
		},

		// Test case 4
		{
			name: "Test case 4",
			args: args{
				arr: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1},
		},

		// Test case 5
		{
			name: "Test case 5",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},

		// Test case 6
		{
			name: "Test case 6",
			args: args{
				arr: []int{1},
			},
			want: []int{1},
		},

		// Test case 7
		{
			name: "Test case 7",
			args: args{
				arr: []int{1, 2},
			},
			want: []int{2, 1},
		},

		// Test case 8
		{
			name: "Test case 8, contains one zero.",

			args: args{
				arr: []int{1, 2, 0, 4, 5},
			},

			want: []int{0, 0, 40, 0, 0},
		},

		// Test case 9
		{
			name: "Test case 9, contains two zeros.",

			args: args{
				arr: []int{1, 2, 0, 4, 0},
			},

			want: []int{0, 0, 0, 0, 0},
		},

		// Test case 10
		{
			name: "Test case 10, contains one zero and negative values.",

			args: args{
				arr: []int{1, 2, 0, 4, -5},
			},

			want: []int{0, 0, -40, 0, 0},
		},

		// Test case 11
		{
			name: "Test case 11, contains two zeros and negative values.",
			args: args{
				arr: []int{1, 2, 0, -4, 0},
			},

			want: []int{0, 0, 0, 0, 0},
		},

		// Test case 12
		{
			name: "Test case 12, large array.",

			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},

			want: []int{3628800, 1814400, 1209600, 907200, 725760, 604800, 518400, 453600, 403200, 362880},
		},

		// Test case 13
		{
			name: "Test case 13, large array with negative values.",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10},
			},

			want: []int{-3628800, -1814400, -1209600, -907200, -725760, -604800, -518400, -453600, -403200, 362880},
		},

		// Test case 14
		{
			name: "Test case 14, large array with negative values and zeros.",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10, 0},
			},

			want: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -3628800},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := partialProduct(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partialProduct() = %v, want %v", got, tt.want)
			}

		})
	}
}
