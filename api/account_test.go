package api

import (
	"fmt"
	"testing"
)

func TestFormatCurrency(t *testing.T) {
	tests := []struct {
		amount   int
		currency CurrencyType
		expected string
	}{
		{1000, GBP, "£10.00"},
		{1500, USD, "$15.00"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %s", tt.amount, tt.currency), func(t *testing.T) {
			result := formatCurrency(tt.amount, tt.currency)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
