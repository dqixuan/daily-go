package array

import "testing"

func Test_searchFirstEqualElement(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args:args{
				nums:   []int{1, 2, 3, 3, 3, 5},
				target: 3,
			},
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchFirstEqualElement(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchFirstEqualElement() = %v, want %v", got, tt.want)
			}
		})
	}
}