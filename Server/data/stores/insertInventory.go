package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	p "../products"
)

func (vector *Vector) SetInventrorys(inventorysJSON *p.InventorysJSON) {
	inv := inventorysJSON.Inventorys
	for i := 0; i < len(inv); i++ {
		if vector.existDepartment(inv[i].Department) {
			vector.mapDepartments(&inv[i])
		} else {
			fmt.Println("El departamento no existe >>> " + inv[i].Department)
		}
	}
}

func (vector *Vector) existDepartment(department string) bool {
	dpt := vector.Alldepartments.Department
	for i := 0; i < len(dpt); i++ {
		deptt, depart := strings.ToLower(dpt[i]), strings.ToLower(department)
		if strings.Contains(deptt, depart) {
			return true
		}
	}
	return false
}

func (vector *Vector) mapDepartments(storeInfo *p.StoreJSON) {
	department := low(storeInfo.Department + string([]byte(storeInfo.Name)[0]))
	qualif := strconv.Itoa(storeInfo.Qualif)

	for i := 0; i < len(vector.Vector); i++ {
		aux := low(vector.Vector[i].ID)
		vdepartment := string([]byte(aux)[:len(aux)-1])
		vqualif := string([]byte(aux)[len(aux)-1])

		if department == vdepartment {
			if qualif == vqualif {
				//Diferent qualification??????????
				vector.mapStores(storeInfo, i)
			}
		} else if i%4 == 0 {
			i += 5
		}
	}
}

func (vector *Vector) mapStores(storeInfo *p.StoreJSON, i int) {
	stores := vector.Vector[i].Stores
	next, prev := stores.Start, stores.Lastest
	storeName := low(storeInfo.Name)

	for i := 0; i < (stores.Size+1)/2; i++ {
		if low(next.Name) == storeName {
			next.Products = addToTree(*next.Products, storeInfo.Products, i)
			break
		} else if low(prev.Name) == storeName {
			prev.Products = addToTree(*prev.Products, storeInfo.Products, i)
			break
		}
		next = next.Next
		prev = prev.Previous
	}
}

/// Agregar el comparar productos
func addToTree(tree p.Tree, products []p.ProductJSON, i int) *p.Tree {
	for i := 0; i < len(products); i++ {
		pt := products[i]

		price := int(pt.Price * 100)
		tree.InsertProduct(pt.Name, pt.Code, pt.Desc, price, pt.Quant, pt.Image)
	}
	return &tree
}

func low(st string) string {
	return strings.ToLower(st)
}

/// JUST FOR Develop
func (vector *Vector) AuxSetInventrorys() {
	jsonFile, _ := os.Open("./Inventarios.json")
	defer jsonFile.Close()

	inventoryJSON := p.NewInventorys()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err2 := json.Unmarshal([]byte(byteValue), &inventoryJSON)

	if err2 != nil {
		fmt.Println("error: ", err2)
	}

	vector.SetInventrorys(inventoryJSON)
}
