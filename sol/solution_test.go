package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	for idx := 0; idx < b.N; idx++ {
		twoSum(numbers, target)
	}
}
func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "numbers = [2,7,11,15], target = 9",
			args: args{numbers: []int{2, 7, 11, 15}, target: 9},
			want: []int{1, 2},
		},
		{
			name: "numbers = [2,3,4], target = 6",
			args: args{numbers: []int{2, 3, 4}, target: 6},
			want: []int{1, 3},
		},
		{
			name: "numbers = [-1,0], target = -1",
			args: args{numbers: []int{-1, 0}, target: -1},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
