package data

type vstore struct {
	previous      *vstore
	Next          *vstore
	name          string
	description   string
	contact       string
	qualification int
}

type stores struct {
	Start   *vstore
	Lastest *vstore
	Size    int
}

//Crear Lista de tiendas
func NewList() *stores {
	return &stores{nil, nil, 0}
}

//Insertar nueva tienda
func (stores *stores) setStore(name string, description string, contact string, qualification int) {
	newStore := &vstore{nil, nil, name, description, contact, qualification}

	if stores.Start == nil {
		stores.Start = newStore
		stores.Lastest = newStore
	} else {
		stores.Lastest.Next = newStore
		newStore.previous = stores.Lastest
		stores.Lastest = newStore
	}
	stores.Size++
}

//Imprime las tiendas
func (stores stores) printStores() {

}
