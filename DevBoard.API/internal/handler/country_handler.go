
package handler

import (
	"net/http"
	"project-devboard/internal/dtos"
	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CountryHandler struct {
	service   services.CountryService
	validator *validator.Validator
}

func NewCountryHandler(service services.CountryService, validator *validator.Validator) *CountryHandler {
	return &CountryHandler{service: service, validator: validator}
}

// @Summary Create a new country
// @Description Create a new country
// @Tags Countries
// @Accept json
// @Produce json
// @Param request body dtos.CountryCreateRequest true "Country details"
// @Success 201 {object} response.Response{data=dtos.CountryResponse}
// @Failure 400 {object} response.Response
// @Security BearerAuth
// @Router /countries [post]
func (h *CountryHandler) Create(c *gin.Context) {
	var req dtos.CountryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	userID := userIDFromContext(c, uuid.Nil)
	countryId, err := h.service.CreateCountry(req.Name, req.FlagCode, req.ShortCode, req.PhonePrefix, userID)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusCreated, dtos.CountryResponse{
		Id:          countryId,
		Name:        req.Name,
		FlagCode:    req.FlagCode,
		ShortCode:   req.ShortCode,
		PhonePrefix: req.PhonePrefix,
	})
}

// @Summary Get all countries
// @Description Get a list of all countries
// @Tags Countries
// @Produce json
// @Success 200 {object} response.Response{data=[]dtos.CountryResponse}
// @Router /countries [get]
func (h *CountryHandler) GetAll(c *gin.Context) {
	countries, err := h.service.GetAllCountries()
	if err != nil {
		c.Error(err)
		return
	}
	result := make([]dtos.CountryResponse, len(countries))
	for i, ct := range countries {
		result[i] = dtos.CountryResponse{
			Id:          ct.Id,
			Name:        ct.Name,
			FlagCode:    ct.FlagCode,
			ShortCode:   ct.ShortCode,
			PhonePrefix: ct.PhonePrefix,
		}
	}
	response.Success(c, http.StatusOK, result)
}

// @Summary Get all countries alphabetically
// @Description Get a list of all countries sorted by name
// @Tags Countries
// @Produce json
// @Success 200 {object} response.Response{data=[]dtos.CountryResponse}
// @Router /countries/alphabetical [get]
func (h *CountryHandler) GetAllAlphabetical(c *gin.Context) {
	countries, err := h.service.GetAllCountriesAlphabetical()
	if err != nil {
		c.Error(err)
		return
	}
	result := make([]dtos.CountryResponse, len(countries))
	for i, ct := range countries {
		result[i] = dtos.CountryResponse{
			Id:          ct.Id,
			Name:        ct.Name,
			FlagCode:    ct.FlagCode,
			ShortCode:   ct.ShortCode,
			PhonePrefix: ct.PhonePrefix,
		}
	}
	response.Success(c, http.StatusOK, result)
}

// @Summary Get a country by ID
// @Description Get details of a country by its ID
// @Tags Countries
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} response.Response{data=dtos.CountryResponse}
// @Failure 404 {object} response.Response
// @Router /countries/{id} [get]
func (h *CountryHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	countryId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	country, err := h.service.GetCountryById(countryId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.CountryResponse{
		Id:          country.Id,
		Name:        country.Name,
		FlagCode:    country.FlagCode,
		ShortCode:   country.ShortCode,
		PhonePrefix: country.PhonePrefix,
	})
}

// @Summary Update a country
// @Description Update an existing country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Country ID"
// @Param request body dtos.CountryUpdateRequest true "Updated details"
// @Success 200 {object} response.Response{data=dtos.CountryResponse}
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /countries/{id} [put]
func (h *CountryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	countryId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	var req dtos.CountryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	err = h.service.UpdateCountry(countryId, req.Name, req.FlagCode, req.ShortCode, req.PhonePrefix)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.CountryResponse{
		Id:          countryId,
		Name:        req.Name,
		FlagCode:    req.FlagCode,
		ShortCode:   req.ShortCode,
		PhonePrefix: req.PhonePrefix,
	})
}

// @Summary Delete a country
// @Description Delete a country by its ID
// @Tags Countries
// @Param id path int true "Country ID"
// @Success 204 "No Content"
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /countries/{id} [delete]
func (h *CountryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	countryId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	err = h.service.DeleteCountry(countryId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}