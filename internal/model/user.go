package model

type User struct {
	ID              string `json:"id,omitempty"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"passwordConfirm,omitempty"`
	EmailVisibility bool   `json:"emailVisibility"`
	FitnessGoal     string `json:"fitnessGoal,omitempty"` // e.g., "weight loss", "muscle gain"
	ActivityLevel   string `json:"activityLevel,omitempty"` // e.g., "sedentary", "active", "very active"
}

type LoginRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type PasswordResetRequest struct {
	Email string `json:"email"`
}

type ConfirmPasswordResetRequest struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UpdateUserRequest struct {
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	FitnessGoal   string `json:"fitnessGoal,omitempty"`
	ActivityLevel string `json:"activityLevel,omitempty"`
}
