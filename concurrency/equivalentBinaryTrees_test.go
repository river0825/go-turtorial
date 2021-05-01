package main

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "same tree return true", args: args{t1: tree.New(10), t2: tree.New(10)}, want: true},
		{name: "same tree return true", args: args{t1: tree.New(10), t2: tree.New(20)}, want: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}
