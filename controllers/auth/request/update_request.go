package request

import "lokajatim/entities"

type UpdateUserRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Address     string `json:"address,omitempty"`   
	PhoneNumber string `json:"phone_number,omitempty"` 
	NIK         string `json:"nik,omitempty"`        
}

func (req *UpdateUserRequest) ToEntities() entities.User {
	return entities.User{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		NIK:         req.NIK,
	}
}
