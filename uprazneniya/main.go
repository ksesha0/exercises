package main

import (
    "fmt"
    "log"
    "uprazneniya/db"
    "uprazneniya/model"
    "uprazneniya/service"
)

func main() {
    connStr := "user=postgres dbname=postgres sslmode=disable password=Wordux80@"
    database, err := db.Connect(connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    muscleGroups := map[string]int{
        "Ноги":    3,
        "Спина":   2,
        "Грудь":   2,
        "Плечи":   1,
        "Трицепс": 1,
        "Бицепс":  1,
        "Пресс":   3,
    }

    var trainingPlan []model.Exercise

    for muscle, count := range muscleGroups {
        exercises, err := db.GetExercises(database, muscle)
        if err != nil {
            log.Fatalf("Error getting exercises for muscle %s: %v", muscle, err)
        }

        selected := service.PickRandomExercises(exercises, count)
        trainingPlan = append(trainingPlan, selected...)
    }

    fmt.Println("Ваш тренировочный план:")
    for _, ex := range trainingPlan {
        fmt.Printf("%s (%s - %s)\n", ex.Exercise, ex.Muscles, ex.MuscleSubgroup)
    }
}
