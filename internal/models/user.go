package models

// CreateUserRequest defines the incoming payload for user creation.
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// UpdateUserRequest defines the incoming payload for user updates.
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// UserResponse represents the outward-facing user structure.
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}
