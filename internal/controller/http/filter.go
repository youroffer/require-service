package httpctrl

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

// @Summary Добавить фильтр
// @Description Добавляет новый фильтр
// @Tags Filters
// @Accept json
// @Produce json
// @Param filter query string true "Название фильтра"
// @Success 201 {object} models.Filter
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /admin/filters [post]
func (h *Handler) addFilter(c *gin.Context) {
	var query filterQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	filterResp, err := h.uc.Filter.Add(c.Request.Context(), query.Filter)
	if err != nil {
		if errors.Is(err, repoerr.ErrFilterExist) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, filterResp)
}

// @Summary Удалить фильтр
// @Description Удаляет фильтр по его названию
// @Tags Filters
// @Param filter path string true "Название фильтра"
// @Success 204 "No Content"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /filters/{filter} [delete]
func (h *Handler) deleteFilter(c *gin.Context) {
	var uri filterURI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	err := h.uc.Filter.Delete(c.Request.Context(), uri.Filter)
	if err != nil {
		if errors.Is(err, repoerr.ErrFilterNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Получить список фильтров
// @Description Возвращает список фильтров с поддержкой пагинации
// @Tags Filters
// @Param limit query int false "Количество фильтров"
// @Param offset query int false "Смещение для фильтров"
// @Success 200 {object} models.FilterResp
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /filters [get]
func (h *Handler) getFilters(c *gin.Context) {
	var query *PaginationQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{fmt.Sprintf("неккоректный query: %v", err)})
		return
	}

	filterResp, err := h.uc.Filter.GetWithPagination(c.Request.Context(), query.Limit, query.Offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	if len(filterResp.Filters) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"filters not found"})
		return

	}

	c.JSON(http.StatusOK, filterResp)
}
