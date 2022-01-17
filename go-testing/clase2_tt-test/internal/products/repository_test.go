package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/pkg/store"
	"github.com/stretchr/testify/assert"
)

/* Se crea el archivo respositoy_test.go en la misma carpeta que esta
repository.go.*/

const errorGetAll = "error for GetAll"

func TestGetAllError(t *testing.T) {
	// Initializing Input/output
	expectedError := errors.New(errorGetAll)
	dbMock := store.Mock{
		Data: nil,
		Err:  expectedError,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)

	_, err := myRepo.GetAll() // Se llamó al GetAll del repositorio

	assert.Equal(t, err, expectedError)
}

/* En este caso se crea la función "TestGetAll.
Para inicializar es necesario crear el stub. Primero se crea el conjunto de
productos que deseamos usar, y ese slice llamado input se convierte a []byte
porque asi lo requiere FileStore en el método Read.*/
func TestGetAll(t *testing.T) {
	/* Initializing Input/output, en este caso se crea la lista de
	productos q va a emular lo q se lee desde el archivo product.json */
	input := []Product{
		{
			ID:    1,
			Name:  "CellPhone",
			Type:  "Tech",
			Count: 3,
			Price: 250,
		}, {
			ID:    2,
			Name:  "Notebook",
			Type:  "Tech",
			Count: 10,
			Price: 1750.5,
		},
	}
	dataJson, _ := json.Marshal(input)
	dbStub := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}

	/* Para probar el Repository, es necesario instanciarlo con el mock
	o stub de storage.  Posteriormente se ejecuta el test y se valida
	que la información que devuelve el repository sea igual a la que
	se definió stub del Storage*/
	myRepo := NewRepository(&storeMocked)
	// Test Excution
	resp, _ := myRepo.GetAll()
	// Validation
	assert.Equal(t, resp, input)
}
