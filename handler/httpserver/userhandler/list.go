package userhandler

import (
	"net/http"
	"strconv"

	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/httpmapper"

	"github.com/labstack/echo/v4"
)

// TODO: add them to the config
const (
	DEFAULT_PAGE_SIZE_STR = "20"
	DEFAULT_MAX_PAGE_SIZE = 50
)

func (h *Handler) ListWithPagination(c echo.Context) error {
	// TODO: Move the sanitization in the appropriate package
	pageStr := c.QueryParam("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			// TODO: add to ErrMsg
			"message": "Invalid Page Number",
		})
	}

	pageSizeStr := c.QueryParam("page_size")
	if pageSizeStr == "" {
		pageSizeStr = DEFAULT_PAGE_SIZE_STR
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			// TODO: add to ErrMsg
			"message": "Invalid Page Size",
		})
	}

	// To protect the server check against the default max page size
	if pageSize > DEFAULT_MAX_PAGE_SIZE {
		pageSize = DEFAULT_MAX_PAGE_SIZE
	}

	lReq := domain.ListRequest{
		Page:     page,
		PageSize: pageSize,
	}

	ctx := c.Request().Context()
	users, respErr := h.userSvc.ListWithPagination(ctx, lReq)
	if respErr != nil {
		code, msg := httpmapper.MapResponseCustomErrorToHttp(respErr)

		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, users)
}
