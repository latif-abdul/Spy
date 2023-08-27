package controller

import (
	"net/http"
	"spy-api/model"

	"github.com/labstack/echo/v4"
	"github.com/latif-abdul/Spy/spy-api/model"
	"github.com/latif-abdul/Spy/spy-api/storage"
)

func GetIpAddress(c echo.Context) error {
	ipaddress, _ := GetRepoIpAddress()
	return c.JSON(http.StatusOK, ipaddress)
}

func GetRepoIpAddress() ([]model.IpAddress, error) {
	db := storage.GetDBInstance()
	ipaddress := []model.IpAddress{}

	if err := db.Find(&ipaddress).Error; err != nil {
		return nil, err
	}

	return ipaddress, nil
}
