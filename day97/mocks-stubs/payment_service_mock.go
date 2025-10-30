package main

import "github.com/stretchr/testify/mock"

// MockPaymentService мок для платежного сервиса
type MockPaymentService struct {
    mock.Mock
}

func (m *MockPaymentService) ProcessPayment(amount float64, currency string) (string, error) {
    args := m.Called(amount, currency)
    return args.String(0), args.Error(1)
}

func (m *MockPaymentService) RefundPayment(paymentID string) error {
    args := m.Called(paymentID)
    return args.Error(0)
}

func (m *MockPaymentService) GetPaymentStatus(paymentID string) (string, error) {
    args := m.Called(paymentID)
    return args.String(0), args.Error(1)
}

// StubPaymentService стаб для платежного сервиса
type StubPaymentService struct {
    ProcessPaymentFunc func(amount float64, currency string) (string, error)
    RefundPaymentFunc  func(paymentID string) error
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
