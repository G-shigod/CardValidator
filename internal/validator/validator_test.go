package validator

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		cardNumber string
		wantValid  bool
		wantType   string
	}{
		{"4111111111111111", true, "Visa"},       // валидная Visa
		{"5500000000000004", true, "MasterCard"}, // валидный MasterCard
		{"4111111111111", false, "Visa"},         // неправильный номер Visa
		{"1234567890123456", false, "Unknown"},   // невалидный номер
		{"2221000000000009", true, "MasterCard"}, // новый MasterCard диапазон
	}

	for _, tt := range tests {
		gotValid, gotType := Validate(tt.cardNumber)
		if gotValid != tt.wantValid || gotType != tt.wantType {
			t.Errorf("Validate(%q) = (%v, %q), want (%v, %q)", tt.cardNumber, gotValid, gotType, tt.wantValid, tt.wantType)
		}
	}
}
