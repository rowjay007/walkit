package config

const (
	PocketBaseURL   = "http://127.0.0.1:8090"
	UsersAPI        = PocketBaseURL + "/api/collections/users/records"
	AuthAPI         = PocketBaseURL + "/api/collections/users/auth-with-password"
	ResetAPI        = PocketBaseURL + "/api/collections/users/request-password-reset"
	ConfirmResetAPI = PocketBaseURL + "/api/collections/users/confirm-password-reset"
	JWTSecret       = "your_jwt_secret_key" 
)
