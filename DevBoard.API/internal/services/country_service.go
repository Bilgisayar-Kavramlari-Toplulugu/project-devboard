package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type CountryService interface {
	CreateCountry(name, flagCode, shortCode, phonePrefix string, userId uuid.UUID) (int, error)
	GetCountryById(id int) (*entities.Country, error)
	GetAllCountries() ([]entities.Country, error)
	GetAllCountriesAlphabetical() ([]entities.Country, error)
	UpdateCountry(id int, name, flagCode, shortCode, phonePrefix string) error
	DeleteCountry(id int) error
}

type countryService struct {
	repo repository.CountryRepository
}

func NewCountryService(repo repository.CountryRepository) CountryService {
	return &countryService{repo: repo}
}

func (s *countryService) CreateCountry(name, flagCode, shortCode, phonePrefix string, userId uuid.UUID) (int, error) {
	country := &entities.Country{
		Name:        name,
		FlagCode:    flagCode,
		ShortCode:   shortCode,
		PhonePrefix: phonePrefix,
		BaseEntity: entities.BaseEntity{
			CreatedBy:      userId,
			LastModifiedBy: userId,
		},
	}
	err := s.repo.Create(country)
	if err != nil {
		return 0, err
	}
	return country.Id, nil
}

func (s *countryService) GetCountryById(id int) (*entities.Country, error) {
	country, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if country == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return country, nil
}

func (s *countryService) GetAllCountries() ([]entities.Country, error) {
	return s.repo.List(1000, 0)
}

func (s *countryService) GetAllCountriesAlphabetical() ([]entities.Country, error) {
	return s.repo.ListAllAlphabetical()
}

func (s *countryService) UpdateCountry(id int, name, flagCode, shortCode, phonePrefix string) error {
	country, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if country == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	country.Name = name
	country.FlagCode = flagCode
	country.ShortCode = shortCode
	country.PhonePrefix = phonePrefix
	return s.repo.Update(country)
}

func (s *countryService) DeleteCountry(id int) error {
	country, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if country == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return s.repo.Delete(country.Id)
}