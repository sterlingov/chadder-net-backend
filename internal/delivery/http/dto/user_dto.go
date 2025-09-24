package httpdto

type CreateUserRequest struct {
	Username string  `json:"username" validate:"required,min=3"`
	Name     *string `json:"name,omitempty"`
	Bio      *string `json:"bio,omitempty"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"created_at"`
}

type UpdateUserRequest struct {
	Username *string
	Name     *string
	Email    *string
	Bio      *string
	Password *string
	Avatar   *string
}
