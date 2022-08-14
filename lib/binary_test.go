package lib

import "testing"

type Fn func(string) string

func TestBinToOct(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = BinToOct
			},
			args: args{
				n: map[string]string{
					"11111111": "377",
					"111111111111111111111111111111111111111111111111111111111111111": "777777777777777777777",
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
				if got := f.c(k); got != v {
					t.Errorf("BinToOct() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestBinToDec(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = BinToDec
			},
			args: args{
				n: map[string]string{
					"11111111": "255",
					"111111111111111111111111111111111111111111111111111111111111111": "9223372036854775807",
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
				if got := f.c(k); got != v {
					t.Errorf("BinToDec() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestBinToHex(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = BinToHex
			},
			args: args{
				n: map[string]string{
					"11111111": "FF",
					"111111111111111111111111111111111111111111111111111111111111111": "7FFFFFFFFFFFFFFF",
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
				if got := f.c(k); got != v {
					t.Errorf("BinToHex() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestDecToBin(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = DecToBin
			},
			args: args{
				n: map[string]string{
					"255":                 "11111111",
					"9223372036854775807": "111111111111111111111111111111111111111111111111111111111111111",
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
				if got := f.c(k); got != v {
					t.Errorf("DecToBin() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestDecToOct(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = DecToOct
			},
			args: args{
				n: map[string]string{
					"255":                 "377",
					"9223372036854775807": "777777777777777777777",
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
				if got := f.c(k); got != v {
					t.Errorf("DecToOct() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestDecToHex(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = DecToHex
			},
			args: args{
				n: map[string]string{
					"255":                 "FF",
					"9223372036854775807": "7FFFFFFFFFFFFFFF",
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
				if got := f.c(k); got != v {
					t.Errorf("DecToHex() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestOctToBin(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = OctToBin
			},
			args: args{
				n: map[string]string{
					"377":                   "11111111",
					"777777777777777777777": "111111111111111111111111111111111111111111111111111111111111111",
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
				if got := f.c(k); got != v {
					t.Errorf("OctToBin() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestOctToDec(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = OctToDec
			},
			args: args{
				n: map[string]string{
					"377":                   "255",
					"777777777777777777777": "9223372036854775807",
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
				if got := f.c(k); got != v {
					t.Errorf("OctToDec() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestOctToHex(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = OctToHex
			},
			args: args{
				n: map[string]string{
					"377":                   "FF",
					"777777777777777777777": "7FFFFFFFFFFFFFFF",
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
				if got := f.c(k); got != v {
					t.Errorf("OctToHex() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestHexToBin(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = HexToBin
			},
			args: args{
				n: map[string]string{
					"FF":               "11111111",
					"7FFFFFFFFFFFFFFF": "111111111111111111111111111111111111111111111111111111111111111",
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
				if got := f.c(k); got != v {
					t.Errorf("HexToBin() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestHexToOct(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = HexToOct
			},
			args: args{
				n: map[string]string{
					"FF":               "377",
					"7FFFFFFFFFFFFFFF": "777777777777777777777",
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
				if got := f.c(k); got != v {
					t.Errorf("HexToOct() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestHexToDec(t *testing.T) {
	type fields struct {
		c Fn
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
			name: "test",
			prepare: func(f *fields) {
				f.c = HexToDec
			},
			args: args{
				n: map[string]string{
					"FF":               "255",
					"7FFFFFFFFFFFFFFF": "9223372036854775807",
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
				if got := f.c(k); got != v {
					t.Errorf("HexToDec() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestBaseConvert(t *testing.T) {
	type Input struct {
		Number           string
		FromBase, ToBase int
	}

	type fields struct {
		c func(number string, fromBase, toBase int) string
	}

	type args struct {
		n map[Input]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "test",
			prepare: func(f *fields) {
				f.c = baseConvert
			},
			args: args{
				n: map[Input]string{
					{"10", 10, 2}: "1010",
					{"10", -1, 2}: "",
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
				if got := f.c(k.Number, k.FromBase, k.ToBase); got != v {
					t.Errorf("HexToDec() = %v, want %v", got, v)
				}
			}
		})
	}
}
