package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем входную строку на слайс строк
	// Проверяем, что слайс содержит ровно два элемента
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid string format: expected 2 elements")
	}
	// Преобразование слайса (количество шагов)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid format steps: %w", err)
	}
	// Проверка кол-ва шагов на 0
	if steps <= 0 {
		return errors.New("the number of steps must be positive")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid format duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("the duration must be positive")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем дистанцию

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Вычисляем количество потраченных калорий

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return result, nil
}
