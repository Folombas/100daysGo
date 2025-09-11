package game

import (
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	ID           string
	Secret       int
	Attempts     int
	MaxAttempts  int
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		ID:          strconv.FormatInt(time.Now().UnixNano(), 10),
		Secret:      rand.Intn(100) + 1,
		Attempts:    0,
		MaxAttempts: 10,
	}
}

func (g *Game) Guess(n int) (string, bool) {
	g.Attempts++
	if g.Attempts > g.MaxAttempts {
		return "Превышено количество попыток", false
	}
	if n < g.Secret {
		return "Загаданное число больше", false
	}
	if n > g.Secret {
		return "Загаданное число меньше", false
	}
	return "Вы угадали!", true
}
