package data

import (
	"fmt"
	"strings"
)

func (vector *Vector) GetVector(data Data) {
	fmt.Println()
	if data.Data != nil {	
		matrix := NewMatrix()
		data.byQualification()
		matrix.SetDataMatrix(data, vector.Alldepartments)
		//fmt.Println(matrix)
		vector.setVector(matrix, vector.Alldepartments)
	}
}

type Vector struct {
	Vector []NodeVector
	Alldepartments *AllDepartments
}

//ID = ASCII( Dept + A-Z + int(calificacion) )
type NodeVector struct {
	ID     string
	Stores *Stores
}

func NewVector() *Vector {
	return &Vector{[]NodeVector{}, &AllDepartments{[]string{}}}
}

func NewnodeVector() *NodeVector {
	return &NodeVector{}
}

func (node *NodeVector) setNodeVector(id string, str *Stores) {
	node.ID = id
	node.Stores = str
}

//Se recorre la lista de DEPARTAMENTOS EIXISTENTES con DEPARTAMENTES DEL ARCHIVO
func (v *Vector) setVector(matrix *AuxMatrix, alldepartments *AllDepartments) {
	for i := 0; i < len(alldepartments.Department); i++ {
		_allDepartment := strings.ToLower(alldepartments.Department[i])
		
		for j := 0; j < len(matrix.Matrix); j++ {
			matrixDepartment := strings.ToLower(matrix.Matrix[j].Department)
			
			if strings.Contains(matrixDepartment, _allDepartment) {
				v.addToVector(matrix, &j)		
			}
		}
	}
}

//agrega al vector las 5 tiendas de una categoria y un departamento
//Los agregados, se eliminan del vector de entrada, para mejorar rendimiento
func (vector *Vector) addToVector(matrixx *AuxMatrix, index *int) {
	matrix := matrixx.Matrix
	n0 := NodeVector{matrix[*index].Vector[0].ID, matrix[*index].Vector[0].Stores}
	n1 := NodeVector{matrix[*index].Vector[1].ID, matrix[*index].Vector[1].Stores}
	n2 := NodeVector{matrix[*index].Vector[2].ID, matrix[*index].Vector[2].Stores}
	n3 := NodeVector{matrix[*index].Vector[3].ID, matrix[*index].Vector[3].Stores}
	n4 := NodeVector{matrix[*index].Vector[4].ID, matrix[*index].Vector[4].Stores}
	
	auxIndex := *index + 1
	if *index == 0 {
		matrixx.Matrix = matrixx.Matrix[auxIndex:]
	} else if *index != len(matrixx.Matrix) {
		matrixx.Matrix = append(matrixx.Matrix[:auxIndex-1], matrixx.Matrix[auxIndex:]...)
	} else {
		matrixx.Matrix = matrixx.Matrix[:auxIndex-1]
	}
	*index = -1

	vector.Vector = append(vector.Vector, []NodeVector{n0, n1, n2, n3, n4}...)
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