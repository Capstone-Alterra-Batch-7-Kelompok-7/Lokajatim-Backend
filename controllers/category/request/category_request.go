package request

import "lokajatim/entities"

// CategoryRequest is the request for the Category endpoint
// @Description CategoryRequest is the request for the Category endpoint
// @Param Name string true "Name of the category"
type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (categoryRequest CategoryRequest) ToEntities() (entities.Category, error) {
	return entities.Category{
		Name: categoryRequest.Name,
	}, nil
}
