package dto

import "mymodule/model"

type RegisterDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) RegisterDto {
	return RegisterDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
