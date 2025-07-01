package userhandler

import (
	"net/http"
	"strconv"

	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/httpmapper"
	"github.com/labstack/echo/v4"
)

func (h Handler) DeleteHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delReq := domain.DelAccountRequest{
		UserID: uint(id),
	}

	resp, regErr := h.userSvc.Delete(delReq)
	if regErr != nil {
		code, msg := httpmapper.MapResponseCustomErrorToHttp(regErr)

		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
