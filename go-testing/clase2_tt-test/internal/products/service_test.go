package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/pkg/store"
	"github.com/stretchr/testify/assert"
)

/* El test de integración se diseñará usando como objeto de prueba la capa
"service" y su integración con "repository" y tambien se usará un Stub de
FileStorage.  En este ejemplo para comenzar se testearan GetAll y Storre
para respuestas válidas y también cuando FileStorage devuelve error.

En el test unitario de GetAll de "Servicio" vamos a crear todas las
estructuras necesarias para después proveer a la función y en este caso esas
estructuras van a ser stub*/

type DummyRepo struct{}

func (dr *DummyRepo) GetAll() ([]Product, error) {
	return []Product{}, nil
}
func (dr *DummyRepo) Store(id int, name, productType string, count int, price float64) (Product, error) {
	return Product{}, nil
}
func (dr *DummyRepo) LastID() (int, error) {
	return 0, nil
}
func (dr *DummyRepo) UpdateName(id int, name string) (Product, error) {
	return Product{}, nil
}
func (dr *DummyRepo) Update(id int, name, productType string, count int, price float64) (Product, error) {
	return Product{}, nil
}
func (dr *DummyRepo) Delete(id int) error {
	return nil
}

type mockStore struct {
	Data          []Product
	ReadWasCalled bool
}

func (s *mockStore) Read(data interface{}) error {
	s.ReadWasCalled = true

	fileData, _ := json.Marshal(s.Data)

	return json.Unmarshal(fileData, &data)
}

func (s *mockStore) Write(data interface{}) error {
	return nil
}

/*De forma análoga al stub q se diseño para el test unitario, se hace
lo mismo para el test de integración entre repo y el service*/
func TestServiceGetAll(t *testing.T) {
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
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	/*Para invocar la ejecución del test, es necesario instanciar el
	Service con el repositorio que contiene el stub del storage. Luego
	se ejecuta el test y se valida en este caso que los resultados sean
	exactamente igual a lo esperado y que el error sea nil*/
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo) // Se llamó al GetAll de service

	result, err := myService.GetAll()

	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

/* Con esto se prueba la integración cuando la respuesta desde FileStore
es errónea.

Para esto se debe usar otro Stub.  Donde se establece que la data en el
FileStore es nil y el error es igual a "expectedError*/
func TestServiceGetAllError(t *testing.T) {
	// Initializing Input/output
	expectedError := errors.New("error for GetAll")
	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

/* En este test de integración se comprobará que desde el service se
pueda almacenar información correctamente.  Para esto se define el Stub
inicial vacío y se ejecuta el método Store. La respuesta debe retornar un
producto con las mismas características y con ID=1.*/
func TestStore(t *testing.T) {
	testProduct := Product{
		ID:    20,
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	db := mockStore{[]Product{testProduct}, false}
	repo := NewRepository(&db)

	result, _ := repo.Store(testProduct.ID, testProduct.Name, testProduct.Type, testProduct.Count, testProduct.Price)

	assert.Equal(t, testProduct.Name, result.Name)
	assert.Equal(t, testProduct.Type, result.Type)
	assert.Equal(t, testProduct.Price, result.Price)

}

/* Con esta integración se comprueba que si ocurre un error durante la
escritura de FileStore, el service reciba del repositorio el error correcto
y además que retorne un objeto vacío de Product*/
func TestStoreError(t *testing.T) {
	testProduct := Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	expectedError := errors.New("error for Storage")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.Store(testProduct.Name, testProduct.Type, testProduct.Count, testProduct.Price)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, Product{}, result)
}

func TestServiceUpdate(t *testing.T) {
	testProduct := Product{
		ID:    1,
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 250.0,
	}
	productoEsperado := Product{
		ID:    1,
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 4,
		Price: 100.0,
	}

	input := []Product{
		testProduct,
	}
	dataJson, _ := json.Marshal(input)

	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	// Resultado esperado
	result, err := myService.Update(testProduct.ID, testProduct.Name, testProduct.Type, 4, 100.0)

	assert.Nil(t, err, "Error en el update")
	assert.Equal(t, productoEsperado, result)
	// assert.Equal(t, testProduct.Name, result.Name)
	// assert.Equal(t, testProduct.Type, result.Type)
	// assert.Equal(t, 100.0, result.Price)
	// assert.Equal(t, 1, result.ID)
}

func TestServiceDelete(t *testing.T) {
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
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	// Resultado esperado
	resultOk := myService.Delete(1)
	resultNotOk := myService.Delete(3)

	assert.Nil(t, resultOk, "Hubo un error en el delete")
	assert.NotNil(t, resultNotOk, "Hubo un error en el delete")
}
