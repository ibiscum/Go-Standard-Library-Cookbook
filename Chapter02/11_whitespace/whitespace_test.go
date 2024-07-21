package main

import "testing"

func Test_pad(t *testing.T) {
	type args struct {
		input  string
		padLen int
		align  string
	}
	args_v := args{"need space", 14, "LEFT"}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test7",
			args_v,
			"    need space",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pad(tt.args.input, tt.args.padLen, tt.args.align); got != tt.want {
				t.Errorf("pad() = %v, want %v", got, tt.want)
			}
		})
	}
}
