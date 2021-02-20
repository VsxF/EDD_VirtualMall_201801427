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


//Une y Ordena 2 vectores
// func JoinVectors(vector, second *Vector) *Vector {
// 	response := NewVector()
// 	j := 0
// 	i := 0
// 	response.mapVectors(vector, second, &i, &j)
// 	fmt.Println(response)
// 	return response
// }

// func (response *Vector) mapVectors(vector, second *Vector, i, j *int) {
// 	_, index := getIDnIndex(*vector)
// 	_, index2 := getIDnIndex(*second)

// 	if index < 97 && len(second.Vector) != 0 {
// 		response.addItem(second)
// 		//response.mapVectors(vector, second, i, j)
// 	} else if index2 < 97 && len(vector.Vector) != 0 {
// 		response.addItem(vector)
// 	} else if index == index2 {
// 		a := getDepsIndex(*vector, index2)
// 		response.mergeIndex(vector, second, a)
// 	} else if index < index2 {
// 		response.addItem(vector)
// 	} else if index2 < index {
// 		response.addItem(second)
// 	}

// 	if len(vector.Vector) != 0 || len(second.Vector) != 0 {
// 		response.mapVectors(vector, second, i, j)
// 	}
// }

// func (vector *Vector) addItem(slices *Vector) {
// 	vector.Vector = append(vector.Vector, slices.Vector[0], slices.Vector[1], slices.Vector[2], slices.Vector[3], slices.Vector[4])
// 	slices.Vector = slices.Vector[5:]
// }

// func (response *Vector) mergeIndex(first, second *Vector, deptsFirst []string) {
// 	_, index := getIDnIndex(*first)
// 	id2, index2 := getIDnIndex(*second)
// 	added := false

// 	for i := 0; i < len(deptsFirst); i++ {
// 		if strings.Contains(id2, strings.ToLower(deptsFirst[i])) {
// 			//merge stores
// 			deptsFirst[i] = ""
// 			added = true
// 			break
// 		}
// 	}

// 	if !added {
// 		first.appendDeptIndex(second)
// 	}

// 	if index == index2 {
// 		response.mergeIndex(first, second, deptsFirst)
// 	}
// }

// func (vector *Vector) appendDeptIndex(slices *Vector) {
// 	a := []NodeVector{ slices.Vector[0], slices.Vector[1], slices.Vector[2], slices.Vector[3], slices.Vector[4] }
// 	vector.Vector = append(a, vector.Vector[0:]...)
// 	slices.Vector = slices.Vector[5:]
// }

// func getDepsIndex(vector Vector, index byte) []string {
// 	var resp []string
// 	for i := 0; i < len(vector.Vector); i++ {
// 		id := strings.ToLower(vector.Vector[i].ID)
// 		if strings.Contains(id, string(index)) {
// 			resp = append(resp, id)
// 		} else {
// 			break
// 		}
// 	}
// 	return resp
// }

// func getIDnIndex(vector Vector) (string, byte) {
// 	id := ""
// 	index := []byte("^")[0]
// 	if len(vector.Vector) > 0 {
// 		id = strings.ToLower(vector.Vector[0].ID)
// 		index = []byte(id)[len(id) - 2]
// 	}
// 	return id, index
// }