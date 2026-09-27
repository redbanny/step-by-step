package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	var splitData = strings.Split(data, ",")
	if len(splitData) != 2 {
		return 0, 0, errors.New("Incorrect input string")
	}

	stepsCount, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, 0, errors.New("Incorrect steps value")
	}

	trainingDuration, err := time.ParseDuration(splitData[1])
	if err != nil {
		return 0, 0, errors.New("Incorrect durations value")
	}

	return stepsCount, trainingDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	var result string
	stepsCount, traningDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return result
	}

	if stepsCount < 0 {
		return result
	}

	distance := (float32(stepsCount) * stepLength) / mInKm

	walkingCaloriesCount, err := spentcalories.WalkingSpentCalories(stepsCount, weight, height, traningDuration)
	if err != nil {
		return result
	}
	result = fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", stepsCount, distance, walkingCaloriesCount)

	return result
}
