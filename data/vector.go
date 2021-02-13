package data

import (
	"fmt"
	"strings"
)

func VectorMain() {
	fmt.Println("vector")

	vector := NewVector()
	matrix := NewMatrix()

	alldepartments := matrix.SetDataMatrix(dataa())

	vector.setVector(matrix, *alldepartments)
	fmt.Println(vector)
}

type vector struct {
	vector []nodeVector
}

//ID = ASCII( A-Z + int(calificacion) )
type nodeVector struct {
	ID     string
	Stores *stores
}

func NewVector() *vector {
	return &vector{}
}

func NewnodeVector() *nodeVector {
	return &nodeVector{}
}

func (node *nodeVector) setnodeVector(id string, str *stores) {
	node.ID = id
	node.Stores = str
}

//Se recorre la lista de DEPARTAMENTOS EIXISTENTES con DEPARTAMENTES DEL ARCHIVO
func (v *vector) setVector(matrix *AuxMatrix, alldepartments allDepartments) {
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
func (vector *vector) addToVector(matrixx *AuxMatrix, index *int, i int) {
	matrix := matrixx.Matrix
	n0 := nodeVector{matrix[*index].Vector[0].ID, matrix[*index].Vector[0].Stores}
	n1 := nodeVector{matrix[*index].Vector[1].ID, matrix[*index].Vector[1].Stores}
	n2 := nodeVector{matrix[*index].Vector[2].ID, matrix[*index].Vector[2].Stores}
	n3 := nodeVector{matrix[*index].Vector[3].ID, matrix[*index].Vector[3].Stores}
	n4 := nodeVector{matrix[*index].Vector[4].ID, matrix[*index].Vector[4].Stores}

	if i != 0 && *index != len(matrix) {
		matrix = append(matrix[:*index-1], matrix[*index:]...)
	} else {
		matrix = append(matrix[*index:])
	}

	vector.vector = append(vector.vector, []nodeVector{n0, n1, n2, n3, n4}...)
}
