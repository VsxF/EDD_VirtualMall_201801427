package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func DataA() {

	jsonFile, err := os.Open("./data/categorias.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	var data data

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err2 := json.Unmarshal([]byte(byteValue), &data)

	if err2 != nil {

		fmt.Println("error:", err)
	}

	// fmt.Println(">>")
	// fmt.Println(data.Data)
	// data.Data[0].Index = "ss"
	// fmt.Println(data.Data)
	fmt.Println(">>matriz.go")

}

type data struct {
	Data []departments `json:"Datos"`
}

type departments struct {
	Index      string       `json:"Indice"`
	Department []department `json:"Departamentos"`
}

type department struct {
	Name  string  `json:"Nombre"`
	Store []store `json:"Tiendas"`
}

type store struct {
	Name    string `json:"Nombre"`
	Desc    string `json:"Descripcion"`
	Contact string `json:"Contacto"`
	Qualifi int    `json:"Calificacion"`
}
