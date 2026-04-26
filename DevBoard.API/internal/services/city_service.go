package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type CityService interface {
	CreateCity(name string, code string, displayOrder *int, countryId *int, userId uuid.UUID) (int, error)
	GetCityById(id int) (*entities.City, error)
	GetAllCities() ([]entities.City, error)
	GetCitiesByCountryId(countryId int) ([]entities.City, error)
	UpdateCity(id int, name string, code string, displayOrder *int, countryId *int) error
	DeleteCity(id int) error
}

type cityService struct {
	repo repository.CityRepository
}

func NewCityService(repo repository.CityRepository) CityService {
	return &cityService{repo: repo}
}

func (s *cityService) CreateCity(name string, code string, displayOrder *int, countryId *int, userId uuid.UUID) (int, error) {
	var codePtr *string
	if code != "" {
		codePtr = &code
	}

	city := &entities.City{
		Name:         name,
		Code:         codePtr,
		DisplayOrder: displayOrder,
		CountryId:    countryId,
		BaseEntity: entities.BaseEntity{
			CreatedBy:      userId,
			LastModifiedBy: userId,
		},
	}
	err := s.repo.Create(city)
	if err != nil {
		return 0, err
	}
	return city.Id, nil
}

func (s *cityService) GetCityById(id int) (*entities.City, error) {
	city, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if city == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return city, nil
}

func (s *cityService) GetAllCities() ([]entities.City, error) {
	return s.repo.ListAll()
}

func (s *cityService) GetCitiesByCountryId(countryId int) ([]entities.City, error) {
	return s.repo.GetByCountryId(countryId)
}

func (s *cityService) UpdateCity(id int, name string, code string, displayOrder *int, countryId *int) error {
	city, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if city == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}

	city.Name = name
	if code != "" {
		city.Code = &code
	} else {
		city.Code = nil
	}
	city.DisplayOrder = displayOrder
	city.CountryId = countryId

	return s.repo.Update(city)
}

func (s *cityService) DeleteCity(id int) error {
	city, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if city == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return s.repo.Delete(city.Id)
}