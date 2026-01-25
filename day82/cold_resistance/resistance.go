package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Challenge struct {
	Err         *SentinelError
	Description string
	Power       int
}

type ResistanceSystem struct {
	TemptationsResisted int
	FocusLevel          int
	ChallengesFaced     []*Challenge
	DailyChallenges     []*Challenge
}

func NewResistanceSystem() *ResistanceSystem {
	rand.Seed(time.Now().UnixNano())

	rs := &ResistanceSystem{
		TemptationsResisted: 0,
		FocusLevel:          75,
	}

	// Генерация ежедневных вызовов
	rs.generateDailyChallenges()

	return rs
}

func (rs *ResistanceSystem) generateDailyChallenges() {
	rs.DailyChallenges = []*Challenge{
		{
			Err:         NewSentinelError(ErrTVTemptation, 6, "distraction"),
			Description: "Мама смотрит сериал, звук доносится из кухни",
			Power:       65,
		},
		{
			Err:         NewSentinelError(ErrColdRoom, 4, "environment"),
			Description: "В комнате холодно, пальцы замерзают на клавиатуре",
			Power:       45,
		},
		{
			Err:         NewSentinelError(ErrWarmBed, 8, "comfort"),
			Description: "Кровать выглядит очень inviting после холодного дня",
			Power:       85,
		},
		{
			Err:         NewSentinelError(ErrCapCutTemptation, 7, "distraction"),
			Description: "Вспомнились тропические видео, хочется монтировать",
			Power:       75,
		},
		{
			Err:         NewSentinelError(ErrSocialMedia, 5, "digital"),
			Description: "Телефон мигает уведомлениями из соцсетей",
			Power:       55,
		},
	}
}

func (rs *ResistanceSystem) CheckDailyChallenge() *Challenge {
	if len(rs.DailyChallenges) == 0 {
		return nil
	}

	// 30% шанс столкнуться с вызовом
	if rand.Intn(100) < 30 {
		index := rand.Intn(len(rs.DailyChallenges))
		return rs.DailyChallenges[index]
	}

	return nil
}

func (rs *ResistanceSystem) FaceTemptation(challenge *Challenge) bool {
	fmt.Printf("   🛡️  Сопротивление искушению: %s\n", challenge.Err.Err.Error())

	// Расчет шанса успеха
	successChance := rs.FocusLevel - challenge.Power + 50

	if successChance > 0 && rand.Intn(100) < successChance {
		// Успешное сопротивление
		rs.TemptationsResisted++
		rs.FocusLevel += 10
		rs.ChallengesFaced = append(rs.ChallengesFaced, challenge)

		// Удаляем вызов из списка ежедневных
		for i, c := range rs.DailyChallenges {
			if c == challenge {
				rs.DailyChallenges = append(rs.DailyChallenges[:i], rs.DailyChallenges[i+1:]...)
				break
			}
		}

		return true
	}

	// Неудачное сопротивление
	rs.FocusLevel -= 15
	if rs.FocusLevel < 0 {
		rs.FocusLevel = 0
	}

	return false
}

func (rs *ResistanceSystem) CalculateDopamine() int {
	base := rs.TemptationsResisted * 50
	focusBonus := rs.FocusLevel
	return base + focusBonus
}
