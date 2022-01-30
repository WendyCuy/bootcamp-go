package routes

import (
	"database/sql"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/cmd/server/handler"
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/product"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildProductRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildProductRoutes() {
	repository := product.NewRepository(r.db)
	service := product.NewService(repository)
	handler := handler.NewProduct(service)
	routes := r.rg
	{
		routes.GET("/products/name", handler.GetByName())
		routes.POST("/", handler.Store())
		routes.GET("/products/id", handler.GetOne())
		routes.PUT("/products/id", handler.Update())
		routes.GET("/products", handler.GetAll())
		routes.DELETE("/products/id", handler.Delete())
		routes.PUT("/products/eficcient/:id", handler.UpdateWithContext())
	}

}
