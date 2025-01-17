package model

type WorkoutLog struct {
    ID          string           `json:"id,omitempty"`
    UserID      string           `json:"userId"`
    WorkoutID   string           `json:"workoutId"`
    Date        string           `json:"date"`
    Duration    int              `json:"duration"` // in minutes
    Exercises   []ExerciseLog    `json:"exercises"`
    Comments    string           `json:"comments"`
    Rating      int              `json:"rating"` // 1-5 rating of workout difficulty
    CreatedAt   string           `json:"created"`
}

type ExerciseLog struct {
    ExerciseID string         `json:"exerciseId"`
    Sets       []ExerciseSet  `json:"sets"`
    Notes      string         `json:"notes"`
}