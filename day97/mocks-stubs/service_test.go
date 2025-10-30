package main

import (
    "errors"
    "testing"
)

func TestOrderService_ProcessOrder_Success(t *testing.T) {
    // Arrange
    paymentMock := &MockPaymentService{
        ProcessPaymentFunc: func(amount float64, currency string) (string, error) {
            return "pay_mock_123", nil
        },
    }

    userRepoMock := &MockUserRepository{
        FindByIDFunc: func(id string) (*User, error) {
            return &User{
                ID:      "user_123",
                Name:    "Гоша",
                Email:   "gosha@example.com",
                Balance: 1000.0,
            }, nil
        },
        UpdateBalanceFunc: func(userID string, newBalance float64) error {
            return nil
        },
    }

    service := NewOrderService(paymentMock, userRepoMock)

    // Act
    paymentID, err := service.ProcessOrder("user_123", 100.0)

    // Assert
    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }
    if paymentID != "pay_mock_123" {
        t.Errorf("Expected payment ID 'pay_mock_123', but got: %s", paymentID)
    }

    // Проверяем вызовы методов
    userRepoMock.AssertFindByIDCalled(t, "user_123")
    userRepoMock.AssertUpdateBalanceCalled(t, "user_123", 900.0)
    paymentMock.AssertProcessPaymentCalled(t, 1)
}

func TestOrderService_ProcessOrder_InsufficientFunds(t *testing.T) {
    // Arrange
    paymentMock := &MockPaymentService{}

    userRepoMock := &MockUserRepository{
        FindByIDFunc: func(id string) (*User, error) {
            return &User{
                ID:      "user_123",
                Name:    "Гоша",
                Email:   "gosha@example.com",
                Balance: 50.0, // Меньше требуемой суммы
            }, nil
        },
    }

    service := NewOrderService(paymentMock, userRepoMock)

    // Act
    paymentID, err := service.ProcessOrder("user_123", 100.0)

    // Assert
    if err == nil {
        t.Error("Expected error for insufficient funds, but got none")
    }
    if paymentID != "" {
        t.Errorf("Expected empty payment ID, but got: %s", paymentID)
    }
    if err.Error() != "недостаточно средств: текущий баланс 50.00, требуется 100.00" {
        t.Errorf("Unexpected error message: %v", err)
    }
}

func TestOrderService_ProcessOrder_PaymentFailure(t *testing.T) {
    // Arrange
    paymentMock := &MockPaymentService{
        ProcessPaymentFunc: func(amount float64, currency string) (string, error) {
            return "", errors.New("ошибка банка")
        },
    }

    userRepoMock := &MockUserRepository{
        FindByIDFunc: func(id string) (*User, error) {
            return &User{
                ID:      "user_123",
                Name:    "Гоша",
                Email:   "gosha@example.com",
                Balance: 1000.0,
            }, nil
        },
    }

    service := NewOrderService(paymentMock, userRepoMock)

    // Act
    paymentID, err := service.ProcessOrder("user_123", 100.0)

    // Assert
    if err == nil {
        t.Error("Expected error for payment failure, but got none")
    }
    if paymentID != "" {
        t.Errorf("Expected empty payment ID, but got: %s", paymentID)
    }
}

// Простая функция Assert для базовых проверок
func assertTrue(t testing.TB, condition bool, message string) {
    if !condition {
        t.Error(message)
    }
}

func assertEqual(t testing.TB, expected, actual interface{}, message string) {
    if expected != actual {
        t.Errorf("%s: expected %v, got %v", message, expected, actual)
    }
}
