package user

import (
	"context"
	"time"

	"github.com/dennyaris/portal-news/internal/entity"
	"github.com/dennyaris/portal-news/internal/repository"
)

type Input struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
}

type Service interface {
	Create(ctx context.Context, in Input) (*entity.User, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context, limit, page int) ([]*entity.User, int, error)
}

type svc struct {
	repo repository.UserRepository
	id   func() string
	now  func() time.Time
	val  func(any) error
}

func New(repo repository.UserRepository, id func() string, now func() time.Time, val func(any) error) Service {
	return &svc{repo: repo, id: id, now: now, val: val}
}
func (s *svc) Create(ctx context.Context, in Input) (*entity.User, error) {
	if err := s.val(in); err != nil {
		return nil, err
	}
	u := &entity.User{ID: s.id(), Name: in.Name, Email: in.Email, CreatedAt: s.now(), UpdatedAt: s.now()}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}
func (s *svc) Get(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *svc) List(ctx context.Context, limit, page int) ([]*entity.User, int, error) {
	return s.repo.List(ctx, limit, page)
}
