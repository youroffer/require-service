package http

import (
	"context"
	"crypto/rsa"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/require/internal/config"
	"github.com/himmel520/uoffer/require/models"
	"github.com/sirupsen/logrus"

	_ "github.com/himmel520/uoffer/require/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:generate mockery --name=Service --with-expecter

type (
	Service interface {
		Auth
		Cache
		
		Category
		Post
		Analytic
		Filter
	}

	Category interface {
		GetAllCategories(ctx context.Context) ([]*models.Category, error)
		GetCategoriesWithPosts(ctx context.Context) (map[string][]*models.PostResponse, error)
		GetCategoriesWithPublicPosts(ctx context.Context) (map[string][]*models.PostResponse, error)
		AddCategory(ctx context.Context, category *models.Category) (*models.Category, error)
		UpdateCategory(ctx context.Context, category, title string) (*models.Category, error)
		DeleteCategory(ctx context.Context, category string) error
	}

	Post interface {
		AddPost(ctx context.Context, post *models.Post) (*models.PostResponse, error)
		UpdatePost(ctx context.Context, id int, post *models.PostUpdate) (*models.PostResponse, error)
		DeletePost(ctx context.Context, id int) error
	}

	Analytic interface {
		AddAnalytic(ctx context.Context, analytic *models.Analytic) (*models.Analytic, error)
		UpdateAnalytic(ctx context.Context, id int, analytic *models.AnalyticUpdate) (*models.Analytic, error)
		DeleteAnalytic(ctx context.Context, id int) error
		GetAnalyticWithWords(ctx context.Context, postID int, role string) (*models.AnalyticWithWords, error)
	}

	Auth interface {
		GetUserRoleFromToken(jwtToken string, publicKey *rsa.PublicKey) (string, error)
		IsUserAuthorized(requiredRole, userRole string) bool
	}

	Filter interface {
		AddFilter(ctx context.Context, filter string) (*models.Filter, error)
		DeleteFilter(ctx context.Context, word string) error
		GetFilters(ctx context.Context, limit, offset int) (*models.FilterResp, error)
	}
)

type Cache interface {
	DeleteCacheCategoriesAndPosts(ctx context.Context) error
}

type Handler struct {
	srv Service
	log *logrus.Logger
	cfg *config.JWT
}

func New(srv Service, cfg *config.JWT, log *logrus.Logger) *Handler {
	return &Handler{srv: srv, log: log, cfg: cfg}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Swagger
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// user
		analytic := api.Group("/analytics").Use(h.jwtAuthAccess(models.RoleUser))
		{
			analytic.GET("/post/:id", h.getAnalyticWithWordsByPostID) // Get analytics for a post
		}

		category := api.Group("/categories").Use(h.jwtAuthAccess(models.RoleUser))
		{
			category.GET("/posts", h.getCategoriesWithPublicPosts) // Get all public posts with categories
		}

		// admin
		admin := api.Group("/admin", h.jwtAuthAccess(models.RoleAdmin))
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
