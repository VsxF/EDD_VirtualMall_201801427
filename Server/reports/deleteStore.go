package reports

import (
	"strings"
	"strconv"
	data "../data/stores"
)

func DeleteStore(store *data.Vstore, vector *data.Vector) string {
	auxIndex := []byte(store.Name)
	id := strings.ToLower(store.Department + string(auxIndex[0]))
	deleted := false
	response := "No se encontro la tienda"
	
	for i := 0; i < len(vector.Vector); i++ {
		vectID := strings.ToLower(vector.Vector[i].ID)

		if strings.Contains(vectID, id) {
			auxID := id + strconv.Itoa(store.Qualification)
			
			if auxID == vectID {
				deleted = vector.DeleteStore(*store, i)
				break		
			}
		} else {
			i += 4
		}
	}
	if deleted {
		response = "Se elimino la tineda"
	}
	return response
}