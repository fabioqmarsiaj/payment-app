package mappers

import (
	"api/dto"
	"api/models"
)

func UserToUserDTO(User *models.User) (*dto.UserDTO){
	userID := User.ID.Hex()
	userDTO := dto.UserDTO{ID: userID}
	return &userDTO
}