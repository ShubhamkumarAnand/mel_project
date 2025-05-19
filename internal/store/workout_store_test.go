package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("opening test db: %v", err)
	}

	// run the migrations for our test
	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("migrating test db error: %v", err)
	}

	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("truncating tables : %v", err)
	}
	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)
	tests := []struct {
		name      string
		workout   *Workout
		wantError bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "push day",
				Description:     "upper body day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Bench Press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(135.5),
						Notes:        "Warm up properly",
						OrderIndex:   1,
					},
				},
			},
			wantError: false,
		},
		{
			name: "workout with invalid enteris",
			workout: &Workout{
				Title:           "full body",
				Description:     "complete workout",
				DurationMinutes: 90,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Sets:         3,
						Reps:         IntPtr(3),
						Notes:        "Keep Form",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "squats",
						Sets:            4,
						Reps:            IntPtr(12),
						DurationSeconds: IntPtr(60),
						Weight:          FloatPtr(165.88),
						OrderIndex:      2,
					},
				},
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createWorkout, err := store.CreateWorkout((tt.workout))
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createWorkout.Title)
			assert.Equal(t, tt.workout.Description, createWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, createWorkout.DurationMinutes)
			assert.Equal(t, tt.workout.CaloriesBurned, createWorkout.CaloriesBurned)

			retrieved, err := store.GetWorkoutByID(int64(createWorkout.ID))
			require.NoError(t, err)

			assert.Equal(t, createWorkout.ID, retrieved.ID)
			assert.Equal(t, len(tt.workout.Entries), len(retrieved.Entries))

			for i := range retrieved.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, tt.workout.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Sets, tt.workout.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].Reps, tt.workout.Entries[i].Reps)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, tt.workout.Entries[i].OrderIndex)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func FloatPtr(i float64) *float64 {
	return &i
}
