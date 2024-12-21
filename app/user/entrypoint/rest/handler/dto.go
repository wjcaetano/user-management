package handler

import entity "user-management/app/user"

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (req *CreateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToResponse(user *entity.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
