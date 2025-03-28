package handler

import "testing"

func TestParseDigits(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		wantErr bool
		wantOut string
	}{{
		name:    "success",
		input:   "/pin 1234",
		wantOut: "1234",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res, err := parseDigits(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("SavePin() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if res != tt.wantOut {
					t.Errorf("SavePin() = %v, want %v", res, tt.wantOut)
				}
			}
		})
	}
}
