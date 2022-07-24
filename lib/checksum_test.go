package lib

import (
	"testing"

	"github.com/cage1016/alfred-devtoys/testdata"
)

type FuncPack struct {
	TheFunc func() string
}

func TestNewCheckSum(t *testing.T) {
	cs, _ := NewCheckSum(testdata.Path("file1.txt"))

	type args struct {
		n map[*FuncPack]string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test checksum",
			args: args{
				n: func() map[*FuncPack]string {
					m := make(map[*FuncPack]string)
					m[&FuncPack{TheFunc: cs.MD5}] = "453d08cc48054b46fa25f7efc8462705"
					m[&FuncPack{TheFunc: cs.SHA1}] = "29f2a8b2bdd5eae69a3225361118864318a878ea"
					m[&FuncPack{TheFunc: cs.SHA256}] = "bb4cf9f8a0ad48df97ce5f996981c85d6ab3a9f177a9c455343cc27c776753b3"
					m[&FuncPack{TheFunc: cs.SHA512}] = "49f9458e0e5e9ace276610d7e6d9f623bbe8d244195adf0bf30a889033395e0220336f4273423b51b54a09733ff714f1bcc024a3baa01a65582b674af614f78a"
					return m
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.n {
				if got := k.TheFunc(); got != v {
					t.Errorf("OctToDec() = %v, want %v", got, v)
				}
			}
		})
	}
}
