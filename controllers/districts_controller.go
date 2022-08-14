package controllers

import (
	"context"
	"net/http"
	"odisha_gov_be/models"
	services "odisha_gov_be/services/district"
	"odisha_gov_be/utils"

	"github.com/gin-gonic/gin"
)

type DistrictCtrl struct{}

// @Summary Get all items
// @Description Get all items
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.District
// @Router /District [get]
func (ctrl *DistrictCtrl) Get(c *gin.Context) {
	utils.Logger().Info("parsing get request for District")
	districts, err := models.Districts().AllG(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": districts})
}

// @Summary Add item
// @Description Add item
// @Accept  json
// @Produce  json
// @Param role body model.Districttrue "data"
// @Success 200 {object} model.District
// @Router /District [post]
func (ctrl *DistrictCtrl) Post(c *gin.Context) {
	var t models.District
	utils.Logger().Info("value before request", t)
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	utils.Logger().Info("value after request", t)
	t = services.Create()
	utils.Logger().Info("value after create", t)
	c.JSON(http.StatusOK, gin.H{"data": t})
}
