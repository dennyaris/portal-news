package sql

import (
	"context"

	"github.com/dennyaris/portal-news/internal/entity"
	"gorm.io/gorm"
)

type UserRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) Create(ctx context.Context, u *entity.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var u entity.User
	if err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) List(ctx context.Context, limit, page int) ([]*entity.User, int, error) {
	var users []*entity.User
	var total int64
	db := r.db.WithContext(ctx).Model(&entity.User{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	if err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, int(total), nil
}

func (r *UserRepo) Update(ctx context.Context, u *entity.User) error {
	return r.db.WithContext(ctx).Save(u).Error
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.User{}, "id = ?", id).Error
}
