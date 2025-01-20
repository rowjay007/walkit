package model

// Exercise struct to hold the details of the exercise.
type Exercise struct {
    ID          string   `json:"id,omitempty"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Category    string   `json:"category"`
    MuscleGroup []string `json:"muscleGroup,omitempty"`
    Duration    int      `json:"duration"`  // Duration in seconds
    Intensity   string   `json:"intensity"` // Intensity (e.g., Low, Medium, High)
    CreatedAt   string   `json:"created"`
    UpdatedAt   string   `json:"updated"`
}

// ExerciseFilter struct to handle filter parameters for listing exercises.
type ExerciseFilter struct {
    Name        string `form:"name"`
    Category    string `form:"category"`
    MuscleGroup string `form:"muscleGroup"`
    SortBy      string `form:"sortBy"`
    SortOrder   string `form:"sortOrder"`
    Page        int    `form:"page"`
    PerPage     int    `form:"perPage"`
}

// ExerciseListResponse struct to hold the response data for listing exercises.
type ExerciseListResponse struct {
    Items      []Exercise `json:"items"`
    TotalItems int        `json:"totalItems"`
    Page       int        `json:"page"`
    PerPage    int        `json:"perPage"`
    TotalPages int        `json:"totalPages"`
}
