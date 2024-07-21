package main

import "testing"

func Test_toCamelCase(t *testing.T) {
	type args struct {
		input string
	}
	args_v := args{"hello_world"}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"tests6",
			args_v,
			"helloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCamelCase(tt.args.input); got != tt.want {
				t.Errorf("toCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
