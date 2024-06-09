package service

import (
	"math/rand"
	"time"
	"uprazneniya/model"
)

func PickRandomExercises(exercises []model.Exercise, count int) []model.Exercise {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(exercises), func(i, j int) { exercises[i], exercises[j] = exercises[j], exercises[i] })

	selected := make(map[string]bool)
	var result []model.Exercise

	for _, ex := range exercises {
		if !selected[ex.MuscleSubgroup] {
			result = append(result, ex)
			selected[ex.MuscleSubgroup] = true
			if len(result) == count {
				break
			}
		}
	}
	return result
}
