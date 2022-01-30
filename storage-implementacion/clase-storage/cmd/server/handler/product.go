package handler

import (
	"strconv"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/product"
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"name"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	productService product.Service
}

func NewProduct(p product.Service) *Product {
	return &Product{
		productService: p,
	}
}

func (s *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			web.Error(c, 400, "Error")
			return
		}
		pr, err := s.productService.GetByName(name)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		web.Success(c, 200, pr)
	}
}

func (s *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request

		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		product, err := s.productService.Store(req.Name, req.Type, req.Count, req.Price)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 201, product)

	}
}

func (s *Product) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		product, err := s.productService.GetOne(id)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, product)
	}
}

func (s *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		s, err := s.productService.Update(int(id), req.Name, req.Type, req.Count, req.Price)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, s)

	}
}

func (s *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		products, err := s.productService.GetAll()

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, products)
	}
}

func (s *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, "invalid ID")
			return
		}

		err = s.productService.Delete(int(id))
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}

		web.Success(c, 204, "Producto eliminado")
	}
}
