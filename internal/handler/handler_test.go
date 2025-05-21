package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateCardHandler(t *testing.T) {
	// Тестируем успешный POST-запрос с валидной картой Visa
	body := []byte(`{"card_number":"4111111111111111"}`)
	req := httptest.NewRequest(http.MethodPost, "/validate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Создаем ResponseRecorder для захвата ответа
	rr := httptest.NewRecorder()

	// Вызываем наш хендлер
	ValidateCard(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ожидали статус %d, получили %d", http.StatusOK, status)
	}

	// Проверяем тело ответа (минимально)
	expected := `"valid":true`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("ожидали, что тело ответа содержит %s, получили %s", expected, rr.Body.String())
	}

	// Тестируем метод GET — должен вернуть ошибку
	req = httptest.NewRequest(http.MethodGet, "/validate", nil)
	rr = httptest.NewRecorder()

	ValidateCard(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("ожидали статус %d для GET, получили %d", http.StatusMethodNotAllowed, status)
	}
}
