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
	Name          string
	Description   string
	Contact       string
	Qualification int
	Department    string
}

type Stores struct {
	Start   *Vstore
	Lastest *Vstore
	Size    int
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
