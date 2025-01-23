package model

// Response represents a standard API response
type Response struct {
	Message string      `json:"message" example:"Success message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"Error message"`
}