package main

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestOrderService_ProcessOrder_Success(t *testing.T) {
    // Arrange
    paymentMock := &MockPaymentService{}
    userRepoMock := &MockUserRepository{}

    user := &User{
        ID:      "user_123",
        Name:    "Гоша",
        Email:   "gosha@example.com",
        Balance: 1000.0,
    }

    // Настраиваем ожидания для моков
    userRepoMock.On("FindByID", "user_123").Return(user, nil)
    paymentMock.On("ProcessPayment", 100.0, "RUB").Return("pay_mock_123", nil)
    userRepoMock.On("UpdateBalance", "user_123", 900.0).Return(nil)

    service := NewOrderService(paymentMock, userRepoMock)

    // Act
    paymentID, err := service.ProcessOrder("user_123", 100.0)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "pay_mock_123", paymentID)
    paymentMock.AssertExpectations(t)
    userRepoMock.AssertExpectations(t)
}

func TestOrderService_ProcessOrder_InsufficientFunds(t *testing.T) {
    // Arrange
    paymentMock := &MockPaymentService{}
    userRepoMock := &MockUserRepository{}

    user := &User{
        ID:      "user_123",
        Name:    "Гоша",
        Email:   "gosha@example.com",
        Balance: 50.0, // Меньше требуемой суммы
    }

    userRepoMock.On("FindByID", "user_123").Return(user, nil)

    service := NewOrderService(paymentMock, userRepoMock)

    // Act
    paymentID, err := service.ProcessOrder("user_123", 100.0)

    // Assert
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "недостаточно средств")
    assert.Empty(t, paymentID)
    userRepoMock.AssertExpectations(t)
}

func TestOrderService_ProcessOrder_PaymentFailure(t *testing.T) {
    // Arrange
    paymentMock := &MockPaymentService{}
    userRepoMock := &MockUserRepository{}

    user := &User{
        ID:      "user_123",
        Name:    "Гоша",
        Email:   "gosha@example.com",
        Balance: 1000.0,
    }

    userRepoMock.On("FindByID", "user_123").Return(user, nil)
    paymentMock.On("ProcessPayment", 100.0, "RUB").Return("", errors.New("ошибка банка"))

    service := NewOrderService(paymentMock, userRepoMock)

    // Act
    paymentID, err := service.ProcessOrder("user_123", 100.0)

    // Assert
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "ошибка обработки платежа")
    assert.Empty(t, paymentID)
    paymentMock.AssertExpectations(t)
    userRepoMock.AssertExpectations(t)
}

func TestOrderService_RefundOrder_Success_WithStub(t *testing.T) {
    // Arrange
    paymentStub := &StubPaymentService{
        RefundPaymentFunc: func(paymentID string) error {
            return nil
        },
        GetPaymentStatusFunc: func(paymentID string) (string, error) {
            return "refunded", nil
        },
    }

    userRepoStub := &StubUserRepository{
        FindByIDFunc: func(id string) (*User, error) {
            return &User{
                ID:      id,
                Balance: 500.0,
            }, nil
        },
        UpdateBalanceFunc: func(userID string, newBalance float64) error {
            return nil
        },
    }

    service := NewOrderService(paymentStub, userRepoStub)

    // Act
    err := service.RefundOrder("pay_123", "user_123")

    // Assert
    assert.NoError(t, err)
}

// Пример теста с табличными данными
func TestOrderService_ProcessOrder_TableDriven(t *testing.T) {
    tests := []struct {
        name          string
        userID        string
        amount        float64
        userBalance   float64
        expectedError string
        setupMocks    func(payment *MockPaymentService, userRepo *MockUserRepository)
    }{
        {
            name:        "успешный платеж",
            userID:      "user_123",
            amount:      100.0,
            userBalance: 1000.0,
            setupMocks: func(payment *MockPaymentService, userRepo *MockUserRepository) {
                userRepo.On("FindByID", "user_123").Return(&User{ID: "user_123", Balance: 1000.0}, nil)
                payment.On("ProcessPayment", 100.0, "RUB").Return("pay_123", nil)
                userRepo.On("UpdateBalance", "user_123", 900.0).Return(nil)
            },
        },
        {
            name:          "недостаточно средств",
            userID:        "user_123",
            amount:        1500.0,
            userBalance:   1000.0,
            expectedError: "недостаточно средств",
            setupMocks: func(payment *MockPaymentService, userRepo *MockUserRepository) {
                userRepo.On("FindByID", "user_123").Return(&User{ID: "user_123", Balance: 1000.0}, nil)
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            paymentMock := &MockPaymentService{}
            userRepoMock := &MockUserRepository{}

            tt.setupMocks(paymentMock, userRepoMock)

            service := NewOrderService(paymentMock, userRepoMock)
            _, err := service.ProcessOrder(tt.userID, tt.amount)

            if tt.expectedError != "" {
                assert.Error(t, err)
                assert.Contains(t, err.Error(), tt.expectedError)
            } else {
                assert.NoError(t, err)
            }

            paymentMock.AssertExpectations(t)
            userRepoMock.AssertExpectations(t)
        })
    }
}
