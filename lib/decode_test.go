package lib_test

import (
	"testing"

	"github.com/cage1016/alfred-devtoys/lib"
)

func TestDecode_Base64(t *testing.T) {
	type args struct {
		input map[string]string
	}
	tests := []struct {
		name string
		e    *lib.Decode
		args args
	}{
		{
			name: "Base64",
			e:    &lib.Decode{},
			args: args{
				input: map[string]string{
					"aGVsbG8gd29ybGQ=": "hello world",
				},
			},
		},
		{
			name: "Invalid Base64",
			e:    &lib.Decode{},
			args: args{
				input: map[string]string{
					"fake-base64": "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				if got := lib.NewDecoder().Base64(k); got != v {
					t.Errorf("Decode.Base64() = %v, want %v", got, k)
				}
			}
		})
	}
}

func TestDecode_URL(t *testing.T) {
	type args struct {
		input map[string]string
	}
	tests := []struct {
		name string
		e    *lib.Decode
		args args
	}{
		{
			name: "URL",
			e:    &lib.Decode{},
			args: args{
				input: map[string]string{
					"hello%20world": "hello world",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				if got := lib.NewDecoder().URL(k); got != v {
					t.Errorf("Decode.URL() = %v, want %v", got, k)
				}
			}
		})
	}
}

func TestDecode_HTML(t *testing.T) {
	type args struct {
		input map[string]string
	}
	tests := []struct {
		name string
		e    *lib.Decode
		args args
	}{
		{
			name: "HTML",
			e:    &lib.Decode{},
			args: args{
				input: map[string]string{
					"&lt;html&gt;hello world&lt;/html&gt;": "<html>hello world</html>",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				if got := lib.NewDecoder().HTML(k); got != v {
					t.Errorf("Decode.HTML() = %v, want %v", got, k)
				}
			}
		})
	}
}
