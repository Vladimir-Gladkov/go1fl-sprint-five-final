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

// Вычисляем калории, потраченные при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем на отрицательное занчение
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("data must be positive")
	}
	// Расчитываем и возвращем калории при ходьбе
	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes() / minInH) * walkingCaloriesCoefficient
	return calories, nil
}

// Вычисляем беговые калории
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем на отрицательные значения шаги
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}
	// Проверяем на отрицательные значения веса, роста и длительности
	if weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("data must be positive")
	}
	// Расчитываем и возвращем калории при беге
	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories, nil
}

// Вычисляем среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка отрицательных шагов
	if duration <= 0 {
		return 0
	}
	// Вычисляем среднюю скорость
	speed := Distance(steps, height) / duration.Hours()
	return speed
}

// Вычисляем дистанцию
func Distance(steps int, height float64) float64 {
	distance := (float64(steps) * height * stepLengthCoefficient) / mInKm
	return distance
}
