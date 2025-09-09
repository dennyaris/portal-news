package handler

import (
	"context"
	"net/http"

	"github.com/dennyaris/portal-news/internal/usecase/content"
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/news-portal-gorm/pkg/pagination"
	"github.com/yourname/news-portal-gorm/pkg/response"
)

type ContentHandler struct{ svc content.Service }

func NewContentHandler(s content.Service) *ContentHandler { return &ContentHandler{svc: s} }

func (h *ContentHandler) Create(c *fiber.Ctx) error {
	var in content.Input
	if err := c.BodyParser(&in); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	v, err := h.svc.Create(context.Background(), in)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusCreated, v)
}
func (h *ContentHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	v, err := h.svc.Get(context.Background(), id)
	if err != nil {
		return response.ErrorJSON(c, http.StatusNotFound, err.Error())
	}
	return response.JSON(c, http.StatusOK, v)
}
func (h *ContentHandler) List(c *fiber.Ctx) error {
	p := pagination.Parse(c, 10, 1)
	filters := map[string]string{
		"status": c.Query("status"),
		"cat":    c.Query("cat"),
		"author": c.Query("author"),
		"q":      c.Query("q"),
	}
	items, total, err := h.svc.List(context.Background(), p.Limit, p.Page, filters)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusOK, pagination.Result[any]{Data: sliceAny(items), Page: p.Page, Limit: p.Limit, Total: total})
}
func (h *ContentHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var in content.UpdateInput
	if err := c.BodyParser(&in); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	v, err := h.svc.Update(context.Background(), id, in)
	if err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return response.JSON(c, http.StatusOK, v)
}
func (h *ContentHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Delete(context.Background(), id); err != nil {
		return response.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	return c.SendStatus(http.StatusNoContent)
}
