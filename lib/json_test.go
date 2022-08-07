package lib

import (
	"testing"
)

func TestJSONFormat_TabIndent(t *testing.T) {
	type fields struct {
		fn func(s string) string
	}

	type args struct {
		n map[string]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "TabIndent",
			prepare: func(f *fields) {
				f.fn = func(s string) string {
					return NewJSONFormat().TabIndent(s)
				}
			},
			args: args{
				n: map[string]string{
					`{"alg":"HS256","typ":"JWT"}`: `{
	"alg": "HS256",
	"typ": "JWT"
}
`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got := f.fn(k); got != v {
					t.Errorf("%s() = %v, want %v", tt.name, got, v)
				}
			}
		})
	}
}

func TestJSONFormat_TwoSpacesIndent(t *testing.T) {
	type fields struct {
		fn func(s string) string
	}

	type args struct {
		n map[string]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "TwoSpacesIndent",
			prepare: func(f *fields) {
				f.fn = func(s string) string {
					return NewJSONFormat().TwoSpacesIndent(s)
				}
			},
			args: args{
				n: map[string]string{
					`{"alg":"HS256","typ":"JWT"}`: `{
  "alg": "HS256",
  "typ": "JWT"
}
`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got := f.fn(k); got != v {
					t.Errorf("%s() = %v, want %v", tt.name, got, v)
				}
			}
		})
	}
}

func TestJSONFormat_FourSpacesIndent(t *testing.T) {
	type fields struct {
		fn func(s string) string
	}

	type args struct {
		n map[string]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "FourSpacesIndent",
			prepare: func(f *fields) {
				f.fn = func(s string) string {
					return NewJSONFormat().FourSpacesIndent(s)
				}
			},
			args: args{
				n: map[string]string{
					`{"alg":"HS256","typ":"JWT"}`: `{
    "alg": "HS256",
    "typ": "JWT"
}
`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got := f.fn(k); got != v {
					t.Errorf("%s() = %v, want %v", tt.name, got, v)
				}
			}
		})
	}
}

func TestJSONFormat_Minify(t *testing.T) {
	type fields struct {
		fn func(s string) string
	}

	type args struct {
		n map[string]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "Minify",
			prepare: func(f *fields) {
				f.fn = func(s string) string {
					return NewJSONFormat().Minify(s)
				}
			},
			args: args{
				n: map[string]string{
					`{"alg":"HS256","typ":"JWT"}`: `{"alg":"HS256","typ":"JWT"}`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got := f.fn(k); got != v {
					t.Errorf("%s() = %v, want %v", tt.name, got, v)
				}
			}
		})
	}
}

func TestJSONFormat_IsJSON(t *testing.T) {
	type fields struct {
		fn func(s string) bool
	}

	type args struct {
		n map[string]bool
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "IsJSON",
			prepare: func(f *fields) {
				f.fn = func(s string) bool {
					return NewJSONFormat().IsJSON(s)
				}
			},
			args: args{
				n: map[string]bool{
					`{"alg":"HS256","typ":"JWT"}`: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got := f.fn(k); got != v {
					t.Errorf("%s() = %v, want %v", tt.name, got, v)
				}
			}
		})
	}
}
