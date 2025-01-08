package httpctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

// @Summary Получить категории с публичными постами
// @Description Возвращает список категорий с публичными постами
// @Tags Categories
// @Produce json
// @Success 200 {object} map[string][]entity.PostResponse
// @Failure 404 {object} errorResponse "Post Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /categories/public-posts [get]
func (h *Handler) getCategoriesWithPublicPosts(c *gin.Context) {
	response, err := h.uc.Category.GetAllWithPublicPosts(c.Request.Context())
	if err != nil {
		if errors.Is(err, repoerr.ErrPostNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Получить категории с постами
// @Description Возвращает список категорий с постами
// @Tags Categories
// @Produce json
// @Success 200 {object} map[string][]entity.PostResponse
// @Failure 404 {object} errorResponse "Post Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /admin/categories/posts [get]
func (h *Handler) getCategoriesWithPosts(c *gin.Context) {
	response, err := h.uc.Category.GetAllWithPosts(c.Request.Context())
	if err != nil {
		if errors.Is(err, repoerr.ErrPostNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Получить все категории
// @Description Возвращает список всех категорий
// @Tags Categories
// @Produce json
// @Success 200 {object} []entity.Category
// @Failure 404 {object} errorResponse "Category Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /admin/categories [get]
func (h *Handler) getAllCategories(c *gin.Context) {
	response, err := h.uc.Category.GetAll(c.Request.Context())
	if err != nil {
		if errors.Is(err, repoerr.ErrCategoryNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// @Summary Добавить новую категорию
// @Description Добавляет новую категорию
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body entity.Category true "Данные новой категории"
// @Success 201 {object} entity.Category
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 409 {object} errorResponse "Category Already Exists"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /admin/categories [post]
func (h *Handler) addCategory(c *gin.Context) {
	category := &entity.Category{}
	if err := c.BindJSON(category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newCategory, err := h.uc.Category.Add(c.Request.Context(), category)
	if err != nil {
		if errors.Is(err, repoerr.ErrCategoryExists) {
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
			return
		}

		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

// @Summary Обновить категорию
// @Description Обновляет название категории
// @Tags Categories
// @Accept json
// @Produce json
// @Param category path string true "Категория"
// @Param title query string true "Новое название категории"
// @Success 200 {object} entity.Category
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Category Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /admin/categories/{category} [put]
func (h *Handler) updateCategory(c *gin.Context) {
	var uri categoryURI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	var query updateCategoryQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newCategory, err := h.uc.Category.Update(c.Request.Context(), uri.Category, query.Title)
	if err != nil {
		if errors.Is(err, repoerr.ErrCategoryNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}

		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, newCategory)
}

// @Summary Удалить категорию
// @Description Удаляет категорию
// @Tags Categories
// @Param category path string true "Категория"
// @Success 204 "No Content"
// @Failure 404 {object} errorResponse "Category Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /admin/categories/{category} [delete]
func (h *Handler) deleteCategory(c *gin.Context) {
	var uri *categoryURI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if err := h.uc.Category.Delete(c.Request.Context(), uri.Category); err != nil {
		if errors.Is(err, repoerr.ErrCategoryNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}

		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
