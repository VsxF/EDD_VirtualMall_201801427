package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type data struct {
	Data []departmentsMatriz `json:"Datos"`
}

type departmentsMatriz struct {
	Index      string             `json:"Indice"`
	Department []departmentMatriz `json:"Departamentos"`
}

type departmentMatriz struct {
	Name  string        `json:"Nombre"`
	Store []storeMatriz `json:"Tiendas"`
}

type storeMatriz struct {
	Name    string `json:"Nombre"`
	Desc    string `json:"Descripcion"`
	Contact string `json:"Contacto"`
	Qualifi int    `json:"Calificacion"`
}

//Lee un archivo local y devuelve un struct data con la informacion
func dataa() data {

	jsonFile, err := os.Open("./data/categorias.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var dt data

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err2 := json.Unmarshal([]byte(byteValue), &dt)

	if err2 != nil {

		fmt.Println("error:", err)
	}
	//jsson, _ := json.Marshal(dt)
	fmt.Println(">>matrix loaded<<")

	dt.byQualification()
	return dt
}

//Ordena las tiendas por calificacion
func (dt *data) byQualification() {
	for i := 0; i < len(dt.Data); i++ {
		for j := 0; j < len(dt.Data[i].Department); j++ {

			sort.SliceStable(dt.Data[i].Department[j].Store, func(k, z int) bool {
				return dt.Data[i].Department[j].Store[k].Qualifi < dt.Data[i].Department[j].Store[z].Qualifi
			})
		}
	}
}
