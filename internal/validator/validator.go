package validator

import (
	"strings"
	"unicode"
)

// Validate принимает номер карты и возвращает:
// valid — true, если номер корректен и соответствует Visa или MasterCard
// cardType — "Visa", "MasterCard" или "Unknown"
func Validate(cardNumber string) (bool, string) {
	// Убираем пробелы и дефисы
	cardNumber = sanitize(cardNumber)

	if !isNumeric(cardNumber) {
		return false, "Unknown"
	}

	cardType := getCardType(cardNumber)
	if cardType == "Unknown" {
		return false, cardType
	}

	if !luhnCheck(cardNumber) {
		return false, cardType
	}

	return true, cardType
}

// sanitize убирает пробелы и дефисы из строки
func sanitize(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "-", "")
}

// isNumeric проверяет, что строка состоит только из цифр
func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// getCardType определяет тип карты по первым цифрам
func getCardType(cardNumber string) string {
	if len(cardNumber) == 0 {
		return "Unknown"
	}

	switch {
	case strings.HasPrefix(cardNumber, "4"):
		return "Visa"
	case hasMasterCardPrefix(cardNumber):
		return "MasterCard"
	default:
		return "Unknown"
	}
}

// hasMasterCardPrefix проверяет MasterCard по первым цифрам (51-55, 2221-2720)
func hasMasterCardPrefix(cardNumber string) bool {
	if len(cardNumber) < 2 {
		return false
	}

	prefix2 := cardNumber[:2]
	prefix4 := ""
	if len(cardNumber) >= 4 {
		prefix4 = cardNumber[:4]
	}

	// 51–55
	if prefix2 >= "51" && prefix2 <= "55" {
		return true
	}

	// 2221–2720
	if prefix4 >= "2221" && prefix4 <= "2720" {
		return true
	}

	return false
}

// luhnCheck реализует алгоритм Луна для проверки номера карты
func luhnCheck(cardNumber string) bool {
	sum := 0
	double := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}
