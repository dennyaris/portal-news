package repository

import (
	"context"

	"github.com/dennyaris/portal-news/internal/entity"
)

type UserRepository interface {
	Create(context.Context, *entity.User) error
	GetByID(context.Context, string) (*entity.User, error)
	List(context.Context, int, int) ([]*entity.User, int, error)
	Update(context.Context, *entity.User) error
	Delete(context.Context, string) error
}

type CategoryRepository interface {
	Create(context.Context, *entity.Category) error
	GetByID(context.Context, string) (*entity.Category, error)
	GetBySlug(context.Context, string) (*entity.Category, error)
	List(context.Context, int, int) ([]*entity.Category, int, error)
	Update(context.Context, *entity.Category) error
	Delete(context.Context, string) error
}

type ContentRepository interface {
	Create(context.Context, *entity.Content) error
	GetByID(context.Context, string) (*entity.Content, error)
	List(context.Context, int, int, map[string]string) ([]*entity.Content, int, error)
	Update(context.Context, *entity.Content) error
	Delete(context.Context, string) error
}
