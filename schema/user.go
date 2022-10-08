package schema

type UserSchema struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
	UserPassword string `json:"user_password,omitempty"`
}
