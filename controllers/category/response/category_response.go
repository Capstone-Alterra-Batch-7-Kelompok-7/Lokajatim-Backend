package response

import (
	"lokajatim/entities"
	"time"
)

// CategoryResponse is the response for the Category controller
// @Description CategoryResponse is the response for the Category controller
// @Param ID int true "ID of the category"
// @Param Name string true "Name of the category"
// @Param CreatedAt string true "Created At of the category"
// @Param UpdatedAt string true "Updated At of the category"
type CategoryResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	}

func CategoryFromEntities(category entities.Category) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}