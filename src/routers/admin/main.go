package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/middlewares"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/users"))
	ProductRouter(r.Group("/products"))
	ProductSizeRouter(r.Group("/productSize"))
	ProductVariantRouter(r.Group("/productVariant"))
	TagsRouter(r.Group("/tags"))
	ProductTagsRouter(r.Group("/productTags"))
	ProductRatingsRouter(r.Group("/productRatings"))
	CategoriesRouter(r.Group("/categories"))
	ProductCategoriesRouter(r.Group("/productCategories"))
	PromoRouter(r.Group("/promo"))

}
