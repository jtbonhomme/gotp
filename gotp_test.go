package gotp

import (
	"errors"
	"fmt"
	"testing"
)

func TestPrefix0(t *testing.T) {
	var tests = []struct {
		otp  string
		want string
	}{
		{"1234", "001234"},
		{"01234", "001234"},
		{"001234", "001234"},
		{"123456", "123456"},
		{"1234567", "1234567"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("test %s", tt.otp)
		t.Run(testname, func(t *testing.T) {
			got := prefix0(tt.otp)

			if got != tt.want {
				t.Errorf("got result %s, want %s", got, tt.want)
			}
		})
	}
}

func TestGetHOTPToken(t *testing.T) {
	var tests = []struct {
		secret   string
		interval int64
		want     string
		err      error
	}{
		{"KZAUYVKFGA======", 52709041, "097417", nil},
		{"KZAUYVKFGE======", 52709041, "451039", nil},
		{"KZAUYVKFGI======", 52709041, "530217", nil},
		{"KZAUYVKFGM======", 52709041, "752372", nil},
		{"KZAUYVKFGQ======", 52709041, "521971", nil},
		{"KZAUYVKFGA======", 52709051, "526980", nil},
		{"KZAUYVKFGE======", 52709051, "674321", nil},
		{"KZAUYVKFGI======", 52709051, "744542", nil},
		{"KZAUYVKFGM======", 52709051, "430197", nil},
		{"KZAUYVKFGQ======", 52709051, "493393", nil},
		{"this is not a correct secret", 52709051, "493393", errors.New("wrong format")},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("test %s", tt.secret)
		t.Run(testname, func(t *testing.T) {
			got, err := getHOTPToken(tt.secret, tt.interval)
			if err != nil && tt.err == nil {
				t.Errorf("got error %s, want nil", err.Error())
			}
			if err == nil && tt.err != nil {
				t.Errorf("got no error but want one")
			}
			if err == nil && got != tt.want {
				t.Errorf("got result %s, want %s", got, tt.want)
			}
		})
	}
}
