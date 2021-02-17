package host

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"	
	"../data"
	"../reports"
)

var MainVector *data.Vector

func Request() {
	fmt.Println("Listening And Serving ...")
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/cargartienda", setStores).Methods("POST")
	myrouter.HandleFunc("/getStores", getStores).Methods("GET")
	myrouter.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	myrouter.HandleFunc("/TiendaEspecifica", searchByName).Methods("POST")
	myrouter.HandleFunc("/{Departamento:{Departamento}}", searchByName).Methods("POST")
	myrouter.HandleFunc("/", searchByPosition).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", myrouter))
}

//Ingrsa las tiendas POST
func setStores(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Set")
	body, _ := ioutil.ReadAll(r.Body)
	var response data.Data
	err := json.Unmarshal(body, &response)

	Error(err)

	// reports.CreateFile(reports.File{"categorias", string(body[:]), ".json"})

	MainVector.GetVector(response)
}

//Imprimir tiendas en consola desde un GET
func getStores(w http.ResponseWriter, r *http.Request) {
	
	//fmt.Println(r.ReadResponse())
	fmt.Fprintf(w, "El vector se imprimio en consola")
	fmt.Println(MainVector)
}

//Obtener imagen del vector GET
func getArreglo(w http.ResponseWriter, r *http.Request) {
	reports.GetComplete(MainVector)
	fmt.Println("Archivo creado")
	fmt.Fprintf(w, "Archivo creado")
}

func searchByName(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	
	//var response reports.SearchedStore
	
	//err := json.Unmarshal(body, &response)

	//Error(err)

	//reports.GetSearch(response)
	fmt.Println("ss")
	fmt.Println(r.URL.Query())
	//fmt.Fprintf(w, string(r.URL.Query()))
}

func searchByPosition(w http.ResponseWriter, r *http.Request) {

}

func Error(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}