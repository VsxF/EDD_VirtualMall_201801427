package reports

import "fmt"

type SearchedStore struct {
	Department string `json:"Departamento"`
	Name string `json:"Nombre"`
	Qualification int `"json:"Calificacion"`
}

func GetSearch(store SearchedStore) {
	fmt.Println("search")
	//fmt.Println(store)
}