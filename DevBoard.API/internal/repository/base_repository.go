package repository

import (
	"errors"
	"math"
	"time"

	"gorm.io/gorm"
	"project-devboard/pkg/pagination"
)

type BaseRepository[T any, ID any] interface {
	Create(entity *T) error
	GetByID(id ID) (*T, error)
	List(limit int) ([]T, error)
	Update(entity *T) error
	Delete(id ID) error
	PaginatedList(page, pageSize int) (*pagination.PaginatedResult[T], error)
}



type baseRepository[T any, ID any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any, ID any](db *gorm.DB) BaseRepository[T, ID] {
	return &baseRepository[T, ID]{db: db}
}

func (r *baseRepository[T, ID]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepository[T, ID]) GetByID(id ID) (*T, error) {
	var entity T
	err := r.db.First(&entity, "id = ?", id).Error // ← güvenli hali
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T, ID]) List(limit int) ([]T, error) {
	var entities []T

	if limit < 1 {
		limit = 10
	}
	if limit > 1000 {
        return nil, errors.New(
            "limit cannot exceed 1000, please use ListPaginated(page, pageSize) for larger datasets",
        )
    }

	err := r.db.Order("created_at desc").
		Limit(limit).
		Find(&entities).Error
	return entities, err
}

func (r *baseRepository[T, ID]) Update(entity *T) error {
	return r.db.Updates(entity).Error //sadece değişenleri update etmek için 
}

func (r *baseRepository[T, ID]) Delete(id ID) error {
    var entity T
    return r.db.Model(&entity).
        Where("id = ?", id).
        Updates(map[string]any{
            "is_deleted": true,
            "deleted_on": time.Now(),
        }).Error
}

func (r *baseRepository[T, ID]) PaginatedList(page, pageSize int) (*pagination.PaginatedResult[T], error) {
	var entities []T
	var total int64

	if err := r.db.Model(new(T)).Count(&total).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 1000 {
        pageSize = 1000
    }

	offset := (page - 1) * pageSize

	err := r.db.Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &pagination.PaginatedResult[T]{
		Data:       entities,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}