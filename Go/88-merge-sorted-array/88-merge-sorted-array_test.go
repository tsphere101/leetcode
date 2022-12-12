package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
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
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},

		{
			name: "Test 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},

		{
			name: "Test 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},

		{
			name: "Test 4",
			args: args{
				nums1: []int{4, 5, 6, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},

		{
			name: "Test 5",
			args: args{
				nums1: []int{0, 0, 0, 0, 0},
				m:     0,
				nums2: []int{1, 2, 3, 4, 5},
				n:     5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.args.nums1
			merge(got, tt.args.m, tt.args.nums2, tt.args.n)
			t.Log(tt.args.nums1)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}

		})
	}
}
