package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// Делим строку запятой
	parts := strings.Split(datastring, ",")

	// Проверяем что длинну слайса
	if len(parts) != 3 {
		return errors.New("invalid string format")
	}

	// Преобразуем количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid format steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("the number of steps must be positive")
	}
	//Сохраняем значения в структуру
	t.Steps = steps
	// Сохраняем тип тренировки
	t.TrainingType = parts[1]

	// Преобразуем продолжительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid format duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("the duration of steps must be positive")
	}
	//Сохраняем значения в структуру
	t.Duration = duration

	// Возвращаем отсутсвие ошибки
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Вычислям дистанцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	// Вычисляем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var calcErr error

	switch t.TrainingType {
	case "Ходьба":
		// Расчет калорий ходьбы
		calories, calcErr = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if calcErr != nil {
			return "", calcErr
		}

	case "Бег":
		// Расчет калорий для бега
		calories, calcErr = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if calcErr != nil {
			return "", calcErr
		}

	default:
		//неизвестный тип тренировки
		return "", errors.New("неизвестный тип тренировки")
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return result, nil

}
