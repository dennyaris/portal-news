package category

import (
	"context"
	"time"

	"github.com/dennyaris/portal-news/internal/entity"
	"github.com/dennyaris/portal-news/internal/repository"
)

type Input struct {
	Name string `json:"name" validate:"required,min=2"`
	Slug string `json:"slug" validate:"required,min=2"`
}

type UpdateInput struct {
	Name string `json:"name" validate:"required,min=2"`
	Slug string `json:"slug" validate:"required,min=2"`
}

type Service interface {
	Create(ctx context.Context, in Input) (*entity.Category, error)
	Get(ctx context.Context, id string) (*entity.Category, error)
	List(ctx context.Context, limit, page int) ([]*entity.Category, int, error)
	Update(ctx context.Context, id string, in UpdateInput) (*entity.Category, error)
	Delete(ctx context.Context, id string) error
}

type svc struct {
	repo repository.CategoryRepository
	id   func() string
	now  func() time.Time
	val  func(any) error
}

func New(repo repository.CategoryRepository, id func() string, now func() time.Time, val func(any) error) Service {
	return &svc{repo: repo, id: id, now: now, val: val}
}

func (s *svc) Create(ctx context.Context, in Input) (*entity.Category, error) {
	if err := s.val(in); err != nil {
		return nil, err
	}
	c := &entity.Category{ID: s.id(), Name: in.Name, Slug: in.Slug, CreatedAt: s.now(), UpdatedAt: s.now()}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
func (s *svc) Get(ctx context.Context, id string) (*entity.Category, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *svc) List(ctx context.Context, limit, page int) ([]*entity.Category, int, error) {
	return s.repo.List(ctx, limit, page)
}
func (s *svc) Update(ctx context.Context, id string, in UpdateInput) (*entity.Category, error) {
	if err := s.val(in); err != nil {
		return nil, err
	}
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	c.Name, c.Slug, c.UpdatedAt = in.Name, in.Slug, s.now()
	if err := s.repo.Update(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
func (s *svc) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
