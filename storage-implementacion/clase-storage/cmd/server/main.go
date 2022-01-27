package main

/* Se importa el package database/sql y el driver*/
import (
	"database/sql" // Es la librería con la que se interactura con bases de datos SQL

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Es la librería del driver
)

//Se crea la variable donde se almacena la base de datos
var (
	StorageDB *sql.DB
)

func main() {
	/*Ahora se puede ejecutar la conexión a la nuestra base de datos.
	  Se invoca la función Open del package sql la cual recibe por parámetro
	  el nombre del driver y los datos*/
	dataSource := "root:root@tcp(localhost:3306)/storage" // Se tiene conection string, usuario root que se conecta al localhost 3306 de la base de datos storage
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource) // Se abre la conexión con el driver mysql a este dataSource y cuando este abierta me guarda en StorageDB.

	r := gin.Default()
	router := routes.NewRouter(r, StorageDB)
	router.MapRoutes()

	if err != nil {
		panic(err)
	}
}
