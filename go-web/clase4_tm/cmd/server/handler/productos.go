package handler

/* IMPLEMENTAR REQUEST Se debe implementar el controlador de productos. Primero generamos el
request con los valores que esperamos recibir en la petición e importamos
nuestro paquete interno de productos*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/WendyCuy/bootcamp-go/go-web/clase3_tt/internal/products"
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

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		fmt.Printf("Token desde nuestro env: %s\n", tokenFromEnv)

		if token != tokenFromEnv {
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
		//Validar token
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
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

/* Se agrega el controlador Update en el handler de productos */

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.Type == "" {
			ctx.JSON(400, gin.H{"error": "El tipo del producto es requerido"})
			return
		}
		if req.Count == 0 {
			ctx.JSON(400, gin.H{"error": "La cantidad es requerida"})
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}

		//Hago el Update
		p, err := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		//Hago el patch del producto
		p, err := c.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		//Hago la baja fisica del producto
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
