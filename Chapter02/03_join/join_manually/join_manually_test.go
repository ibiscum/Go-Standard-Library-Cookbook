package main

import "testing"

func TestJoinWithFunc(t *testing.T) {
	refStringSlice_v := []string{
		"FIRST_NAME = 'Jack' ",
		" INSURANCE_NO = 333444555 ",
		" EFFECTIVE_FROM = SYSDATE"}

	type args struct {
		refStringSlice []string
		joinFunc       JoinFunc
	}

	jF_v := func(p string) string {
		return "AND"
	}
	var args_v = args{refStringSlice_v, jF_v}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test3", args_v, `FIRST_NAME = 'Jack' AND INSURANCE_NO = 333444555 AND EFFECTIVE_FROM = SYSDATE`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinWithFunc(tt.args.refStringSlice, tt.args.joinFunc); got != tt.want {
				t.Errorf("JoinWithFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
