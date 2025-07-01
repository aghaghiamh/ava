package userhandler

import (
	"net/http"
	"strconv"

	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/httpmapper"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetProfileHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	profileReq := domain.ProfileRequest{
		UserID: uint(id),
	}

	profileResp, respErr := h.userSvc.GetProfile(c.Request().Context(), profileReq)
	if respErr != nil {
		code, msg := httpmapper.MapResponseCustomErrorToHttp(respErr)

		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, profileResp)
}
