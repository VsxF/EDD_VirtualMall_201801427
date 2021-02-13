package data

import "fmt"

func Listdd() {
	lista := NewStoresList()
	lista.setStore("nombre", "descrip", "contac", 5, "dep")
	fmt.Println(lista.Start)
}

type vstore struct {
	Previous      *vstore
	Next          *vstore
	Name          string
	Description   string
	Contact       string
	Qualification int
	Department    string
}

type stores struct {
	Start   *vstore
	Lastest *vstore
	Size    int
}

func NewStoresList() *stores {
	return &stores{nil, nil, 0}
}

//Insertar nueva tienda
func (stores *stores) setStore(name string, description string, contact string, qualification int, dep string) {
	newStore := &vstore{nil, nil, name, description, contact, qualification, dep}

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

//Imprime las tiendas
func (stores stores) printStores() {

}
