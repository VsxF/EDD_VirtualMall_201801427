package data

import "fmt"

func Listdd() {
	lista := NewStoresList()
	lista.setStore("nombre", "descrip", "contac", 5, "dep")
	fmt.Println(lista.Start)
}

type Vstore struct {
	Previous      *Vstore
	Next          *Vstore
	Name          string `json:"Nombre"`
	Description   string `json:"Descripcion"`
	Contact       string `json:"Contacto"`
	Qualification int `json:"Calificacion"`
	Department    string `json:"Departamento"`
}

type Stores struct {
	Start   *Vstore
	Lastest *Vstore
	Size    int
}

func NewVstore() *Vstore {
	return &Vstore{}
}

func NewStoresList() *Stores {
	return &Stores{nil, nil, 0}
}

//Insertar nueva tienda
func (stores *Stores) setStore(name string, description string, contact string, qualification int, dep string) {
	newStore := &Vstore{nil, nil, name, description, contact, qualification, dep}

	if stores.Start == nil {
		stores.Start = newStore
		stores.Lastest = newStore
	} else {
		stores.Lastest.Next = newStore
		newStore.Previous = stores.Lastest
		stores.Lastest = newStore
	}
	stores.Size++
}
