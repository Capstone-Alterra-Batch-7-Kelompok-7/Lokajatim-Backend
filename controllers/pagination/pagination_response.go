package pagination

import (
	"math"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Meta struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalCount int64       `json:"total_count"`
	TotalPages int         `json:"total_pages"`
}

type PaginatedResponse struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       Meta        `json:"meta"`
}

func SuccessPaginatedResponse(c echo.Context, data interface{}, page, limit int, totalCount int64) error {
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	
	return c.JSON(http.StatusOK, PaginatedResponse{
		Status:     true,
		Message:    "success",
		Data:       data,
		Meta: Meta{
			Page:       page,
			Limit:      limit,
			TotalCount: totalCount,
			TotalPages: totalPages,
		},
	})
}
