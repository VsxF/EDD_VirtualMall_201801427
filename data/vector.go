package data

import (
	"fmt"
)

func main() {
	fmt.Println("vector")
}

type vector struct {
	vector []dataVector
}

type dataVector struct {
	ID     string
	Stores *stores
}

func newVector() *vector {
	return &vector{}
}

func (vector *vector) setVectorData(data data) {

}

//Insertar tiendas
func (dataVector *dataVector) setStores(stores stores) {

}
