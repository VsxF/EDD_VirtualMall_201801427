package reports

import (
	"fmt"
	"encoding/json"
	"../data"
)

func SaveVector(vector data.Vector) {
	fmt.Println()
	if len(vector.Vector) > 0 {
		matrix := data.NewMatrixLista(vector)
		
		ids := matrix.MapVector(vector)
		jsondata := data.NewData() 
		data.Matrix2Json(jsondata, *matrix, ids)
		fmt.Println(jsondata)
		js, _ := json.Marshal(jsondata)

		file := NewFile("categorias", ".json")
		file.AddText(string(js))
		CreateFile(*file)
	}
}
