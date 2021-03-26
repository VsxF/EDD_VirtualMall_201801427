package data

import (
	"fmt"
	"strconv"
	"strings"

	p "../products"
	c "../cart"
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
	Vector         []NodeVector
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
	vector.Vector = append(vector.Vector, []NodeVector{n0, n1, n2, n3, n4}...)
}

//????????????????????????????????????
//UPDATA VECTOR
//actualizar la cantidad de una tienda
//El producto trae en el nodo izq la informacion de la tienda y el dep
var correct = false

func (vector *Vector) UpdateQuant(cart *c.Cart) bool {
	for i := 0; i < len(cart.Products); i++ {
		vector.mapIndexx(cart.Products[i])
	}
	return correct
}

func (vector *Vector) mapIndexx(product *p.Product) {
	for i := 0; i < len(vector.Vector); i++ {
		id := strings.ToLower(vector.Vector[i].ID)
		aux := []byte(product.Left.Name)[0]
		search := strings.ToLower(product.Left.Desc) + strings.ToLower(string(aux))
		
		if strings.Contains(id, search) {
			if id == search + strconv.Itoa(product.Left.Code)   {
				vector.mapStoresV(product, i)
				break
			} 	
		} else if i%4==0 {
			i += 5
		}
	}
}

func (vector *Vector) mapStoresV(product *p.Product, i int) {
	prev := vector.Vector[i].Stores.Lastest
	next := vector.Vector[i].Stores.Start
	
	for j := 0; j < (vector.Vector[i].Stores.Size+1)/2; j++ {
		if prev.Name == product.Left.Name {
			mapProducts(prev.Products, product)
			break
		}
		if next.Name == product.Left.Name {
			mapProducts(next.Products, product)
			break
		}
		prev = prev.Previous
		next = next.Next
	}
}


func mapProducts(products *p.Tree, product *p.Product) {
	inOrder(products.Root, product)
}

func inOrder(products *p.Product, product *p.Product) {
	if products.Left != nil {
		inOrder(products.Left, product)
	}
	
	if products.Name == product.Name {
		if products.Quant - product.Quant > 0 {
			products.Quant = products.Quant - product.Quant
			correct = true
		}
		correct = false
	} else if products.Right != nil {
		inOrder(products.Right, product)
	}
}