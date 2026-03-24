package config_test

import (
	"testing"

	"github.com/equinor/radix-cli/pkg/config"
)

func TestIsValidContext(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		context string
		want    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.IsValidContext(tt.context)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("IsValidContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
