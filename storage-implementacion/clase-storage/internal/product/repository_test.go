package product

import (
	"database/sql"
	"testing"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/models"
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {

	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	newProduct := models.Product{
		Name:  "Perfume",
		Type:  "Belleza",
		Count: 10,
		Price: 100.10,
	}

	serv := NewRepository(db)

	res, _ := serv.GetByName(newProduct.Name)

	assert.Equal(t, newProduct.Name, res.Name)

}

func TestSqlRepositoryStore(t *testing.T) {

	db, err := util.InitDb()

	assert.NoError(t, err)

	repository := NewRepository(db)

	newProduct := models.Product{
		Name:  "Perfume",
		Type:  "Belleza",
		Count: 10,
		Price: 100.10,
	}

	respStore, err := repository.Store(newProduct)
	assert.NoError(t, err)
	assert.NotEqual(t, newProduct, respStore)

	respGet, err := repository.GetOne(newProduct.ID)

	assert.NoError(t, err)
	assert.Equal(t, newProduct.Name, respGet.Name)

}
