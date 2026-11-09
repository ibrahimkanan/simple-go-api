package handlers

import (
	"myapp/cmd/api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func PostIndexHanlder(c echo.Context)error{
	data, err := service.GetAll()
	if err != nil{
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandle(c echo.Context) error{
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unalbe to process")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unalbe to process")
	}
	res := make(map[string]any)
	res["status"] = "OK"
	res["datat"] = data
	
	return c.JSON(http.StatusOK, res)
}