package main

import (
    "testing" // Добавляем этот импорт
)

// MockPaymentService мок для платежного сервиса (упрощенная версия)
type MockPaymentService struct {
    ProcessPaymentFunc func(amount float64, currency string) (string, error)
    RefundPaymentFunc  func(paymentID string) error
    GetPaymentStatusFunc func(paymentID string) (string, error)

    // Для отслеживания вызовов
    ProcessPaymentCalls []callArgs
    RefundPaymentCalls  []string
}

type callArgs struct {
    amount   float64
    currency string
}

func (m *MockPaymentService) ProcessPayment(amount float64, currency string) (string, error) {
    m.ProcessPaymentCalls = append(m.ProcessPaymentCalls, callArgs{amount, currency})
    if m.ProcessPaymentFunc != nil {
        return m.ProcessPaymentFunc(amount, currency)
    }
    return "mock_payment_id", nil
}

func (m *MockPaymentService) RefundPayment(paymentID string) error {
    m.RefundPaymentCalls = append(m.RefundPaymentCalls, paymentID)
    if m.RefundPaymentFunc != nil {
        return m.RefundPaymentFunc(paymentID)
    }
    return nil
}

func (m *MockPaymentService) GetPaymentStatus(paymentID string) (string, error) {
    if m.GetPaymentStatusFunc != nil {
        return m.GetPaymentStatusFunc(paymentID)
    }
    return "completed", nil
}

// Assertions для ручной проверки
func (m *MockPaymentService) AssertProcessPaymentCalled(t testing.TB, times int) {
    if len(m.ProcessPaymentCalls) != times {
        t.Errorf("Expected ProcessPayment to be called %d times, but was called %d times",
            times, len(m.ProcessPaymentCalls))
    }
}

func (m *MockPaymentService) AssertRefundPaymentCalled(t testing.TB, paymentID string) {
    found := false
    for _, call := range m.RefundPaymentCalls {
        if call == paymentID {
            found = true
            break
        }
    }
    if !found {
        t.Errorf("Expected RefundPayment to be called with %s", paymentID)
    }
}

// StubPaymentService стаб для платежного сервиса
type StubPaymentService struct {
    ProcessPaymentFunc   func(amount float64, currency string) (string, error)
    RefundPaymentFunc    func(paymentID string) error
    GetPaymentStatusFunc func(paymentID string) (string, error)
}

func (s *StubPaymentService) ProcessPayment(amount float64, currency string) (string, error) {
    if s.ProcessPaymentFunc != nil {
        return s.ProcessPaymentFunc(amount, currency)
    }
    return "stub_payment_id", nil
}

func (s *StubPaymentService) RefundPayment(paymentID string) error {
    if s.RefundPaymentFunc != nil {
        return s.RefundPaymentFunc(paymentID)
    }
    return nil
}

func (s *StubPaymentService) GetPaymentStatus(paymentID string) (string, error) {
    if s.GetPaymentStatusFunc != nil {
        return s.GetPaymentStatusFunc(paymentID)
    }
    return "completed", nil
}
