package httpctrl

import (
	"github.com/gin-gonic/gin"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/usecase"

	"github.com/sirupsen/logrus"

	_ "github.com/himmel520/uoffer/require/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


type Handler struct {
	uc  *usecase.Usecase
	log *logrus.Logger
}

func New(uc *usecase.Usecase, log *logrus.Logger) *Handler {
	return &Handler{uc: uc, log: log}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Swagger
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// user
		analytic := api.Group("/analytics").Use(h.jwtAuthAccess(entity.RoleUser))
		{
			analytic.GET("/post/:id", h.getAnalyticWithWordsByPostID) // Get analytics for a post
		}

		category := api.Group("/categories").Use(h.jwtAuthAccess(entity.RoleUser))
		{
			category.GET("/posts", h.getCategoriesWithPublicPosts) // Get all public posts with categories
		}

		// admin
		admin := api.Group("/admin", h.jwtAuthAccess(entity.RoleAdmin))
		{
			category := admin.Group("/categories", h.deleteCategoriesCache())
			{
				category.GET("/posts", h.getCategoriesWithPosts) // Get all posts with categories
				category.GET("/", h.getAllCategories)            // Get all categories
				category.POST("/", h.addCategory)                // Add a new category
				category.DELETE("/:category", h.deleteCategory)  // Delete a category
				category.PUT("/:category", h.updateCategory)     // Update a category
			}

			post := admin.Group("/posts", h.deleteCategoriesCache())
			{
				post.POST("/", h.addPost) // Add a new post

				post.Use(h.validateID())
				post.PUT("/:id", h.updatePost)    // Update a post
				post.DELETE("/:id", h.deletePost) // Delete a post
			}

			analytic := admin.Group("/analytics")
			{
				analytic.POST("/", h.addAnalytic) // Add analytics for a post

				analytic.Use(h.validateID())
				analytic.PUT("/:id", h.updateAnalytic)    // Update analytics by ID
				analytic.DELETE("/:id", h.deleteAnalytic) // Delete analytics by ID
			}

			filter := admin.Group("/filters")
			{
				filter.GET("/", h.getFilters)
				filter.POST("/", h.addFilter)
				filter.DELETE("/:filter", h.deleteFilter)
			}
		}
	}

	return r
}
