package main

import (
	"errors"
	"math"
	"strings"
	"time"
	"unicode/utf8"
)

var currentTime = TimeNow()

var (
	strokeLengthError = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func currentDayOfTheWeek() string {
	weekday := currentTime.Weekday()
	switch int(weekday) {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	default:
		return "Воскресенье"
	}
}

func dayOrNight() string {
	h := currentTime.Hour()
	m := currentTime.Minute()
	if h >= 10 && h <= 21 || h == 22 && m == 0 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	return int(math.Abs((float64(currentTime.Weekday())) - 5))
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	if strings.EqualFold(answer, currentDayOfTheWeek()) {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, strokeLengthError
	}
	return strings.EqualFold(answer, dayOrNight()), nil
}