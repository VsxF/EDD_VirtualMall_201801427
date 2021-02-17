//Vector de departamentos con sus tiendas
//Estructura intermedia para poder ordenar el vector a row major

package data

import (
	"fmt"
	"strconv"
	"strings"
)

type AuxMatrix struct {
	Matrix []AuxVector
}

type AuxVector struct {
	Department string
	Vector     []NodeVector
}

type allDepartments struct {
	department []string
}

func NewMatrix() *AuxMatrix {
	return &AuxMatrix{}
}

func NewAuxVector() *AuxVector {
	return &AuxVector{"", nil}
}

func (mt *AuxMatrix) addToMatrix(dept string, node []NodeVector) {
	aux := AuxVector{dept, node}
	mt.Matrix = append(mt.Matrix, aux)
}

func (dpt *allDepartments) AddDepartmentAll(dept string) {
	if dept != "" {
		saved := false
		if dpt == nil {
			dpt = &allDepartments{[]string{dept}}
		}
		for i := 0; i < len(dpt.department); i++ {
			if strings.ToLower(dept) == strings.ToLower(dpt.department[i]) {
				saved = true
				break
			}
		}
		if !saved {
			dpt.department = append(dpt.department, dept)
		}		
	}	
}

//Insertar informacion en el vector 
// > Linealizado, tiendas lista doble
// >vector de departamentos existentes (sin repetir)
func (mt *AuxMatrix) SetDataMatrix(data Data, alldpt *allDepartments ) {
	for i := 0; i < len(data.Data); i++ {
		mt.SetDepartmentMatrix(data.Data[i].Index, data.Data[i].Department, alldpt)
	}
}

//Obtiene los departamentos departamanetos
func (mt *AuxMatrix) SetDepartmentMatrix(id string, department []departmentMatriz, alldpt *allDepartments) {
	for i := 0; i < len(department); i++ {
		dpt := department[i]
		alldpt.AddDepartmentAll(dpt.Name)
		mt.SetStoresAux(id, dpt.Name, dpt.Store)
	}
}

//Crea e ingresa las tiendas al vector
func (mt *AuxMatrix) SetStoresAux(idVector string, dept string, storeinfo []storeMatriz) {
	node1 := NewnodeVector()
	node2 := NewnodeVector()
	node3 := NewnodeVector()
	node4 := NewnodeVector()
	node5 := NewnodeVector()
	previousQual := 0
	id := dept + idVector

	node1.ID = id + "1"
	node1.Stores = NewStoresList()
	node2.ID = id + "2"
	node2.Stores = NewStoresList()
	node3.ID = id + "3"
	node3.Stores = NewStoresList()
	node4.ID = id + "4"
	node4.Stores = NewStoresList()
	node5.ID = id + "5"
	node5.Stores = NewStoresList()

	for i := 0; i < len(storeinfo); i++ {
		str := storeinfo[i]

		if previousQual != 0 && previousQual == str.Qualifi {

			switch str.Qualifi {
			case 1:
				node1.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept)
			case 2:
				node2.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept)
			case 3:
				node3.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept)
			case 4:
				node4.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept)
			case 5:
				node5.Stores.setStore(str.Name, str.Desc, str.Contact, str.Qualifi, dept)
			default:
				fmt.Println("error: ubiacion " + idVector + dept + strconv.Itoa(str.Qualifi) + " tienda: " + str.Name)
				fmt.Println("No se encontre calificacion")
			}
		} else {
			i--
		}

		previousQual = str.Qualifi
	}

	node := []NodeVector{*node1, *node2, *node3, *node4, *node5}
	mt.addToMatrix(id, node)
}
