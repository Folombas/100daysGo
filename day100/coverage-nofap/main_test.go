package main

import "testing"

func TestNoFapWarrior(t *testing.T) {
	warrior := NewWarrior("Тестовый Участник Челленджа", 25)

	// Тест начального состояния
	if warrior.DaysClean != 0 {
		t.Errorf("Ожидалось 0 дней чистоты, получено %d", warrior.DaysClean)
	}

	if warrior.EnergyLevel != 50.0 {
		t.Errorf("Ожидался уровень энергии 50.0, получено %.1f", warrior.EnergyLevel)
	}
}

func TestAddCleanDay(t *testing.T) {
	warrior := NewWarrior("Тестовый Участник", 25)

	warrior.AddCleanDay()
	if warrior.DaysClean != 1 {
		t.Errorf("Ожидалось 1 день чистоты, получено %d", warrior.DaysClean)
	}

	if warrior.EnergyLevel != 52.5 {
		t.Errorf("Ожидался уровень энергии 52.5, получено %.1f", warrior.EnergyLevel)
	}
}

func TestGetCoverageReport(t *testing.T) {
	warrior := NewWarrior("Тестовый Воин", 25)

	// 10 дней чистоты из 30
	for i := 0; i < 10; i++ {
		warrior.AddCleanDay()
	}

	report := warrior.GetCoverageReport()
	expected := "Покрытие жизни: 33.3%"

	if len(report) < len(expected) {
		t.Errorf("Отчет coverage не соответствует ожиданиям")
	}
}

func TestSkillsAcquisition(t *testing.T) {
	warrior := NewWarrior("Тестовый Участник", 25)

	// 6 дней должно дать 2 навыка (каждые 3 дня)
	for i := 0; i < 6; i++ {
		warrior.AddCleanDay()
	}

	if len(warrior.GoSkills) != 2 {
		t.Errorf("Ожидалось 2 навыка, получено %d", len(warrior.GoSkills))
	}
}
