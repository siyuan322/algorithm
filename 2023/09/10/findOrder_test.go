package main

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "case02",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "case03",
			args: args{
				numCourses:    1,
				prerequisites: nil,
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
