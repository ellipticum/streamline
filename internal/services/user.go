package services

import (
	"github.com/ellipticum/streamline/internal/dto"
)

type User struct{}

func (user User) Create(data dto.UserCreateDTO) dto.UserCreateDTO {
	return data
}
