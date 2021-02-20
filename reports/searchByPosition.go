package reports

import (

	"strconv"
	"../data"
)

func GetSearchByPosition(position int, vector data.Vector) string {
	if position < len(vector.Vector) {
		if vector.Vector[position].Stores.Size > 0 {
			return getStoresDList(*vector.Vector[position].Stores.Start)
		}
		return "No hay ninguna tienda en esa posicion"
	}
	return "No existe esa posicion"
}

func getStoresDList(store data.Vstore) string {
	result := "["
	for store.Name != "" {
		result += "{\"Nombre\": \"" + store.Name +
		"\",\"Descripcion\":\"" + store.Description +
		"\",\"Contacto\":\"" + store.Contact +
		"\",\"Calificacion\":" + strconv.Itoa(store.Qualification) +
		   "}"

		if store.Next != nil {
			result += ","
			store = *store.Next
		} else {
			break
		}	
	}
	result += "]"
	return result
}
