package lib

import (
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestJWTdecode(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name string
		args args
		want *jwt.Token
	}{
		{
			name: "Decode",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhIjoiYiJ9.9RfgR0OhCFw1pz-g9gLzDEuFSRe1sgsqedGz5e4MkWc",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := JWTdecode(tt.args.tokenString); !reflect.DeepEqual(got, tt.want) && err != nil {
				t.Errorf("JWTdecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
