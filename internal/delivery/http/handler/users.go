package handler

import (
	"context"
	"net/http"

	"github.com/dennyaris/portal-news/internal/usecase/user"
	"github.com/dennyaris/portal-news/pkg/pagination"
	"github.com/dennyaris/portal-news/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{ svc user.Service }

func NewUserHandler(s user.Service) *UserHandler { return &UserHandler{svc: s} }

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var in user.Input
	if err := c.BodyParser(&in); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	u, err := h.svc.Create(context.Background(), in)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusCreated, u)
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	u, err := h.svc.Get(context.Background(), id)
	if err != nil {
		return response.ErrorJSON(c, http.StatusNotFound, err.Error())
	}
	return response.JSON(c, http.StatusOK, u)
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	p := pagination.Parse(c, 10, 1)
	items, total, err := h.svc.List(context.Background(), p.Limit, p.Page)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusOK, pagination.Result[any]{Data: sliceAny(items), Page: p.Page, Limit: p.Limit, Total: total})
}

func sliceAny[T any](in []T) []any {
	out := make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}
