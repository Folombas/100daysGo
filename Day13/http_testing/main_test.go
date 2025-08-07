package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Run("Базовый тест", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/hello?name=Антон", nil)
		rr := httptest.NewRecorder()
		helloHandler(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Ожидался статус 200, получен %d", status)
		}

		expected := "Привет, Антон!"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("Ожидалось: '%s', получено: '%s'", expected, rr.Body.String())
		}
	})

	t.Run("Без параметра", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/hello", nil)
		rr := httptest.NewRecorder()
		helloHandler(rr, req)

		if !strings.Contains(rr.Body.String(), "Привет, Гость!") {
			t.Error("Ожидалось приветствие по умолчанию")
		}
	})
}

func TestEchoHandler(t *testing.T) {
	payload := []byte("Тестовое сообщение")
	req := httptest.NewRequest("POST", "/echo", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	echoHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус 200, получен %d", status)
	}

	expected := "Получено: Тестовое сообщение"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Ожидалось: '%s', получено: '%s'", expected, rr.Body.String())
	}
}

func TestMethodValidation(t *testing.T) {
	// Тест неверного метода для /hello
	req := httptest.NewRequest("POST", "/hello", nil)
	rr := httptest.NewRecorder()
	helloHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Error("Должен возвращать 405 для POST")
	}

	// Тест неверного метода для /echo
	req = httptest.NewRequest("GET", "/echo", nil)
	rr = httptest.NewRecorder()
	echoHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Error("Должен возвращать 405 для GET")
	}
}

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	homeHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус 200, получен %d", status)
	}

	expected := "Добро пожаловать на Day13!"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Ожидалось, что тело ответа содержит '%s'", expected)
	}
}
