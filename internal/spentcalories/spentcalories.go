package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	splitData := strings.Split(data, ",")
	if len(splitData) != 3 {
		return 0, "", 0, errors.New("Incorrect input data")
	}

	stepsCount, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, "", 0, errors.New("Incorrect steps value")
	}

	trainingDuration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return 0, "", 0, errors.New("Incorrect durations value")
	}

	return stepsCount, splitData[1], trainingDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var result string
	var kkalCount float64
	var err error
	stepsCount, trainingType, trainingDuration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return result, err
	}

	switch strings.ToLower(trainingType) {
	case "ходьба":
		kkalCount, err = WalkingSpentCalories(stepsCount, weight, height, trainingDuration)
	case "бег":
		kkalCount, err = RunningSpentCalories(stepsCount, weight, height, trainingDuration)
	default:
		err = errors.New("unknown type of training")
	}
	if err != nil {
		return result, err
	}

	speed := meanSpeed(stepsCount, height, trainingDuration)
	distance := distance(stepsCount, height)
	result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2fч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		trainingType, trainingDuration.Hours(), distance, speed, kkalCount)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Incorrect input args")
	}

	avgSpeed := meanSpeed(steps, height, duration)

	return (weight * avgSpeed * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Incorrect input args")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	kkalCount := (weight * avgSpeed * duration.Minutes()) / minInH

	return kkalCount * walkingCaloriesCoefficient, nil
}
