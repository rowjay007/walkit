package model

type Progress struct {
    ID          string    `json:"id,omitempty"`
    UserID      string    `json:"userId"`
    Date        string    `json:"date"`
    Weight      float64   `json:"weight"`      // in kg
    BodyFat     float64   `json:"bodyFat"`     // percentage
    Measurements BodyMeasurements `json:"measurements"`
    Photos      []string  `json:"photos"`      // photo URLs
    Notes       string    `json:"notes"`
    CreatedAt   string    `json:"created"`
}

type BodyMeasurements struct {
    Chest      float64 `json:"chest"`      // in cm
    Waist      float64 `json:"waist"`
    Hips       float64 `json:"hips"`
    Arms       float64 `json:"arms"`
    Thighs     float64 `json:"thighs"`
    Calves     float64 `json:"calves"`
}