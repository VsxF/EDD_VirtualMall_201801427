package reports

import (
	"strings"
	"strconv"
	"../data"
)

func GetSearchByStore(reqstore *data.Vstore, vector *data.Vector) string {
	result := getSearch(reqstore, vector)

	stringResult := "No se encontro ningun resultado"
	
	if result.Name != "" {
		stringResult = "{\"Nombre\":\"" + result.Name + 
		"\",\"Descripcion\":\"" + result.Description + 
		"\",\"Contacto\":\"" + result.Contact + 
		"\",\"Calificacion\":" + strconv.Itoa(result.Qualification) + 
			"}"
	}
	return stringResult
}

func getSearch(reqstore *data.Vstore, vector *data.Vector) data.Vstore {
	auxIndex := []byte(reqstore.Name)
	index := string(auxIndex[0]) + strconv.Itoa(reqstore.Qualification)
	vect := vector.Vector
	var result data.Vstore

	for i := 0; i < len(vect); i++ {		
		if strings.Contains(vect[i].ID, index){
			result = searchStore(reqstore.Name, *vect[i].Stores)
			
			if result.Name != "" {
				break
			}
		}
	}

	return result
}

func searchStore(name string, stores data.Stores) data.Vstore{
	start := stores.Start
	last := stores.Lastest
	var auxVstore data.Vstore
	if start != nil {
		if strings.ToLower(name) == strings.ToLower(start.Name) {
			auxVstore = *start
		} else if strings.ToLower(name) == strings.ToLower(last.Name){
			auxVstore = *last
		} else {
			start = start.Next 
			last = last.Previous
		}
	}
	return auxVstore
}