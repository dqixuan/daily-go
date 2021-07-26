package dp

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		interval [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			args:args{interval: [][]int{{1,3}, {2,6}, {8, 10}}},
			want: [][]int{{1,6}, {8,10}},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.interval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}