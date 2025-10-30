package main

import "testing"

// MockUserRepository мок для репозитория пользователей (упрощенная версия)
type MockUserRepository struct {
    FindByIDFunc      func(id string) (*User, error)
    SaveFunc          func(user *User) error
    UpdateBalanceFunc func(userID string, newBalance float64) error

    // Для отслеживания вызовов
    FindByIDCalls      []string
    UpdateBalanceCalls []balanceCall
}

type balanceCall struct {
    userID    string
    newBalance float64
}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
    m.FindByIDCalls = append(m.FindByIDCalls, id)
    if m.FindByIDFunc != nil {
        return m.FindByIDFunc(id)
    }
    return &User{
        ID:      id,
        Name:    "Mock User",
        Email:   "mock@example.com",
        Balance: 100.0,
    }, nil
}

func (m *MockUserRepository) Save(user *User) error {
    if m.SaveFunc != nil {
        return m.SaveFunc(user)
    }
    return nil
}

func (m *MockUserRepository) UpdateBalance(userID string, newBalance float64) error {
    m.UpdateBalanceCalls = append(m.UpdateBalanceCalls, balanceCall{userID, newBalance})
    if m.UpdateBalanceFunc != nil {
        return m.UpdateBalanceFunc(userID, newBalance)
    }
    return nil
}

// Assertions для ручной проверки
func (m *MockUserRepository) AssertFindByIDCalled(t testing.TB, userID string) {
    found := false
    for _, call := range m.FindByIDCalls {
        if call == userID {
            found = true
            break
        }
    }
    if !found {
        t.Errorf("Expected FindByID to be called with %s", userID)
    }
}

func (m *MockUserRepository) AssertUpdateBalanceCalled(t testing.TB, userID string, expectedBalance float64) {
    found := false
    for _, call := range m.UpdateBalanceCalls {
        if call.userID == userID && call.newBalance == expectedBalance {
            found = true
            break
        }
    }
    if !found {
        t.Errorf("Expected UpdateBalance to be called with %s and balance %f", userID, expectedBalance)
    }
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
