package main

// Speaker - интерфейс для животных
type Speaker interface {
    Speak() string
}

// Animal - базовая структура
type Animal struct {
    Species string
    Sound   string
}

// Реализация интерфейса Speaker
func (a Animal) Speak() string {
    return "Я " + a.Species + "! " + a.Sound
}

// AnimalFactory - фабрика для создания животных
func AnimalFactory(species string) Speaker {
    sounds := map[string]string{
        "кошка":    "Мяу!",
        "собака":   "Гав!",
        "корова":   "Му-Му!",
        "утка":     "Кря-кря!",
        "петух":    "Кукареку!",
        "свинья":   "Хрю-Хрю!",
        "мышь":     "Пи-Пи!",
        "лягушка":  "Ква-Ква!",
        "ворона":   "Кар-Кар!",
        "кукушка":  "Ку-Ку!",
    }
    
    if sound, ok := sounds[species]; ok {
        return Animal{Species: species, Sound: sound}
    }
    return Animal{Species: "неизвестное существо", Sound: "???"}
}