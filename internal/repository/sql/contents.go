package sql

import (
	"context"

	"github.com/dennyaris/portal-news/internal/entity"
	"gorm.io/gorm"
)

type ContentRepo struct{ db *gorm.DB }

func NewContentRepo(db *gorm.DB) *ContentRepo { return &ContentRepo{db: db} }

func (r *ContentRepo) Create(ctx context.Context, c *entity.Content) error {
	return r.db.WithContext(ctx).Create(c).Error
}
func (r *ContentRepo) GetByID(ctx context.Context, id string) (*entity.Content, error) {
	var c entity.Content
	if err := r.db.WithContext(ctx).First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *ContentRepo) List(ctx context.Context, limit, page int, filters map[string]string) ([]*entity.Content, int, error) {
	var list []*entity.Content
	var total int64
	db := r.db.WithContext(ctx).Model(&entity.Content{})
	if v := filters["status"]; v != "" {
		db = db.Where("status = ?", v)
	}
	if v := filters["author"]; v != "" {
		db = db.Where("author_id = ?", v)
	}
	if v := filters["cat"]; v != "" {
		db = db.Where("category_id = ?", v)
	}
	if v := filters["q"]; v != "" {
		like := "%" + v + "%"
		db = db.Where("title LIKE ? OR body LIKE ?", like, like)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	if err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, int(total), nil
}
func (r *ContentRepo) Update(ctx context.Context, c *entity.Content) error {
	return r.db.WithContext(ctx).Save(c).Error
}
func (r *ContentRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Content{}, "id = ?", id).Error
}
