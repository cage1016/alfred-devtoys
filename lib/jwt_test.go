package lib

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTdecode(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name      string
		args      args
		wantError error
	}{
		{
			name: "Decode",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhIjoiYiJ9.9RfgR0OhCFw1pz-g9gLzDEuFSRe1sgsqedGz5e4MkWc",
			},
			wantError: nil,
		},
		{
			name: "Decode",
			args: args{
				tokenString: "invalid-token-string",
			},
			wantError: errors.New("[jwt: invalid token string] Invalid token"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := JWTdecode(tt.args.tokenString); tt.wantError != nil {
				assert.Error(t, err)
			}
		})
	}
}
