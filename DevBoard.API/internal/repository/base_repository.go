// repository/base_repository.go
package repository

import (
	"errors"

	"gorm.io/gorm"
)

// BaseRepository generic CRUD operasyonları için
type BaseRepository[T any, ID any] interface {
	Create(entity *T) error
	GetByID(id ID) (*T, error)
	List(limit, offset int) ([]T, error)
	ListAll() ([]T, error)
	Update(entity *T) error
	Delete(id ID) error
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