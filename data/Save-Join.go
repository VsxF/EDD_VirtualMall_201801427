package data

import (
	"strings"
	"fmt"
)

//SaveVector -- reports
type MatrixLista struct {
	Index_qual [][]Stores
}

//Crea una nueva matriz, con las posiciones estaticas
func NewMatrixLista(vector Vector) *MatrixLista {
	fmt.Println()
	j := len(vector.Alldepartments.Department)
	i := len(vector.Vector)/ (j*5)
	mat := make([][]Stores, j)
	for k := 0; k < j; k++ {
		mat[k] = make([]Stores, i)		
	}
	return &MatrixLista{mat}
}

//Recorre el vector -- devuelve un vector de ids ROWMAJOR para la matriz
func (mt *MatrixLista) MapVector(vector Vector) []string {
	prevIndex := ""
	prevDept := ""
	indexCount := -1
	deptCount := -1	
	var matrizID []string	

	for i := 0; i < len(vector.Vector); i++ {
			auxIndex := []byte(vector.Vector[i].ID)
			index := strings.ToLower(string(auxIndex[len(auxIndex)-2]))
			dept := strings.ToLower(string(auxIndex[:len(auxIndex)-2]))

			if prevDept != dept {
				i--
					deptCount++
				indexCount = 0
			} else if prevIndex == index {
				mt.addToMatrixLista(indexCount, deptCount, *vector.Vector[i].Stores)
				matrizID = append(matrizID, vector.Vector[i].ID)
			} else {
				i--
					indexCount++				
			}
			prevIndex = index
			prevDept = dept
		}
		return matrizID
}

//Agrega la informacion a la matrixLista
func (mt *MatrixLista) addToMatrixLista(index, dept int, stores Stores) {
	if stores.Size > 0   {
		if mt.Index_qual[dept][index].Size == 0 {
			mt.Index_qual[dept][index] = stores
		} else {

			mt.Index_qual[dept][index].Lastest.Next = stores.Start
			mt.Index_qual[dept][index].Lastest = stores.Lastest
		}
	}
}

//

type auxStruct struct {
	Departamento string
	Stores Stores
}

func Matrix2Json(dataa *Data, matrix MatrixLista, id []string) {
	var deslinealizado [1000]Stores
	prevDpt := ""
	dpt :=""
	for  i, stores := range matrix.Index_qual {
		for j := 0; j < len(stores); j++ {
			if stores[j].Size > 0 {
				k:= j*len(matrix.Index_qual) + i
				dpt = stores[j].Start.Department	
				
				if prevDpt != dpt {
					deslinealizado[k] = stores[j]
				}
			}
		}
		prevDpt = dpt
	}
	deslinealizado2Data(dataa, deslinealizado)
}

func deslinealizado2Data(dataa *Data, deslinealizado [1000]Stores) {
	var allStores []auxStruct
	for i := 0; i < len(deslinealizado); i++ {
		if deslinealizado[i].Size > 0 {
			aux := auxStruct{ deslinealizado[i].Start.Department, deslinealizado[i] }
			allStores = append(allStores, aux)
		}
	}
	setDepts2Index(dataa, allStores)
}

func setDepts2Index(dataa *Data, allStores []auxStruct) {
	dptsMatriz := NewDepartmentsMatriz()
	dptMatriz := NewDepartmentMatriz()
	prevIndex := []byte("^")[0]
	for i := 0; i < len(allStores); i++ {
		index := []byte(strings.ToLower(allStores[i].Stores.Start.Name))[0]
		if prevIndex != index || i == len(allStores)-1{

			if i == len(allStores)-1 {
				
				dptMatriz.Name = allStores[i].Departamento
				dptMatriz.Store = getStoresM2J(*allStores[i].Stores.Start)
				
				
				dptsMatriz.Index = strings.Title(string(prevIndex))
				dptsMatriz.Department = append(dptsMatriz.Department, *dptMatriz)
				
			} 
				dataa.Data = append(dataa.Data, *dptsMatriz)
				dptsMatriz = NewDepartmentsMatriz()		
		}
		dptMatriz.Name = allStores[i].Departamento
		dptMatriz.Store = getStoresM2J(*allStores[i].Stores.Start)
		
		dptsMatriz.Index = strings.Title(string(index))
		dptsMatriz.Department = append(dptsMatriz.Department, *dptMatriz)

		prevIndex = index
	}
	
	if len(dataa.Data) > 0 {
		dataa.Data = dataa.Data[1:]
	}
	
}

func getStoresM2J(store Vstore) []StoreMatriz{
	var response []StoreMatriz
	for store.Name != "" {
		storeMtx := NewStoreMatriz()
		storeMtx.Name = store.Name
		storeMtx.Desc = store.Description
		storeMtx.Contact = store.Contact
		storeMtx.Qualifi = store.Qualification

		response = append(response, *storeMtx)
		if store.Next == nil {
			break
		}
		store = *store.Next
	}
	return response
}
