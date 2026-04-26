// repository/base_repository.go
package repository

import (
	"errors"
	"math"
	"project-devboard/pkg/pagination"

	"gorm.io/gorm"
)

// BaseRepository generic CRUD operasyonları için
type BaseRepository[T any, ID any] interface {
	Create(entity *T) error
	GetByID(id ID) (*T, error)
	ListAll() ([]T, error)
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
	err := r.db.First(&entity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil

}

func (r *baseRepository[T, ID]) List(limit, offset int) ([]T, error) {
	var entities []T
	err := r.db.Order("created_on desc").Limit(limit).Offset(offset).Find(&entities).Error
	return entities, err
}

func (r *baseRepository[T, ID]) ListAll() ([]T, error) {
	var entities []T
	err := r.db.Order("created_on desc").Find(&entities).Error
	return entities, err
}

func (r *baseRepository[T, ID]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T, ID]) Delete(id ID) error {
	var entity T
	return r.db.Delete(&entity, "id = ?", id).Error
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