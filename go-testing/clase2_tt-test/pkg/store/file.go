package store

import (
	"encoding/json"
	"io/ioutil"
)

/* Se implementa la interface Store con los métodos Read y Write, ambos
métodos reciben una interface y devolveran un error.

Se agrega Mock para realizar test*/

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	//AddMock(mock *Mock)
	//ClearMock()
}

/* Se debe implementar una constante de tipo Type para definir el tipo
de store que se utilizará, en este caso solo será por archivo (FileType)*/

type Type string

const (
	FileType Type = "filestorage"
)

/* Se define el mock para hacer el test */
type Mock struct {
	Data []byte
	Err  error
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

/* Método Write . Se utiliza para escribir datos de la estructura en el archivo.
Simplemente recibe una interface y lo convertirá a una representación
JSON en bytes para guardarlo en el archivo que especificamos al
momento de instanciar la función Factory*/

func (fs *FileStore) Write(data interface{}) error {
	if fs.Mock != nil {
		return fs.Mock.Err
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fs.FileName, file, 0644)

}

/* Método Read.  Sirve para leer el archivo y guardar su contenido
empleando la interface que recibirá como parámetro.*/

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}

	file, err := ioutil.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}

/*Factory de Store
Se debe implementar la función Factory que se encarga de generar la
estructura que se desea y recibe el tipo de store que se quiere implementar
y el nombre del archivo.
Se declara la estructura FileStore con el campo que guarde el nombre del
archivo.*/

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName, nil}
	}
	return nil
}

/* Se agrega mock al filestore */
type FileStore struct {
	FileName string
	Mock     *Mock
}
