package reports

import (
	"encoding/json"
	data "../data/stores"
)

func SaveVector(vector data.Vector) {
	if len(vector.Vector) > 0 {
		matrix := data.NewMatrixLista(vector)
		
		ids := matrix.MapVector(vector)
		jsondata := data.NewData() 
		data.Matrix2Json(jsondata, *matrix, ids)
		js, _ := json.Marshal(jsondata)

		file := NewFile("categorias", ".json")
		file.AddText(string(js))
		CreateFile(*file)
	}
}
