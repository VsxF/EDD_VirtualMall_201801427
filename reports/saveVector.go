package reports

import (
	"strings"
	"encoding/json"
	"../data"
)

func SaveVector(vector data.Vector) {
	if len(vector.Vector) > 0 {
		matrix := NewMatrixLista(vector)
		ids := matrix.mapVector(vector)
		
		jsondata := data.NewData() 
		Matrix2Json(jsondata, *matrix, ids)
		js, _ := json.Marshal(jsondata)

		file := NewFile("categorias", ".json")
		file.AddText(string(js))
		CreateFile(*file)
	}
}
//Matrix Lista dobel [A1,A2][Departments]

type MatrixLista struct {
	Index_qual [][]data.Stores
}

//Crea una nueva matriz, con las posiciones estaticas
func NewMatrixLista(vector data.Vector) *MatrixLista {
	
	j := len(vector.Alldepartments.Department)
	i := len(vector.Vector)/ (j*5)
	mat := make([][]data.Stores, j)
	for k := 0; k < j; k++ {
		mat[k] = make([]data.Stores, i)		
	}
	return &MatrixLista{mat}
}

//Recorre el vector -- devuelve un vector de ids ROWMAJOR para la matriz
func (mt *MatrixLista) mapVector(vector data.Vector) []string {
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
func (mt *MatrixLista) addToMatrixLista(index, dept int, stores data.Stores) {
	if stores.Size > 0 {	
		if mt.Index_qual[dept][index].Size == 0 {
			mt.Index_qual[dept][index] = stores
		} else {
			mt.Index_qual[dept][index].Lastest.Next = stores.Start
			mt.Index_qual[dept][index].Lastest = stores.Lastest
		}
	}
}

//
func Matrix2Json(dataa *data.Data, matrix MatrixLista, id []string) {
	prevIndex := ""
	var _departments []data.DepartmentMatriz
	departments := data.NewDepartmentsMatriz()

	for  _, stores := range matrix.Index_qual {
		for j := 0; j < len(stores); j++ {
			store := stores[j].Start
			if store != nil {
				index, dptJSON := setDepartmentsM2J(dataa, store, &prevIndex, &j)
				
				if index != "" && dptJSON.Name != "" {
					_departments = append(_departments, *dptJSON)
					departments.Index = index
				}

				if index == "!5!A" || j == len(stores)-1 {
					departments.Department = _departments
					dataa.Data = append(dataa.Data, *departments)
					_departments = nil
				}								
			}
		}
		
	}	
}

func setDepartmentsM2J(dataa *data.Data, store *data.Vstore, prevIndex *string, j *int) (string, *data.DepartmentMatriz ) {
	_index := []byte(store.Name) 
	index := string(_index[:1])
	dptJSON := data.NewDepartmentMatriz()
	if index == *prevIndex {
		dptJSON.Name = store.Department
		dptJSON.Store = getStoresM2J(*store)
		*prevIndex = index
		return index, dptJSON					
	} else if *prevIndex == "" {
		*j--
	} else {
		*j--
		*prevIndex = index
		return "!5!A", dptJSON
	}
	*prevIndex = index
	return "", dptJSON
}

func getStoresM2J(store data.Vstore) []data.StoreMatriz{
	var response []data.StoreMatriz
	for store.Name != "" {
		storeMtx := data.NewStoreMatriz()
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
