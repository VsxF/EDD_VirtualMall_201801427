package data

import (
	"fmt"
	"strings"
)

func VectorMain() {
	fmt.Println("vector")
	//dataa()
	v1 := newVectorIndexed()

	v1.setDataVector(dataa())
	//fmt.Println(v1)
	//fmt.Println(v1.vector[0].Stores)
	fmt.Println(v1.vector[3].Stores.Start.Name)
	fmt.Println(v1.vector[4].Stores.Lastest.Name)
	fmt.Println(v1.vector[4].Stores.Start.Name)

}

type vector struct {
	vector []dataVector
}

//ID = ASCII( A-Z + int(calificacion) )
type dataVector struct {
	ID     string
	Stores *stores
}

//Inicializa el vector con todas las posiciones estaticas
//asciiID -> Ej. A1 -> {65, 49} -> inicio
func newVectorIndexed() *vector {
	var vect vector
	asciiID := []byte{65, 49}

	for i := 0; i < 26; i++ {
		for j := 0; j < 5; j++ {
			storeData := dataVector{string(asciiID), NewStoresList()}
			vect.vector = append(vect.vector, storeData)
			asciiID[1] = asciiID[1] + 1
		}
		asciiID[0] = asciiID[0] + 1
		asciiID[1] = 49
	}
	return &vect
}

//Actualiza las tiendas en base al id
func (dtvt *dataVector) updateStore(id string, sts stores) {
	for i := 0; i < dtvt.Stores.Size; i++ {
		if dtvt.ID == id {
			dtvt.Stores = &sts
			break
		} else {
			fmt.Println("El id esta mal")
		}
	}
}

//Insertar informacion en el vector
func (vt *vector) setDataVector(data data) {
	for i := 0; i < len(data.Data); i++ {
		vt.mapDepartments(data.Data[i].Index, data.Data[i].Department)
	}
}

//obtiene los departamentos departamanetos
func (vt *vector) mapDepartments(id string, department []departmentMatriz) {
	for i := 0; i < len(department); i++ {
		vt.setStores(id, department[i].Name, department[i].Store)
	}
}

// //Crea e ingresa las tiendas al vector
func (vt *vector) setStores(idVector string, dept string, storeinfo []storeMatriz) {

	for i := 0; i < len(storeinfo); i++ {

		for j := 0; j < len(vt.vector); j++ {
			if strings.Contains(vt.vector[j].ID, idVector) {

				switch storeinfo[i].Qualifi {
				case 2:
					j++
				case 3:
					j = +2
				case 4:
					j = +3
				case 5:
					j = +4
				}
				//mistake
				vt.vector[j].Stores.setStore(storeinfo[i].Name, storeinfo[i].Desc, storeinfo[i].Contact, storeinfo[i].Qualifi, dept)
				break
			} else {
				j = +5
			}

		}
	}
}
