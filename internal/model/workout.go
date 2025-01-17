package model

type WorkoutPlan struct {
    ID          string         `json:"id,omitempty"`
    UserID      string         `json:"userId"`
    Name        string         `json:"name"`
    Description string         `json:"description"`
    Exercises   []ExerciseSet  `json:"exercises"`
    Schedule    WorkoutSchedule `json:"schedule"`
    CreatedAt   string         `json:"created"`
    UpdatedAt   string         `json:"updated"`
}


type ExerciseSet struct {
    ExerciseID string  `json:"exerciseId"`
    Sets       int     `json:"sets"`
    Reps       int     `json:"reps"`
    Weight     float64 `json:"weight"`
    Duration   int     `json:"duration"` // in seconds, for cardio exercises
    Notes      string  `json:"notes"`
}

type WorkoutSchedule struct {
    StartDate string   `json:"startDate"`
    EndDate   string   `json:"endDate"`
    Days      []string `json:"days"` // e.g., ["Monday", "Wednesday", "Friday"]
    Time      string   `json:"time"` // e.g., "18:00"
}