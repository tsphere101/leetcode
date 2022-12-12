package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "Test 2",
			args: args{
				x: -123,
			},
			want: -321,
		},

		{
			name: "Test 3",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "Test 4",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},

		{
			name: "Test 6",
			args: args{
				x: -2147483648,
			},
			want: 0,
		},
		{
			name: "Test 7",
			args: args{
				x: 2147483647,
			},
			want: 0,
		},
		{
			name: "Test 8",
			args: args{
				x: 1463847412,
			},
			want: 2147483641,
		},
		{
			name: "Test 9",
			args: args{
				x: 10000000,
			},
			want: 1,
		},
		{
			name: "Test 10",
			args: args{
				x: -10000000,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
