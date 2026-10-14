package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Ошибка выполнения функции WalkingSpentCalories")
	}
	average_speed := MeanSpeed(steps, height, duration)
	minute := duration.Minutes()
	count_calories := ((weight * average_speed * minute) / minInH) * walkingCaloriesCoefficient
	return count_calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Ошибка выполнения функции RunningSpentCalories")
	}
	average_speed := MeanSpeed(steps, height, duration)
	minute := duration.Minutes()
	calories_burned := (weight * average_speed * minute) / minInH
	return calories_burned, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	if dist == 0 {
		return 0
	}
	average_speed := dist / duration.Hours()
	return average_speed
}

func Distance(steps int, height float64) float64 {
	len_step := height * stepLengthCoefficient
	count_step := (float64(steps) * len_step) / mInKm
	return count_step
}
