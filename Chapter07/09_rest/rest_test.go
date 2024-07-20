package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestCity_toJson(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Location string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			"test1",
			fields{"1", "Berlin", "Germany"},
			`{"name":"Berlin","location":"Germany"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := City{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Location: tt.fields.Location,
			}
			if got := c.toJson(); got != tt.want {
				t.Errorf("City.toJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeCity(t *testing.T) {
	type args struct {
		r io.Reader
	}
	r_v := bytes.NewBuffer([]byte(`{"name":"Berlin","location":"Germany"}`))
	t_v := &args{r_v}
	c_v := &City{ Name: "Berlin", Location: "Germany"}

	tests := []struct {
		name    string
		args    args
		want    City
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test2", *t_v, *c_v, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeCity(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
