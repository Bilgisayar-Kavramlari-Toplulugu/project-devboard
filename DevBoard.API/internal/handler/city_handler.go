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

type CityHandler struct {
	service   services.CityService
	validator *validator.Validator
}

func NewCityHandler(service services.CityService, validator *validator.Validator) *CityHandler {
	return &CityHandler{service: service, validator: validator}
}

// @Summary Create a new city
// @Description Create a new city
// @Tags Cities
// @Accept json
// @Produce json
// @Param request body dtos.CityCreateRequest true "City details"
// @Success 201 {object} response.Response{data=dtos.CityResponse}
// @Failure 400 {object} response.Response
// @Security BearerAuth
// @Router /cities [post]
func (h *CityHandler) Create(c *gin.Context) {
	var req dtos.CityCreateRequest
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
	cityId, err := h.service.CreateCity(req.Name, req.Code, req.DisplayOrder, req.CountryId, userID)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusCreated, dtos.CityResponse{
		Id:           cityId,
		Name:         req.Name,
		Code:         req.Code,
		DisplayOrder: req.DisplayOrder,
		CountryId:    req.CountryId,
	})
}

// @Summary Get all cities
// @Description Get a list of all cities
// @Tags Cities
// @Produce json
// @Success 200 {object} response.Response{data=[]dtos.CityResponse}
// @Router /cities [get]
func (h *CityHandler) GetAll(c *gin.Context) {
	cities, err := h.service.GetAllCities()
	if err != nil {
		c.Error(err)
		return
	}
	result := make([]dtos.CityResponse, len(cities))
	for i, ct := range cities {
		code := ""
		if ct.Code != nil {
			code = *ct.Code
		}
		result[i] = dtos.CityResponse{
			Id:           ct.Id,
			Name:         ct.Name,
			Code:         code,
			DisplayOrder: ct.DisplayOrder,
			CountryId:    ct.CountryId,
		}
	}
	response.Success(c, http.StatusOK, result)
}

// @Summary Get cities by country ID
// @Description Get all cities for a specific country
// @Tags Cities
// @Produce json
// @Param countryId path int true "Country ID"
// @Success 200 {object} response.Response{data=[]dtos.CityResponse}
// @Router /cities/country/{countryId} [get]
func (h *CityHandler) GetByCountryId(c *gin.Context) {
	id := c.Param("countryId")
	countryId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	cities, err := h.service.GetCitiesByCountryId(countryId)
	if err != nil {
		c.Error(err)
		return
	}
	result := make([]dtos.CityResponse, len(cities))
	for i, ct := range cities {
		code := ""
		if ct.Code != nil {
			code = *ct.Code
		}
		result[i] = dtos.CityResponse{
			Id:           ct.Id,
			Name:         ct.Name,
			Code:         code,
			DisplayOrder: ct.DisplayOrder,
			CountryId:    ct.CountryId,
		}
	}
	response.Success(c, http.StatusOK, result)
}

// @Summary Get a city by ID
// @Description Get details of a city by its ID
// @Tags Cities
// @Produce json
// @Param id path int true "City ID"
// @Success 200 {object} response.Response{data=dtos.CityResponse}
// @Failure 404 {object} response.Response
// @Router /cities/{id} [get]
func (h *CityHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	cityId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	city, err := h.service.GetCityById(cityId)
	if err != nil {
		c.Error(err)
		return
	}
	code := ""
	if city.Code != nil {
		code = *city.Code
	}
	response.Success(c, http.StatusOK, dtos.CityResponse{
		Id:           city.Id,
		Name:         city.Name,
		Code:         code,
		DisplayOrder: city.DisplayOrder,
		CountryId:    city.CountryId,
	})
}

// @Summary Update a city
// @Description Update an existing city
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "City ID"
// @Param request body dtos.CityUpdateRequest true "Updated details"
// @Success 200 {object} response.Response{data=dtos.CityResponse}
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /cities/{id} [put]
func (h *CityHandler) Update(c *gin.Context) {
	id := c.Param("id")
	cityId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	var req dtos.CityUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	err = h.service.UpdateCity(cityId, req.Name, req.Code, req.DisplayOrder, req.CountryId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.CityResponse{
		Id:           cityId,
		Name:         req.Name,
		Code:         req.Code,
		DisplayOrder: req.DisplayOrder,
		CountryId:    req.CountryId,
	})
}

// @Summary Delete a city
// @Description Delete a city by its ID
// @Tags Cities
// @Param id path int true "City ID"
// @Success 204 "No Content"
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /cities/{id} [delete]
func (h *CityHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	cityId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	err = h.service.DeleteCity(cityId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}