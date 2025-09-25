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
	Avatar    string `json:"avatar_url,omitempty"`
	CreatedAt string `json:"created_at"`
}

type UpdateUserRequest struct {
	Username *string `json:"username" validate:"required,min=3"`
	Name     *string `json:"name"`
	Email    *string `json:"email" validate:"required,email"`
	Bio      *string `json:"bio"`
	Password *string `json:"password" validate:"required,min=8"`
	Avatar   *string `json:"avatar_url"`
}
