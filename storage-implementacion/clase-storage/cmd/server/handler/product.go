package handler

import (
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/product"
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(p product.Service) *Product {
	return &Product{
		productService: p,
	}
}

func (p *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			web.Error(c, 400, "Error")
			return
		}
		pr, err := p.productService.GetByName(name)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		web.Success(c, 200, pr)
	}
}
