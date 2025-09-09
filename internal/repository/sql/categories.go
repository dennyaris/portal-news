package sql

import (
	"context"

	"github.com/dennyaris/portal-news/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepo struct{ db *gorm.DB }

func NewCategoryRepo(db *gorm.DB) *CategoryRepo { return &CategoryRepo{db: db} }

func (r *CategoryRepo) Create(ctx context.Context, c *entity.Category) error {
	return r.db.WithContext(ctx).Create(c).Error
}
func (r *CategoryRepo) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	var c entity.Category
	if err := r.db.WithContext(ctx).First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *CategoryRepo) GetBySlug(ctx context.Context, slug string) (*entity.Category, error) {
	var c entity.Category
	if err := r.db.WithContext(ctx).First(&c, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *CategoryRepo) List(ctx context.Context, limit, page int) ([]*entity.Category, int, error) {
	var cats []*entity.Category
	var total int64
	db := r.db.WithContext(ctx).Model(&entity.Category{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	if err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&cats).Error; err != nil {
		return nil, 0, err
	}
	return cats, int(total), nil
}
func (r *CategoryRepo) Update(ctx context.Context, c *entity.Category) error {
	return r.db.WithContext(ctx).Save(c).Error
}
func (r *CategoryRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Category{}, "id = ?", id).Error
}
