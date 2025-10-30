package main

import "github.com/stretchr/testify/mock"

// MockUserRepository мок для репозитория пользователей
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Save(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepository) UpdateBalance(userID string, newBalance float64) error {
    args := m.Called(userID, newBalance)
    return args.Error(0)
}

// StubUserRepository стаб для репозитория пользователей
type StubUserRepository struct {
    FindByIDFunc      func(id string) (*User, error)
    SaveFunc          func(user *User) error
    UpdateBalanceFunc func(userID string, newBalance float64) error
}

func (s *StubUserRepository) FindByID(id string) (*User, error) {
    if s.FindByIDFunc != nil {
        return s.FindByIDFunc(id)
    }
    // Возвращаем заглушку по умолчанию
    return &User{
        ID:      id,
        Name:    "Тестовый пользователь",
        Email:   "test@example.com",
        Balance: 100.0,
    }, nil
}

func (s *StubUserRepository) Save(user *User) error {
    if s.SaveFunc != nil {
        return s.SaveFunc(user)
    }
    return nil
}

func (s *StubUserRepository) UpdateBalance(userID string, newBalance float64) error {
    if s.UpdateBalanceFunc != nil {
        return s.UpdateBalanceFunc(userID, newBalance)
    }
    return nil
}
