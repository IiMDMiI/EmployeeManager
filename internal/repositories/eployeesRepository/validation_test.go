package eployeesRepository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr bool
	}{
		{
			name:    "Valid phone number with plus",
			phone:   "+1234567890",
			wantErr: false,
		},
		{
			name:    "Invalid phone number with letters",
			phone:   "+123abc4567",
			wantErr: true,
		},
		{
			name:    "Empty phone number",
			phone:   "",
			wantErr: true,
		},
		{
			name:    "Too short phone number",
			phone:   "+123",
			wantErr: true,
		},
		{
			name:    "Too long phone number",
			phone:   "+12345678901234567890",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := valid.ValidatePhone(tt.phone)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
