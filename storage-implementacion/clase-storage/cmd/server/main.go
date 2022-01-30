package main

/* Se importa el package database/sql y el driver*/
import (
	"database/sql" // Es la librería con la que se interactura con bases de datos SQL
	"fmt"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Es la librería del driver
)

func main() {
	/*Ahora se puede ejecutar la conexión a la nuestra base de datos.
	  Se invoca la función Open del package sql la cual recibe por parámetro
	  el nombre del driver y los datos*/
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	// Open inicia un pool de conexiones. Sólo abrir una vez
	if err != nil {
		fmt.Println("No se pudo conectar a la db")
	} else {
		fmt.Println("Conexion a la db exitosa !")
	}

	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	r.Run()

}
