package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/internal/products"
	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_SaveProduct_OK(t *testing.T) {
	objReq := struct {
		Code string           `json:"code"`
		Data products.Product `json:"data"`
	}{}
	// crear el Server y definir las rutas
	r := createServer()
	// crear request de tipo post y response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 99.99
    }`)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.Code, "200")
	assert.Equal(t, objReq.Data.Price, 99.99)
	assert.Equal(t, objReq.Data.Count, 10)
	assert.Equal(t, objReq.Data.Type, "Funcional")
	assert.Equal(t, objReq.Data.Name, "Tester")
}

func Test_GetProduct_OK(t *testing.T) {
	objReq := struct {
		Code string             `json:"code"`
		Data []products.Product `json:"data"`
	}{}
	// crear el Server y definir las rutas
	r := createServer()
	// crear request de tipo GET y response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.Code, "200")
	assert.Equal(t, len(objReq.Data) > 0, true)
}

func Test_UpdateProduct_OK(t *testing.T) {
	// crear el Server y definir las rutas
	r := createServer()
	// crear Request del tipo Put y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/products/4", `{
        "nombre": "Telefono","tipo": "Funcional","cantidad": 10,"precio": 10.99
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_DeleteProduct_OK(t *testing.T) {
	// crear el Server y definir las rutas
	r := createServer()
	// crear Request del tipo DELETE y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/products/3", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
