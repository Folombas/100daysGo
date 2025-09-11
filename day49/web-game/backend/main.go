package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"backend/game"
)

var (
	games = make(map[string]*game.Game)
	mutex = &sync.Mutex{}
)

func startGame(w http.ResponseWriter, r *http.Request) {
	g := game.NewGame()
	mutex.Lock()
	games[g.ID] = g
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"game_id": g.ID})
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")
	if gameID == "" {
		http.Error(w, "Отсутствует game_id", http.StatusBadRequest)
		return
	}
	guessStr := r.URL.Query().Get("guess")
	if guessStr == "" {
		http.Error(w, "Отсутствует guess", http.StatusBadRequest)
		return
	}
	guess, err := strconv.Atoi(guessStr)
	if err != nil {
		http.Error(w, "Неверный формат числа", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	g, exists := games[gameID]
	mutex.Unlock()
	if !exists {
		http.Error(w, "Игра не найдена", http.StatusNotFound)
		return
	}

	message, won := g.Guess(guess)
	response := map[string]interface{}{
		"message":           message,
		"won":               won,
		"attempts":          g.Attempts,
		"remaining_attempts": g.MaxAttempts - g.Attempts,
	}

	if won || g.Attempts >= g.MaxAttempts {
		response["secret"] = g.Secret
		mutex.Lock()
		delete(games, gameID)
		mutex.Unlock()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/start", startGame)
	http.HandleFunc("/guess", guessHandler)

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	fmt.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
