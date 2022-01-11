package handler

/* IMPLEMENTAR REQUEST Se debe implementar el controlador de productos. Primero generamos el
request con los valores que esperamos recibir en la petición e importamos
nuestro paquete interno de productos*/

import (
	"github.com/WendyCuy/bootcamp-go/go-web/clase2_tt/internal/products"
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
		service: p}
}

/* Método obtener productos. Se encargará de realizar las validaciones
de la petición, pasarle la tarea al Servicio y retornar la respuesta
correspondiente al cliente.*/

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

/* Método guardar.  De la misma forma el método guardar */

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
