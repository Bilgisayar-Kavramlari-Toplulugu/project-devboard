package repository

import (
	"errors"
	"project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSkillRepository interface {
	BaseRepository[entities.UserSkill, int]
	ValidateIds(skillIDs []int) error
	GetByUserID(userID uuid.UUID) ([]entities.UserSkill, error)
	DeleteByUserID(userID uuid.UUID) error
	CreateBatch(userSkills []entities.UserSkill) error
}

type userSkillRepository struct {
	db *gorm.DB
	BaseRepository[entities.UserSkill, int]
}

func NewUserSkillRepository(db *gorm.DB) UserSkillRepository {
	return &userSkillRepository{
		db:             db,
		BaseRepository: NewBaseRepository[entities.UserSkill, int](db),
	}
}

func (r *userSkillRepository) ValidateIds(skillIDs []int) error {
	if len(skillIDs) == 0 {
		return nil
	}

	var count int64
	if err := r.db.Model(&entities.Skill{}).Where("id IN ?", skillIDs).Count(&count).Error; err != nil {
		return err
	}

	if count != int64(len(skillIDs)) {
		return errors.New("some skill IDs do not exist")
	}

	return nil
}

func (r *userSkillRepository) GetByUserID(userID uuid.UUID) ([]entities.UserSkill, error) {
	var userSkills []entities.UserSkill
	err := r.db.Preload("Skill").Where("user_id = ?", userID).Order("created_on desc").Find(&userSkills).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.UserSkill{}, nil
		}
		return nil, err
	}
	return userSkills, nil
}

func (r *userSkillRepository) DeleteByUserID(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&entities.UserSkill{}).Error
}

// CreateBatch normalde create fonksiyonunu override ederdi ancak BaseRepository'de Create fonksiyonu tek bir entity alacak şekilde tanımlandığı için yeni bir fonksiyon olarak ekledim.
func (r *userSkillRepository) CreateBatch(userSkills []entities.UserSkill) error {
	if len(userSkills) == 0 {
		return nil
	}

	err := r.db.Create(userSkills).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil // Eğer duplicate key hatası gelirse bunu görmezden gel. Çünkü bu fonksiyonda zaten var olan kayıtları eklemeye çalışmayacağız.
		}
		return err
	}

	return nil
}
