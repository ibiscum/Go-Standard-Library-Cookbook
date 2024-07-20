package main

import (
	"testing"
)

func TestIndentByRune(t *testing.T) {
	type args struct {
		input  string
		indent int
		r      rune
	}
	args_v := args{"Hello", 4, '.'}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test4",
			args_v,
			"....Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndentByRune(tt.args.input, tt.args.indent, tt.args.r); got != tt.want {
				t.Errorf("IndentByRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndent(t *testing.T) {
	type args struct {
		input  string
		indent int
	}
	args_v := args{"Hello", 4}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test5",
			args_v,
			"    Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Indent(tt.args.input, tt.args.indent); got != tt.want {
				t.Errorf("Indent() = %v, want %v", got, tt.want)
			}
		})
	}
}
