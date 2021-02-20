package data

import (
	"strings"
	"fmt"
)

func JoinVectors(first, second Vector) *Vector {
	matrix1 := NewMatrixLista(first)
	matrix2 := NewMatrixLista(second)
	ids1 := matrix1.MapVector(first)
	ids2 := matrix2.MapVector(second)

	data1 := NewData()
	data2 := NewData()

	Matrix2Json(data1, *matrix1, ids1)
	Matrix2Json(data2, *matrix2, ids2)
	fulldata := JoinMatrixs(data1, data2)
	b := NewVector()
	b.GetVector(*fulldata)
	return b
}

//
//Join Matrixs 
func JoinMatrixs(first, second *Data) *Data{
	response := NewData()
	i := 0
	j := 0

	response.mapMatrixs(first, second, &j, &i)

	return response
}

//Mapeo del Vector Data (sobre)
//Setea Data [ { Index, [Departments}], ... ] 
func (mt *Data) mapMatrixs(first, second *Data, j, i *int) {
	index1 := getIndex(first, *i)
	index2 :=  getIndex(second, *j)
	
	if index1 < 97 || index2 < index1{
		mt.Data = append(mt.Data, second.Data[*j])
		*j++
	} else if index2 < 97 || index1 < index2{
		mt.Data = append(mt.Data, first.Data[*i])
		*i++
	} else if index1 == index2 {
		k := 0
		h := 0
		departments := NewDepartmentsMatriz()
		mapSameDepartments(first.Data[*i], second.Data[*j], &k, &h, departments)
		a := string(index1)
		departments.Index = strings.Title(a)
		mt.Data = append(mt.Data, *departments)
		*i++
		*j++
	}
	
	if *i < len(first.Data) || *j < len(second.Data) {
		mt.mapMatrixs(first, second, j, i)
	}
}

//Devuelve Index en ascii minuscula -> A return 97 (a)
func getIndex(vt *Data, k int) byte {
	response := []byte("^")[0]
	if k < len(vt.Data) {
		response = []byte(strings.ToLower(vt.Data[k].Index))[0]
	}
	return response
}

//Setea [Departments] -> { Index[ Name , []StoreMatriz ], ... } -> para unir 2 index iguals
func mapSameDepartments(dpt1, dpt2 DepartmentsMatriz, k, h *int, departments *DepartmentsMatriz)  {
	name1 := strings.ToLower(dpt1.Department[*k].Name)
	name2 := strings.ToLower(dpt2.Department[*h].Name)
	fmt.Println(*k)
	fmt.Println(*h)
	if name1 != name2 { 
		if *k == len(dpt1.Department)-1 {
			departments.Department = append(departments.Department, dpt1.Department[*k])
			*k=999
			if (*h < 999) {
				departments.Department = append(departments.Department, dpt2.Department[*h:]...)
				*h = 999
			}
		}
		if *h == len(dpt2.Department)-1 {
			departments.Department = append(departments.Department, dpt2.Department[*h])
			*h=999
			if (*k < 999) {
				departments.Department = append(departments.Department, dpt1.Department[*k+1:]...)
				*k = 999 
			}
		}	
		if *k != 999 {
			departments.Department = append(departments.Department, dpt1.Department[*k])
			departments.Department = append(departments.Department, dpt2.Department[*h])
			*k++
			*h++
		}
	} else {
		x := 0
		y := 0
		department := NewDepartmentMatriz()
		mapSameDepartment(dpt1.Department[*h], dpt2.Department[*k], &x, &y, department)
		department.Name = strings.Title(name1)
		departments.Department = append(departments.Department, *department)	
		*h++
		*k++	
	}

	if *k < len(dpt1.Department) || *h < len(dpt2.Department) {
		mapSameDepartments(dpt1, dpt2, k, h, departments)
	}
}

// setea []StoreMatriz = [ {st1}, {st2} ]
func mapSameDepartment(store1, store2 DepartmentMatriz, x, y *int, storeMatriz *DepartmentMatriz ) {
	name1 := strings.ToLower(store1.Store[*x].Name)
	name2 := strings.ToLower(store2.Store[*y].Name)

	if name1 != name2 { 
		if *x == len(store1.Store)-1 {
			storeMatriz.Store = append(storeMatriz.Store, store1.Store[*x])
			*x=999
			if (*y < 999) {
				storeMatriz.Store = append(storeMatriz.Store, store2.Store[*y:]...)
				*y = 999
			}
		}
		if *y == len(store2.Store)-1 {
			storeMatriz.Store = append(storeMatriz.Store, store2.Store[*y])
			*y=999
			if (*x < 999) {
				storeMatriz.Store = append(storeMatriz.Store, store1.Store[*x+1:]...)
				*x = 999 
			}
		}	
		if *x != 999 {
			storeMatriz.Store = append(storeMatriz.Store, store1.Store[*x])
			
			storeMatriz.Store = append(storeMatriz.Store, store2.Store[*y])
			fmt.Println(name2)
			*x++
			*y++
		}
	} else {
		
		// x := 0
		// y := 0
		// store := NewStoreMatriz()
		// mapSameDepartment(store1.Store[*y], store2.Store[*x], &x, &y, store)
		// store.Name = strings.Title(name1)
		// storeMatriz.Store = append(storeMatriz.Store, *store)	
		// *y++
		// *x++	
	}

	if *x < len(store1.Store) || *y < len(store2.Store) {
		mapSameDepartment(store1, store2, x, y, storeMatriz)
	}
}