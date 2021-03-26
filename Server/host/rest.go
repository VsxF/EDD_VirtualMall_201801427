package host

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	cart "../data/cart"
	products "../data/products"
	orders "../data/orders"
	data "../data/stores"
	//"../products"

	"../reports"
	"github.com/go-extras/tahwil"
	"github.com/gorilla/mux"
)

var MainVector = data.NewVector()
var MainCart = cart.NewCart()
var MainOrders = orders.NewCalendar()
var OrderID = 0

func Request() {
	fmt.Println("Listening And Serving ...")

	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/cargartienda", setStores).Methods("POST", "OPTIONS")
	myrouter.HandleFunc("/getTiendas", getStores).Methods("GET", "OPTIONS")
	myrouter.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	myrouter.HandleFunc("/TiendaEspecifica", searchByName).Methods("POST")
	myrouter.HandleFunc("/id/{id}", searchByPosition).Methods("GET")
	myrouter.HandleFunc("/Eliminar", deleteStore).Methods("POST")
	myrouter.HandleFunc("/setInventorys", setInventorys).Methods("POST", "OPTIONS")

	myrouter.HandleFunc("/add2Cart", add2Cart).Methods("POST", "OPTIONS")
	myrouter.HandleFunc("/delCart/{code}", delCart).Methods("DELETE", "OPTIONS")
	myrouter.HandleFunc("/buyCart", buyCart).Methods("POST", "OPTIONS")
	myrouter.HandleFunc("/getCart", getCart).Methods("GET", "OPTIONS")

	myrouter.HandleFunc("/setOrders", setOrders).Methods("POST", "OPTIONS")
	myrouter.HandleFunc("/getOrders", getOrders).Methods("GET", "OPTIONS")

	myrouter.HandleFunc("/getMatrixGraph/{date}", getMatrixGraph).Methods("GET", "OPTIONS")
	myrouter.HandleFunc("/getYearsGraph", getYearsGraph).Methods("GET", "OPTIONS")
	myrouter.HandleFunc("/getDays/{date}", getDays).Methods("GET", "OPTIONS")
	myrouter.HandleFunc("/getMonthsGraph", getMonthsGraph).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":3000", (myrouter)))
}

func getMatrixGraph(w http.ResponseWriter, r *http.Request) {
	if count > 0 {
	url := []byte(r.URL.Path)
	date := string(url[16:])
	image := reports.GetMatrixGraph(MainOrders, date)

	w.Write(image)
	count = 0
	}
	count++
}

func getYearsGraph(w http.ResponseWriter, r *http.Request) {
	if count > 0 {
	image := reports.GetYearsGraph(MainOrders)
	w.Write(image)
	count = 0
	}
	count++
}

func getMonthsGraph(w http.ResponseWriter, r *http.Request) {
	if count > 0 {
	image := reports.GetMonthsGraph()
	w.Write(image)
	count = 0
	}
	count++
}

func getDays(w http.ResponseWriter, r *http.Request) {
	if count > 0 {
		url := []byte(r.URL.Path)
		date := string(url[9:])
		a := reports.GetDayOrders(MainOrders, date)

		ajson, _ := json.Marshal(a)
		w.Write(ajson)
	count = 0
	}
	count++
}



//Ingrsa las tiendas POST
func setStores(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	body, _ := ioutil.ReadAll(r.Body)
	
	var response data.Data
	err := json.Unmarshal(body, &response)
	Error(err)
	if err == nil {
		MainVector = data.NewVector()
		MainVector.GetVector(response)
		
		if len(MainVector.Vector) > 0 {
		// reports.SaveVector(*MainVector)
			fmt.Fprintf(w, "Seted")
			fmt.Println("Seted")
		}

		
	}
}

func getStores(w http.ResponseWriter, r *http.Request) {
	v, _ := tahwil.ToValue(MainVector)
	res, _ := json.Marshal(v) 

	w.Write(res)
	//json.NewEncoder(w).Encode(res)
	//fmt.Fprintf(w, string(res))
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

func setInventorys(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	response := products.NewInventorys()
	err := json.Unmarshal(body, &response)

	if !Error(err) {
		if len(response.Inventorys) > 0 &&
			 len(MainVector.Vector) > 0 {

			MainVector.SetInventrorys(response)
		
			fmt.Println("Seted")
			fmt.Fprintf(w, "Seted")
		}	
	}
}

func add2Cart(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var response products.Product
	json.Unmarshal(body, &response)
	
	if response.Name != "" {
		MainCart.InsertProduct(&response)
	}
	
	fmt.Fprintf(w, "Seted")
	fmt.Println("Seted")
}

var countDel = 0
func delCart(w http.ResponseWriter, r *http.Request) {
	url := []byte(r.URL.Path)
	codeURL := string(url[9:])
	code,_ := strconv.Atoi(codeURL)
	

	if countDel > 0 {
		countDel = 0
		MainCart.DeleteProduct(code)
		getCart(w, r)
	}

	countDel++
	
}

var count = 0
//Falta retornar el estado true, de vendido
func buyCart(w http.ResponseWriter, r *http.Request) {
	if count == 1 {
		MainVector.UpdateQuant(MainCart)
		MainCart = cart.NewCart()
		count = 0
	} else {
		count++
	}
}

func getCart(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(MainCart) 
	Error(err)
	w.Write(res)
}


func setOrders(w http.ResponseWriter, r *http.Request){
	if count > 0 {
		body, _ := ioutil.ReadAll(r.Body)
		var response orders.AuxOrders
		err := json.Unmarshal(body, &response)
		if !Error(err) {
			if len(MainVector.Vector) > 0 {
				MainOrders.SetOrders(response)
				fmt.Println(MainOrders)
				if MainOrders.Size > 0 {
					fmt.Println("Seted")
					fmt.Fprintf(w, "Seted")
				}	
			}	
		}
		count = 0
	}
	count++
}

func getOrders(w http.ResponseWriter, r *http.Request){
	v, _ := tahwil.ToValue(MainOrders)
	res, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(res)
	w.Write(res)
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
