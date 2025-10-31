package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// APIClient клиент для работы с API
type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// GetUser получает пользователя по ID
func (c *APIClient) GetUser(id int) (*User, error) {
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))

	resp, err := c.client.Get(c.baseURL + "/user?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ошибка API: %s - %s", resp.Status, string(body))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("ошибка парсинга ответа: %w", err)
	}

	return &user, nil
}

// CreateUser создает нового пользователя
func (c *APIClient) CreateUser(name, email string, age int) (*User, error) {
	userData := map[string]interface{}{
		"name":  name,
		"email": email,
		"age":   age,
	}

	jsonData, err := json.Marshal(userData)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания JSON: %w", err)
	}

	resp, err := c.client.Post(c.baseURL+"/user", "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ошибка API: %s - %s", resp.Status, string(body))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("ошибка парсинга ответа: %w", err)
	}

	return &user, nil
}

// GetAllUsers получает всех пользователей
func (c *APIClient) GetAllUsers() ([]User, error) {
	resp, err := c.client.Get(c.baseURL + "/users")
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ошибка API: %s - %s", resp.Status, string(body))
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("ошибка парсинга ответа: %w", err)
	}

	return users, nil
}
