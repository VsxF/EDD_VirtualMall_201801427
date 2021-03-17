package data

import (
	"strings"
	"../products"
)

type Vstore struct {
	Previous      *Vstore
	Next          *Vstore
	Name          string `json:"Nombre"`
	Description   string `json:"Descripcion"`
	Contact       string `json:"Contacto"`
	Qualification int `json:"Calificacion"`
	Department    string `json:"Departamento"`
	Logo string `json:"Logo"`
	Products *products.Tree `json:"Productos"`
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
func (stores *Stores) setStore(name string, description string, contact string, qualification int, dep string, logo string) {
	newStore := &Vstore{nil, nil, name, description, contact, qualification, dep, logo, products.NewTree()}

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

func (stores *Stores) AddStores(st *Stores) {
	if stores.Size > 0 {
		if st.Size > 0 {
			st.Start.Previous = stores.Lastest
			*stores.Lastest.Next = *st.Start
			
			*stores.Lastest = *st.Lastest
		}
	} else {
		*stores = *st
	}
}

//Eliminar tienda de la lista doble de una posicion
func (vt *Vector) DeleteStore(delete Vstore, index int) bool {
	if vt.Vector[index].Stores.Size > 0 {
		store := vt.Vector[index].Stores.Start
		
		for store != nil {
			name := strings.ToLower(store.Name)
			delname := strings.ToLower(delete.Name)
			
			if name == delname && delete.Qualification == store.Qualification {
				
				if store.Next == nil && store.Previous == nil {
					vt.Vector[index].Stores = &Stores{nil, nil, 0}

				} else if store.Next != nil && store.Previous != nil {
					store.Previous.Next = store.Next
					store.Next.Previous = store.Previous
					vt.Vector[index].Stores.Size--

				} else  if store.Next != nil {
					store.Next.Previous = nil
					vt.Vector[index].Stores.Start = store.Next
					vt.Vector[index].Stores.Size--

				} else if store.Previous != nil {
					store.Previous.Next = nil
					vt.Vector[index].Stores.Lastest = store.Previous
					vt.Vector[index].Stores.Size--
				} 
				return true
			} 
			store = store.Next
		}
	}
	return false
}
