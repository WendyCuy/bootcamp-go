package store

import (
	"encoding/json"
	"os"
)

/* Se implementa la interface Store con los métodos Read y Write, ambos
métodos reciben una interface y devolveran un error.*/

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

/* Se debe implementar una constante de tipo Type para definir el tipo
de store que se utilizará, en este caso solo será por archivo (FileType)*/

type Type string

const (
	FileType Type = "file"
)

/* Método Write . Se utiliza para escribir datos de la estructura en el archivo.
Simplemente recibe una interface y lo convertirá a una representación
JSON en bytes para guardarlo en el archivo que especificamos al
momento de instanciar la función Factory*/

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

/* Método Read.  Sirve para leer el archvio y guardar su contenido
empleando la interface que recibirá como parámetro.*/

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
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
		return &FileStore{fileName}
	}
	return nil
}

type FileStore struct {
	FileName string
}
