package data

import (
	"fmt"
	"strings"
)

func VectorMain() *Vector{
	fmt.Println("vector")

	vector := NewVector()
	matrix := NewMatrix()

	alldepartments := matrix.SetDataMatrix(dataa())

	vector.setVector(matrix, *alldepartments)
	//fmt.Println(vector)
	return vector
}

type Vector struct {
	Vector []NodeVector
}

//ID = ASCII( A-Z + int(calificacion) )
type NodeVector struct {
	ID     string
	Stores *Stores
}

func NewVector() *Vector {
	return &Vector{}
}

func NewnodeVector() *NodeVector {
	return &NodeVector{}
}

func (node *NodeVector) setNodeVector(id string, str *Stores) {
	node.ID = id
	node.Stores = str
}

//Se recorre la lista de DEPARTAMENTOS EIXISTENTES con DEPARTAMENTES DEL ARCHIVO
func (v *Vector) setVector(matrix *AuxMatrix, alldepartments allDepartments) {
	for i := 0; i < len(alldepartments.department); i++ {
		for j := 0; j < len(matrix.Matrix); j++ {
			if strings.Contains(strings.ToLower(matrix.Matrix[j].Department), strings.ToLower(alldepartments.department[i])) {
				v.addToVector(matrix, &j, i)
			}
		}
	}
}

//agrega al vector las 5 tiendas de una categoria y un departamento
//Los agregados, se eliminan del vector de entrada, para mejorar rendimiento
func (vector *Vector) addToVector(matrixx *AuxMatrix, index *int, i int) {
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

if auxIndex == len(matrixx.Matrix) {
	*index = -1
}

	vector.Vector = append(vector.Vector, []NodeVector{n0, n1, n2, n3, n4}...)
}
