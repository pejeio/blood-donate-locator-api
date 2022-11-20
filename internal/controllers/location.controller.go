package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pejeio/blood-donate-locator-api/internal/models"
	"gorm.io/gorm"
)

type LocationController struct {
	DB *gorm.DB
}

func NewLocationController(DB *gorm.DB) LocationController {
	return LocationController{DB}
}

func (lc *LocationController) CreateLocation(ctx *gin.Context) {
	var payload *models.CreateLocationRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newLocation := models.Location{
		Name:        payload.Name,
		Coordinates: payload.Coordinates,
		Address:     payload.Address,
	}
	result := lc.DB.Create(&newLocation)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
	}

	ctx.JSON(http.StatusCreated, newLocation)
}

func (lc *LocationController) FindLocations(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var locations []models.Location

	var totalCount int64
	lc.DB.Find(&locations).Count(&totalCount)

	lc.DB.
		Limit(intLimit).
		Offset(offset).
		Order("created_at DESC").
		Find(&locations)

	ctx.JSON(http.StatusOK, models.ResponseWithPagination{
		Data: locations,
		Meta: models.ResponseMeta{Count: totalCount},
	})
}
