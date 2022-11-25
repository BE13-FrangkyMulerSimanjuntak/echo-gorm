package controllers

import (
	"frangky/be13/helper"
	"frangky/be13/models"
	"frangky/be13/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {

	result, err := repository.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("ERROR ON READ DATA !"))
	}
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("SUCCESS CREATE USER", result))
}

func AddUserController(c echo.Context) error {
	//
	Inputuser := models.User{}
	errbind := c.Bind(&Inputuser)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("ERROR ON READ DATA"+errbind.Error()))
	}
	errInsert := repository.InsertUser(Inputuser)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("FAIL INSERT DATA"))
	}
	return c.JSON(http.StatusOK, helper.PesanSuksesHelper("SUCCES CREATE USER"))
}

func GetAllUserControllerId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := repository.GetAllUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("ERROR ON READ DATA"))
	}
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("SUCCES ON READ USER", result))
}

func UpddateUserController(c echo.Context) error {
	//

	data := models.User{}
	c.Bind(&data)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("FAIL ON READ BY ID"))
	}
	result, errupdate := repository.UpdateUserId(id, data)
	if errupdate != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("FAIL DELETE DATA"))
	}
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("SUCCES DELETE USERS", result))
}

func DeleteUserContoller(c echo.Context) error {
	//
	id, _ := strconv.Atoi(c.Param("id"))

	errdelete := repository.DeleteUser(id)
	if errdelete != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("FAIL DELETE DATA"))
	}
	return c.JSON(http.StatusOK, helper.PesanSuksesHelper("SUCCES CREATE USER"))
}
