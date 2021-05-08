package dto

import "Gin-HATTS/model"

type UserDto struct {
	Name      string `json:"name"`
	Email string `json:"email"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Email: user.Email,
	}
}
