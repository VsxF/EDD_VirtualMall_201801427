package reports

import (
	"fmt"
	"strings"
	"../data"
)

type SearchedStore struct {
	Department string `json:"Departamento"`
	Name string `json:"Nombre"`
	Qualification int `"json:"Calificacion"`
}

func GetSearch(reqstore *data.Vstore, vector *data.Vector ) {
	fmt.Println()
	
	auxIndex := []byte(reqstore.Name)
	index := string(auxIndex[0]) 
	vect := vector.Vector
	for i := 0; i < len(vect); i++ {
		
		if strings.Contains(vect[i].ID, index) {
			fmt.Println(vect[i].ID)
		}

		// start := vector
		// last := vector.Lastest

		// if strings.Contains(start.Department, index) {
		// 	if strings.Contains(start.Name) {
				
		// 	}
		// } else if strings.Contains(last, index) {

		// } else {
		// 	start = start.Next 
		// 	last = last.Previous
		// }

	}
	
}