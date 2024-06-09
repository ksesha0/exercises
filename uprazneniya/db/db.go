package db

import (
	"database/sql"
	"uprazneniya/model"

	_ "github.com/lib/pq"
)

func Connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetExercises(db *sql.DB, muscle string) ([]model.Exercise, error) {
	rows, err := db.Query("SELECT id, exercise, muscles, COALESCE(muscle_subgroup, '') FROM exercises WHERE muscles=$1", muscle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []model.Exercise
	for rows.Next() {
		var ex model.Exercise
		if err := rows.Scan(&ex.ID, &ex.Exercise, &ex.Muscles, &ex.MuscleSubgroup); err != nil {
			return nil, err
		}
		exercises = append(exercises, ex)
	}
	return exercises, nil
}
