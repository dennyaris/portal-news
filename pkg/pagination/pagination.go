package pagination

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Params struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Result[T any] struct {
	Data  []T `json:"data"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func Parse(c *fiber.Ctx, defLimit, defPage int) Params {
	limit, _ := strconv.Atoi(c.Query("limit", strconv.Itoa(defLimit)))
	page, _ := strconv.Atoi(c.Query("page", strconv.Itoa(defPage)))
	if limit <= 0 {
		limit = defLimit
	}
	if page <= 0 {
		page = defPage
	}
	return Params{Limit: limit, Page: page}
}
