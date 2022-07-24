package lib

import "testing"

type HashFn func(string) string

func TestMD5(t *testing.T) {
	type fields struct {
		c HashFn
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
			name: "test md5",
			prepare: func(f *fields) {
				f.c = MD5
			},
			args: args{
				n: map[string]string{
					"255":                 "fe131d7f5a6b38b23cc967316c13dae2",
					"9223372036854775807": "15767b252275cf5107bba9267b88e787",
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
					t.Errorf("MD5() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	type fields struct {
		c HashFn
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
			name: "test sha1",
			prepare: func(f *fields) {
				f.c = SHA1
			},
			args: args{
				n: map[string]string{
					"255":                 "3028f51407d83338f72f994bc283572452a877de",
					"9223372036854775807": "458b642b137e2c76e0b746c6fa43e64c3d4c47f1",
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
					t.Errorf("SHA1() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	type fields struct {
		c HashFn
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
			name: "test sha256",
			prepare: func(f *fields) {
				f.c = SHA256
			},
			args: args{
				n: map[string]string{
					"255":                 "9556b82499cc0aaf86aee7f0d253e17c61b7ef73d48a295f37d98f08b04ffa7f",
					"9223372036854775807": "b34a1c30a715f6bf8b7243afa7fab883ce3612b7231716bdcbbdc1982e1aed29",
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
					t.Errorf("SHA256() = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	type fields struct {
		c HashFn
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
			name: "test sha512",
			prepare: func(f *fields) {
				f.c = SHA512
			},
			args: args{
				n: map[string]string{
					"255":                 "b84abbb04904e63955cf7b9def018fb974c71e690fbdc8fc56dc02fe5a974821ade3aea25e0658f1aae869330960befaaf7425ecfef6b137a046794263c3a4f0",
					"9223372036854775807": "85a4f762e59771074e103a5b1fedf1f6c262814d8a87cffb08a5082b4be0de43a7b186ab402cb919fb7d6a4dba1b57b0c2a06071d11bb83ee3f1977abbcf59e9",
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
					t.Errorf("SHA512() = %v, want %v", got, v)
				}
			}
		})
	}
}
