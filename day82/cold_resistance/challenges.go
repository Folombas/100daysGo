package main

import (
	"fmt"
	"errors"
)

func CheckAllSentinels(rs *ResistanceSystem) {
	sentinels := []struct {
		err         error
		name        string
		description string
	}{
		{ErrColdRoom, "Холодная комната", "Спал с открытым балконом"},
		{ErrTVTemptation, "Телевизор", "Мама смотрит сериалы на кухне"},
		{ErrWarmBed, "Теплая кровать", "После холода кровать особенно манит"},
		{ErrCapCutTemptation, "Видеомонтаж", "Воспоминания о тропиках"},
		{ErrSocialMedia, "Соцсети", "Уведомления в телефоне"},
		{ErrHunger, "Голод", "После ужина хочется еще"},
		{ErrTiredness, "Усталость", "Долгий день за рулем"},
		{ErrComfortZone, "Зона комфорта", "Хочется расслабиться"},
		{ErrProcrastination, "Прокрастинация", "Откладывание на потом"},
		{ErrDoubt, "Сомнения", "А стоит ли учить Go?"},
	}

	resisted := 0
	for _, sentinel := range sentinels {
		// Проверяем, было ли это искушение сегодня
		faced := false
		for _, challenge := range rs.ChallengesFaced {
			if errors.Is(challenge.Err.Err, sentinel.err) {
				faced = true
				break
			}
		}

		if faced {
			fmt.Printf("   ✅ %s: ПРЕОДОЛЕНО!\n", sentinel.name)
			resisted++
		} else {
			fmt.Printf("   ⏳ %s: не актуально сегодня\n", sentinel.name)
		}
	}

	fmt.Printf("\n🎯 Итог: %d/%d sentinel errors проверены и преодолены\n",
		resisted, len(sentinels))
}
