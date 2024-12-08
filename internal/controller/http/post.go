package httpctrl

// import (
// 	"errors"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"

// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
// )

// // @Summary Добавить должность
// // @Description Добавляет новую должность в систему
// // @Tags Positions
// // @Accept json
// // @Produce json
// // @Param post body entity.Post true "Данные новой должности"
// // @Success 200 {object} entity.PostResponse
// // @Failure 400 {object} errorResponse "Bad Request"
// // @Failure 409 {object} errorResponse "Conflict - Должность уже существует или зависимость не найдена"
// // @Failure 500 {object} errorResponse "Internal Server Error"
// // @Router /admin/positions [post]
// func (h *Handler) addPost(c *gin.Context) {
// 	var post *entity.Post
// 	if err := c.BindJSON(&post); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
// 		return
// 	}

// 	postResp, err := h.uc.Post.Add(c.Request.Context(), post)
// 	if err != nil {
// 		if errors.Is(err, repoerr.ErrPostDependencyNotFound) || errors.Is(err, repoerr.ErrPostExists) {
// 			c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
// 			return
// 		}

// 		h.log.Error(err.Error())
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, postResp)
// }

// // @Summary Обновить должность
// // @Description Обновляет данные существующей должности по её ID
// // @Tags Positions
// // @Accept json
// // @Produce json
// // @Param id path int true "ID должности"
// // @Param post body entity.PostUpdate true "Данные для обновления должности"
// // @Success 200 {object} entity.PostResponse
// // @Failure 400 {object} errorResponse "Bad Request"
// // @Failure 404 {object} errorResponse "Должность не найдена"
// // @Failure 409 {object} errorResponse "Conflict - Должность уже существует или зависимость не найдена"
// // @Failure 500 {object} errorResponse "Internal Server Error"
// // @Router /admin/positions/{id} [put]
// func (h *Handler) updatePost(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	post := &entity.PostUpdate{}
// 	if err := c.BindJSON(post); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
// 		return
// 	}

// 	if post.IsEmpty() {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"post has no changes"})
// 		return
// 	}

// 	postResp, err := h.uc.Post.Update(c.Request.Context(), id, post)
// 	switch {
// 	case errors.Is(err, repoerr.ErrPostDependencyNotFound) || errors.Is(err, repoerr.ErrPostExists):
// 		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
// 		return
// 	case errors.Is(err, repoerr.ErrPostNotFound):
// 		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
// 		return
// 	case err != nil:
// 		h.log.Error(err.Error())
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, postResp)
// }

// // @Summary Удалить должность
// // @Description Удаляет должность по её ID
// // @Tags Positions
// // @Param id path int true "ID должности"
// // @Success 204 "No Content"
// // @Failure 404 {object} errorResponse "Должность не найдена"
// // @Failure 500 {object} errorResponse "Internal Server Error"
// // @Router /admin/positions/{id} [delete]
// func (h *Handler) deletePost(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	err := h.uc.Post.Delete(c.Request.Context(), id)
// 	switch {
// 	case errors.Is(err, repoerr.ErrPostNotFound):
// 		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
// 		return
// 	case err != nil:
// 		h.log.Error(err.Error())
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }
