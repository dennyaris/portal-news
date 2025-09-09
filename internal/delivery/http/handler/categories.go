package handler

import (
	"context"
	"net/http"

	"github.com/dennyaris/portal-news/internal/usecase/category"
	"github.com/dennyaris/portal-news/pkg/pagination"
	"github.com/dennyaris/portal-news/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct{ svc category.Service }

func NewCategoryHandler(s category.Service) *CategoryHandler { return &CategoryHandler{svc: s} }

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var in category.Input
	if err := c.BodyParser(&in); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	cat, err := h.svc.Create(context.Background(), in)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusCreated, cat)
}
func (h *CategoryHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	cat, err := h.svc.Get(context.Background(), id)
	if err != nil {
		return response.ErrorJSON(c, http.StatusNotFound, err.Error())
	}
	return response.JSON(c, http.StatusOK, cat)
}
func (h *CategoryHandler) List(c *fiber.Ctx) error {
	p := pagination.Parse(c, 10, 1)
	items, total, err := h.svc.List(context.Background(), p.Limit, p.Page)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusOK, pagination.Result[any]{Data: sliceAny(items), Page: p.Page, Limit: p.Limit, Total: total})
}
func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var in category.UpdateInput
	if err := c.BodyParser(&in); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	cat, err := h.svc.Update(context.Background(), id, in)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusOK, cat)
}
func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Delete(context.Background(), id); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return c.SendStatus(http.StatusNoContent)
}
