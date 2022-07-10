package lib_test

import (
	"testing"

	"github.com/cage1016/alfred-devtoys/lib"
)

func TestEncode_Base64(t *testing.T) {
	type args struct {
		input map[string]string
	}
	tests := []struct {
		name string
		e    *lib.Encode
		args args
	}{
		{
			name: "Base64",
			e:    &lib.Encode{},
			args: args{
				input: map[string]string{
					"hello world": "aGVsbG8gd29ybGQ=",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				if got := lib.NewEncoder().Base64(k); got != v {
					t.Errorf("Encode.Base64() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestEncode_URL(t *testing.T) {
	type args struct {
		input map[string]string
	}
	tests := []struct {
		name string
		e    *lib.Encode
		args args
	}{
		{
			name: "URL",
			e:    &lib.Encode{},
			args: args{
				input: map[string]string{
					"hello world": "hello%20world",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				if got := lib.NewEncoder().URL(k); got != v {
					t.Errorf("Encode.URL() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestEncode_HTML(t *testing.T) {
	type args struct {
		input map[string]string
	}
	tests := []struct {
		name string
		e    *lib.Encode
		args args
	}{
		{
			name: "HTML",
			e:    &lib.Encode{},
			args: args{
				input: map[string]string{
					"<html>hello world</html>": "&lt;html&gt;hello world&lt;/html&gt;",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				if got := lib.NewEncoder().HTML(k); got != v {
					t.Errorf("Encode.HTML() = %v, want %v", got, v)
				}
			}
		})
	}
}
