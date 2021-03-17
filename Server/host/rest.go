package host

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../data/products"
	data "../data/stores"
	"../reports"
	"github.com/go-extras/tahwil"
	"github.com/gorilla/mux"
)

var MainVector *data.Vector
var MainInventory *products.InventorysJSON

func Request() {
	fmt.Println("Listening And Serving ...")

	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/cargartienda", setStores).Methods("POST", "OPTIONS")
	myrouter.HandleFunc("/getTiendas", getStores).Methods("GET", "OPTIONS")
	myrouter.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	myrouter.HandleFunc("/TiendaEspecifica", searchByName).Methods("POST")
	myrouter.HandleFunc("/id/{id}", searchByPosition).Methods("GET")
	myrouter.HandleFunc("/Eliminar", deleteStore).Methods("POST")
	myrouter.HandleFunc("/setInventarios", setInventarios).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", (myrouter)))
}

//Ingrsa las tiendas POST
func setStores(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	var response data.Data
	err := json.Unmarshal(body, &response)
	Error(err)
	if err == nil {
		MainVector = data.NewVector()
		MainVector.GetVector(response)
		reports.SaveVector(*MainVector)
		fmt.Fprintf(w, "Seted")
		fmt.Println("Seted")
	}
}

func getStores(w http.ResponseWriter, r *http.Request) {
	v, _ := tahwil.ToValue(MainVector)
	res, _ := json.Marshal(v) 

	fmt.Fprintf(w, string(res))
}

func getArreglo(w http.ResponseWriter, r *http.Request) {
	reports.GetComplete(MainVector)
	fmt.Println("Archivo creado")
	fmt.Fprintf(w, "Archivo creado")
}

func searchByName(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	response := data.NewVstore()
	err := json.Unmarshal(body, &response)
	Error(err)
	if err == nil {
		result := reports.GetSearchByStore(response, MainVector)

		fmt.Fprintf(w, result)
		fmt.Println(result)
	}
}

func searchByPosition(w http.ResponseWriter, r *http.Request) {
	url := []byte(r.URL.Path)
	idURL := string(url[4:])
	id, err := strconv.Atoi(idURL)
	Error(err)
	if err == nil {
		result := reports.GetSearchByPosition(id, *MainVector)

		fmt.Println(result)
		fmt.Fprintf(w, result)
	}
}

func deleteStore(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	auxBody := strings.ReplaceAll(string(body), "\"Categoria\":", "\"Departamento\":")

	response := data.NewVstore()
	err := json.Unmarshal([]byte(auxBody), &response)
	Error(err)

	if err == nil {
		result := reports.DeleteStore(response, MainVector)
		reports.SaveVector(*MainVector)
		fmt.Fprintf(w, result)
		fmt.Println(result)
	}
}

func setInventarios(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	response := products.NewInventorys()
	err := json.Unmarshal(body, &response)

	if !Error(err) {
		MainInventory = response

		fmt.Println("Seted")
		fmt.Fprintf(w, "Seted")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Error(err error) bool {
	if err != nil {
		fmt.Println("error:", err)
		return true
	}
	return false
}
