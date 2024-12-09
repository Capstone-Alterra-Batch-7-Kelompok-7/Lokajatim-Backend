package response

import "lokajatim/entities"

type UpdateUserResponse struct {
	ID          int    `json:"id_user"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	NIK         string `json:"nik,omitempty"`
}

func UpdateFromEntities(user entities.User) UpdateUserResponse {
	return UpdateUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		NIK:         user.NIK,
	}
}
