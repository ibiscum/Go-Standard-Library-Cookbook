package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewSyncList(t *testing.T) {
	cap := 10
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *SyncList
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{cap: cap},
			want: &SyncList{m: sync.Mutex{}, slice: make([]interface{}, cap)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSyncList(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSyncList() = %v, want %v", got, tt.want)
			}
		})
	}
}
