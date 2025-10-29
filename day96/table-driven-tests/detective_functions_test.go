package main

import (
	"testing"
)

// TestValidatePassword - table-driven test для функции ValidatePassword
func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
		desc     string
	}{
		{
			name:     "Валидный пароль",
			password: "SecurePass123",
			wantErr:  false,
			desc:     "Пароль с заглавными, строчными буквами и цифрами",
		},
		{
			name:     "Слишком короткий",
			password: "Short1",
			wantErr:  true,
			desc:     "Пароль менее 8 символов",
		},
		{
			name:     "Без заглавных букв",
			password: "nocapital123",
			wantErr:  true,
			desc:     "Пароль без заглавных букв",
		},
		{
			name:     "Без строчных букв",
			password: "NOLOWERCASE123",
			wantErr:  true,
			desc:     "Пароль без строчных букв",
		},
		{
			name:     "Без цифр",
			password: "NoDigitsHere",
			wantErr:  true,
			desc:     "Пароль без цифр",
		},
		{
			name:     "Пустой пароль",
			password: "",
			wantErr:  true,
			desc:     "Пустая строка",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v | Тест: %s", err, tt.wantErr, tt.desc)
			}
		})
	}
}

// TestPasswordStrength - table-driven test для функции PasswordStrength
func TestPasswordStrength(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
		desc     string
	}{
		{
			name:     "Очень сильный пароль",
			password: "Very$ecureP@ss123!",
			want:     "ОЧЕНЬ_СИЛЬНЫЙ",
			desc:     "Длинный пароль с разными типами символов",
		},
		{
			name:     "Сильный пароль",
			password: "StrongPass123",
			want:     "СИЛЬНЫЙ",
			desc:     "Пароль средней длины с буквами и цифрами",
		},
		{
			name:     "Средний пароль",
			password: "Medium12",
			want:     "СРЕДНИЙ",
			desc:     "Короткий пароль с базовыми требованиями",
		},
		{
			name:     "Слабый пароль",
			password: "weak",
			want:     "СЛАБЫЙ",
			desc:     "Очень короткий пароль",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PasswordStrength(tt.password)
			if got != tt.want {
				t.Errorf("PasswordStrength() = %v, want %v | Тест: %s", got, tt.want, tt.desc)
			}
		})
	}
}

// TestValidateEmail - table-driven test для функции ValidateEmail
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
		desc    string
	}{
		{
			name:    "Валидный email",
			email:   "user@example.com",
			wantErr: false,
			desc:    "Стандартный валидный email",
		},
		{
			name:    "Email с поддоменами",
			email:   "user@mail.example.com",
			wantErr: false,
			desc:    "Email с несколькими поддоменами",
		},
		{
			name:    "Без символа @",
			email:   "userexample.com",
			wantErr: true,
			desc:    "Email без символа @",
		},
		{
			name:    "Без домена",
			email:   "user@",
			wantErr: true,
			desc:    "Email без доменной части",
		},
		{
			name:    "Без имени пользователя",
			email:   "@example.com",
			wantErr: true,
			desc:    "Email без локальной части",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v | Тест: %s", err, tt.wantErr, tt.desc)
			}
		})
	}
}

// TestValidateAge - table-driven test для функции ValidateAge
func TestValidateAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr bool
		desc    string
	}{
		{
			name:    "Валидный возраст",
			age:     25,
			wantErr: false,
			desc:    "Стандартный взрослый возраст",
		},
		{
			name:    "Минимальный допустимый",
			age:     18,
			wantErr: false,
			desc:    "Минимальный возраст для взрослых",
		},
		{
			name:    "Несовершеннолетний",
			age:     17,
			wantErr: true,
			desc:    "Возраст меньше 18 лет",
		},
		{
			name:    "Отрицательный возраст",
			age:     -5,
			wantErr: true,
			desc:    "Некорректный отрицательный возраст",
		},
		{
			name:    "Слишком большой возраст",
			age:     151,
			wantErr: true,
			desc:    "Возраст превышает разумные пределы",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAge() error = %v, wantErr %v | Тест: %s", err, tt.wantErr, tt.desc)
			}
		})
	}
}

// TestGenerateSecurePassword - table-driven test для функции GenerateSecurePassword
func TestGenerateSecurePassword(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
		desc    string
	}{
		{
			name:    "Валидная длина",
			length:  12,
			wantErr: false,
			desc:    "Корректная длина пароля",
		},
		{
			name:    "Минимальная длина",
			length:  8,
			wantErr: false,
			desc:    "Минимально допустимая длина",
		},
		{
			name:    "Слишком короткий",
			length:  7,
			wantErr: true,
			desc:    "Длина меньше минимальной",
		},
		{
			name:    "Отрицательная длина",
			length:  -5,
			wantErr: true,
			desc:    "Некорректная отрицательная длина",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateSecurePassword(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSecurePassword() error = %v, wantErr %v | Тест: %s", err, tt.wantErr, tt.desc)
			}
		})
	}
}

