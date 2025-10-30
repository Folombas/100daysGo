package main

import "fmt"

// OrderService основной сервис приложения
type OrderService struct {
    paymentService PaymentService
    userRepo       UserRepository
}

func NewOrderService(paymentService PaymentService, userRepo UserRepository) *OrderService {
    return &OrderService{
        paymentService: paymentService,
        userRepo:       userRepo,
    }
}

// ProcessOrder обрабатывает заказ пользователя
func (s *OrderService) ProcessOrder(userID string, amount float64) (string, error) {
    // 1. Находим пользователя
    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return "", fmt.Errorf("пользователь не найден: %w", err)
    }

    // 2. Проверяем достаточность средств
    if user.Balance < amount {
        return "", fmt.Errorf("недостаточно средств: текущий баланс %.2f, требуется %.2f", user.Balance, amount)
    }

    // 3. Обрабатываем платеж
    paymentID, err := s.paymentService.ProcessPayment(amount, "RUB")
    if err != nil {
        return "", fmt.Errorf("ошибка обработки платежа: %w", err)
    }

    // 4. Обновляем баланс пользователя
    newBalance := user.Balance - amount
    if err := s.userRepo.UpdateBalance(userID, newBalance); err != nil {
        // В реальном приложении здесь была бы компенсирующая транзакция
        return "", fmt.Errorf("ошибка обновления баланса: %w", err)
    }

    return paymentID, nil
}

// RefundOrder обрабатывает возврат заказа
func (s *OrderService) RefundOrder(paymentID string, userID string) error {
    // 1. Возвращаем платеж
    if err := s.paymentService.RefundPayment(paymentID); err != nil {
        return fmt.Errorf("ошибка возврата платежа: %w", err)
    }

    // 2. Получаем статус платежа для возврата
    status, err := s.paymentService.GetPaymentStatus(paymentID)
    if err != nil {
        return fmt.Errorf("ошибка получения статуса платежа: %w", err)
    }

    // 3. Находим пользователя
    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return fmt.Errorf("пользователь не найден: %w", err)
    }

    // 4. Возвращаем средства (в реальном приложении здесь была бы логика расчета суммы возврата)
    newBalance := user.Balance + 100.0 // упрощенная логика
    if err := s.userRepo.UpdateBalance(userID, newBalance); err != nil {
        return fmt.Errorf("ошибка возврата средств: %w", err)
    }

    fmt.Printf("✅ Возврат обработан. Статус платежа: %s\n", status)
    return nil
}
