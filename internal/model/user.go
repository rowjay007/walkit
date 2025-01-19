package model

type User struct {
    ID              string    `json:"id,omitempty"`
    Username        string    `json:"username"`
    Email           string    `json:"email"`
    Password        string    `json:"password,omitempty"`
    PasswordConfirm string    `json:"passwordConfirm,omitempty"`
    EmailVisibility bool      `json:"emailVisibility"`
    FitnessGoal     string    `json:"fitnessGoal,omitempty"`
    ActivityLevel   string    `json:"activityLevel,omitempty"`
    Avatar          string    `json:"avatar,omitempty"`
    CreatedAt       string    `json:"created"`
    UpdatedAt       string    `json:"updated"`
}

type UserFilter struct {
    Username      string `form:"username"`
    Email         string `form:"email"`
    FitnessGoal   string `form:"fitnessGoal"`
    ActivityLevel string `form:"activityLevel"`
    SortBy        string `form:"sortBy"`
    SortOrder     string `form:"sortOrder"` // asc or desc
    Page          int    `form:"page"`
    PerPage       int    `form:"perPage"`
}

type UserListResponse struct {
    Items      []User `json:"items"`
    TotalItems int    `json:"totalItems"`
    Page       int    `json:"page"`
    PerPage    int    `json:"perPage"`
    TotalPages int    `json:"totalPages"`
}

type LoginRequest struct {
    Identity string `json:"identity" binding:"required,min=3,max=100"`
    Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}

type PasswordResetRequest struct {
    Email string `json:"email" binding:"required,email"`
}

type ConfirmPasswordResetRequest struct {
    Token           string `json:"token" binding:"required"`
    Password        string `json:"password" binding:"required,min=8"`
    PasswordConfirm string `json:"passwordConfirm" binding:"required,eqfield=Password"`
}

type UpdateUserRequest struct {
    Username      string `json:"username,omitempty"`
    Email         string `json:"email,omitempty"`
    FitnessGoal   string `json:"fitnessGoal,omitempty"`
    ActivityLevel string `json:"activityLevel,omitempty"`
    Avatar        string `json:"avatar,omitempty"`
}