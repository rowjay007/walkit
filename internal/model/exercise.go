package model

type Exercise struct {
    ID          string   `json:"id,omitempty"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Category    string   `json:"category"` 
    MuscleGroup []string `json:"muscleGroup"` 
    CreatedAt   string   `json:"created"`
    UpdatedAt   string   `json:"updated"`
}


