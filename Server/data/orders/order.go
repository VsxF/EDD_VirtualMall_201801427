package orders

import (
	"fmt"
	"strconv"
)
//
//"../cart"


type ProductsCodes struct {
	Code int `json:"Codigo"`
}

type Order struct {
	ID int //id de la compra
	Date string `json:"Fecha"`
	Store string `json:"Tienda"`
	Department string `json:"Departamento"`
	Qualification int	`json:"Calificacion"`
	Products []ProductsCodes `json:"Productos"`

	X int
	Y int
	NextX *Order
	PreviousX *Order

	NextY *Order
	PreviousY *Order
}


type InnerList struct {
	First *Order
}

type Header struct {
	Order int // dept.Name -> ascii -> eje Y
	Department string //Nul en el eje X
	Next *Header
	Previous *Header
	InnerList *InnerList
}

type HeaderList struct {
	First *Header
	Last *Header
}

type Orders struct {
	Layer int
	HeaderX *HeaderList
	HeaderY *HeaderList
}

func NewOrders() *Orders{
	headerX := NewHeader()
	headerY := NewHeader()
	return &Orders{0, headerX, headerY}
}

func NewHeader() *HeaderList {
	return &HeaderList{nil, nil}
}

func NewInnerList() *InnerList {
	return &InnerList{nil}
}

func (m *HeaderList) buscar(pos int) *Header {
	aux := m.First

	for aux != nil {
		if aux.Order == pos {
			return aux
		}
		aux = aux.Next
	}

	return nil
}

var ID = 0

func InsertarOrden(m *Orders, date, store, department string, qualification int, products []ProductsCodes) {
	aux := string([]byte(date)[:2])
	aux2 := []byte(department)

	posx, _ := strconv.Atoi(aux)
	posy := 0
	
	for i := 0; i < len(aux2); i++ {
		posy += int(aux2[i])
	}
	
	nuevo := &Order{ID, date, store, department, qualification, products, posx, posy,nil,nil,nil,nil}
	ID++

	cabecerax := m.HeaderX.buscar(posx) 
	if cabecerax == nil {
		lista := NewInnerList()
		cabecerax = &Header{posx,department, nil,nil,lista}
		m.HeaderX.Insertarcabecera(cabecerax)
	}

	cabeceray := m.HeaderY.buscar(posy)
	if (cabeceray == nil){
		listay := NewInnerList()
		cabeceray =  &Header{posx, department,nil,nil,listay}
		m.HeaderY.Insertarcabecera(cabeceray)
	}

	listaX := cabecerax.InnerList
	
	
	//**** SE DEBE INSERTAR ORDENADO en X comparando los valores en y
	if (listaX.First == nil){ //si es nulo solo se agrega
		listaX.First = nuevo
	}else{
		if (nuevo.Y < listaX.First.Y){ //el nodo debe ir al inicio de la lista
			nuevo.NextX = listaX.First
			listaX.First.PreviousX = nuevo
			listaX.First = nuevo
			return
		}else { //se recorre la lista para insertar ordenado
			pivote := listaX.First

			for pivote != nil {
				if (nuevo.Y < pivote.Y){
					nuevo.NextX = pivote
					nuevo.PreviousX = pivote.PreviousX
					pivote.PreviousX.NextX = nuevo
				}else if (nuevo.Y == pivote.Y && nuevo.X == pivote.X ){//comparacion para saber si no se ha insertado una mis posicion
					fmt.Println("Ya existe un nodo es estas coordenadas")
					return
				} else{ //else el y del nuevo es mayor al del pivote 
					if (pivote.NextX == nil){ //se valida si se llego al ultimo 
						pivote.NextX = nuevo //si el siginete es nil 
						nuevo.PreviousX = pivote
						return
					}else{
						pivote = pivote.NextX //si no es el ultimo nos pasamos al siguiente y vuelve a iterar el ciclo
					}
				}
			}
		}
	}

	//Insertar el nodo a las cabeceras
	//Insertar en Y
	listaY := cabeceray.InnerList

	//**** SE DEBE INSERTAR ORDENADO en Y comparando los valores en X
	if (listaY.First == nil){ //si es nulo solo se agrega
		listaY.First = nuevo
	}else{
		fmt.Println("entro aqui ****")
		if (nuevo.X < listaY.First.X){ //el nodo debe ir al inicio de la lista
			fmt.Println("entro aqui *")
			nuevo.NextY = listaY.First
			listaY.First.PreviousY = nuevo
			listaY.First = nuevo
			return
		}else { //se recorre la lista para insertar ordenado
			pivote := listaY.First

			for pivote != nil {
				fmt.Println("entro aqui")
				if (nuevo.X < pivote.X){
					nuevo.NextY = pivote
					nuevo.PreviousY = pivote.PreviousY
					pivote.PreviousY.NextY = nuevo
				}else if (nuevo.Y == pivote.Y && nuevo.X == pivote.X ){//comparacion para saber si no se ha insertado una mis posicion
					fmt.Println("Ya existe un nodo es estas coordenadas")
					return
				} else{ //else el y del nuevo es mayor al del pivote 
					if (pivote.NextY == nil){ //se valida si se llego al ultimo 
						pivote.NextY = nuevo //si el siginete es nil 
						nuevo.PreviousY = pivote
						return
					}else{
						pivote = pivote.NextY //si no es el ultimo nos pasamos al siguiente y vuelve a iterar el ciclo
					}
				}
			}
		}
	}	
}

func (m *HeaderList) Insertarcabecera(nuevo *Header) {
	if m.First == nil {
		m.First = nuevo
		m.Last = nuevo
	}else{
		if m.First == m.Last { //solo hay un dato
			if nuevo.Order > m.First.Order {
				m.First.Next = nuevo
				nuevo.Previous = m.First
				m.Last = nuevo
			}else if nuevo.Order < m.First.Order{
				nuevo.Next = m.First
				m.First.Previous = nuevo
				m.First = nuevo
			}
		}else { //hay mas de un dato
			if nuevo.Order < m.First.Order { //es menor al First 
				nuevo.Next = m.First
				m.First.Previous = nuevo
				m.First = nuevo
			}else if nuevo.Order > m.Last.Order { // es mayor al Last
				m.Last.Next = nuevo
				nuevo.Previous = m.Last
				m.Last = nuevo
			}else {
				aux := m.First

				for aux != nil {
					if nuevo.Order < aux.Order {
						nuevo.Next = aux
						nuevo.Previous = aux.Previous
						aux.Previous.Next = nuevo
						aux.Previous = nuevo
						return
					}
				}
			}
		}
	}
}