package handler

/* IMPLEMENTAR REQUEST Se debe implementar el controlador de productos. Primero generamos el
request con los valores que esperamos recibir en la petición e importamos
nuestro paquete interno de productos*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/WendyCuy/bootcamp-go/go-web/clase4_tt/internal/products"
	"github.com/WendyCuy/bootcamp-go/go-web/clase4_tt/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

/* ESTRUCTURA CONTROLADOR. Implementaremos la estructura Producto y una
función que reciba un Servicio (del paquete interno) y devuelva el controlador
instanciado */

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

/* Método obtener productos. Se encargará de realizar las validaciones
de la petición, pasarle la tarea al Servicio y retornar la respuesta
correspondiente al cliente.*/

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		fmt.Printf("Token desde nuestro env: %s\n", tokenFromEnv)

		if token != tokenFromEnv {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		if len(p) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No hay productos almacenados"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

/* Método guardar.  De la misma forma el método guardar */

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Validar token
		token := ctx.GetHeader("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}
		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La cantidad es requerida"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}
		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

/* Se agrega el controlador Update en el handler de productos */

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products [put]
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		//Validaciones
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}
		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La cantidad es requerida"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}

		//Hago el Update
		p, err := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// UpdateNameProducts godoc
// @Summary UpdateName products
// @Tags Products
// @Description updatename products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to updateName"
// @Success 200 {object} web.Response
// @Router /products [patch]
func (c *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		//Validaciones
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		//Hago el patch del producto
		p, err := c.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// DeleteProducts godoc
// @Summary Delete products
// @Tags Products
// @Description delete products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to delete"
// @Success 200 {object} web.Response
// @Router /products [delete]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalido"))
			return
		}

		//Hago la baja fisica del producto
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalido"))
			return
		}
		ctx.JSON(200, web.NewResponse(400, nil, "El producto ha sido eliminado"))
	}
}
