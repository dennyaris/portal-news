package content

import (
	"context"
	"time"

	"github.com/dennyaris/portal-news/internal/entity"
	"github.com/dennyaris/portal-news/internal/repository"
)

type Input struct {
	Title      string               `json:"title" validate:"required,min=3"`
	Slug       string               `json:"slug" validate:"required,min=2"`
	Body       string               `json:"body" validate:"required,min=10"`
	Status     entity.ContentStatus `json:"status" validate:"required,oneof=draft published"`
	AuthorID   string               `json:"author_id" validate:"required"`
	CategoryID string               `json:"category_id" validate:"required"`
}

type UpdateInput struct {
	Title      string                `json:"title" validate:"omitempty,min=3"`
	Slug       string                `json:"slug" validate:"omitempty,min=2"`
	Body       string                `json:"body" validate:"omitempty,min=10"`
	Status     *entity.ContentStatus `json:"status"`
	AuthorID   string                `json:"author_id" validate:"omitempty"`
	CategoryID string                `json:"category_id" validate:"omitempty"`
}

type Service interface {
	Create(ctx context.Context, in Input) (*entity.Content, error)
	Get(ctx context.Context, id string) (*entity.Content, error)
	List(ctx context.Context, limit, page int, filters map[string]string) ([]*entity.Content, int, error)
	Update(ctx context.Context, id string, in UpdateInput) (*entity.Content, error)
	Delete(ctx context.Context, id string) error
}

type svc struct {
	repo repository.ContentRepository
	id   func() string
	now  func() time.Time
	val  func(any) error
}

func New(repo repository.ContentRepository, id func() string, now func() time.Time, val func(any) error) Service {
	return &svc{repo: repo, id: id, now: now, val: val}
}

func (s *svc) Create(ctx context.Context, in Input) (*entity.Content, error) {
	if err := s.val(in); err != nil {
		return nil, err
	}
	var pub *time.Time
	if in.Status == entity.StatusPublished {
		t := s.now()
		pub = &t
	}
	c := &entity.Content{ID: s.id(), Title: in.Title, Slug: in.Slug, Body: in.Body, Status: in.Status, AuthorID: in.AuthorID, CategoryID: in.CategoryID, PublishedAt: pub, CreatedAt: s.now(), UpdatedAt: s.now()}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
func (s *svc) Get(ctx context.Context, id string) (*entity.Content, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *svc) List(ctx context.Context, limit, page int, filters map[string]string) ([]*entity.Content, int, error) {
	return s.repo.List(ctx, limit, page, filters)
}
func (s *svc) Update(ctx context.Context, id string, in UpdateInput) (*entity.Content, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if in.Title != "" {
		c.Title = in.Title
	}
	if in.Slug != "" {
		c.Slug = in.Slug
	}
	if in.Body != "" {
		c.Body = in.Body
	}
	if in.AuthorID != "" {
		c.AuthorID = in.AuthorID
	}
	if in.CategoryID != "" {
		c.CategoryID = in.CategoryID
	}
	if in.Status != nil {
		c.Status = *in.Status
		if c.Status == entity.StatusPublished && c.PublishedAt == nil {
			t := s.now()
			c.PublishedAt = &t
		}
		if c.Status == entity.StatusDraft {
			c.PublishedAt = nil
		}
	}
	c.UpdatedAt = s.now()
	if err := s.repo.Update(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
func (s *svc) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
