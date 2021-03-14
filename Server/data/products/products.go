package products

import "fmt"

type Inventarios struct {
	Inventarios []Store `json:"Inventarios"`
}

type Store struct {
	Store string `json:"Tienda"`
	Department string `json:"Departamento"`
	Qualif int `json:"Calificacion"`
	Products []Product `json:"Productos"`
}

type Product struct {
	Name string `json:"Nombre"`
	Code int `json:"Codigo"`
	Desc string `json:"Descripcion"`
	Price float64 `json:"Precio"`
	Quant int `json:"Cantidad"`
	Image string `json:"Imagen"`
}

func NewInventarios() *Inventarios {
	return &Inventarios{}
}

func NewStore() *Store {
	return &Store{}
}

func NewProduct() *Product {
	return &Product{}
}

func ma() {
	fmt.Println("prodcuts")
}