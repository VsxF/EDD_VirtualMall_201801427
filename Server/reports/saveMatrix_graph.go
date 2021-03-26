package reports

import (
	"strconv"
	"os/exec"
	"fmt"
	ord "../data/orders"
)

type auxDays struct {
	x int
	id int
}

type auxDepartments struct {
	y int
	yRow []auxDays
	department string
}

var contentM string
var headersX, headersY []string
var realtionsY, relationsX string

func GetMatrixGraph(orders *ord.Calendar, date string) []byte {
	contentM = ""
	aux := []byte(date)
	month2search, _ := strconv.Atoi(string(aux[3]) + string(aux[4]))
	year2search, _ := strconv.Atoi(string(aux[6:]))

	contentM = "digraph G {\nnode[shape=\"box\"];\nedge[style=\"bold\"];"
	
	searchAVLM(orders.Root, year2search, month2search)

	contentM += "\n}"
	fmt.Println(contentM)
	CreateFile(File{"matrix", contentM, ".dot"})
	exec.Command("neato", "-Tpng", "matrix.dot", "-o", "matrix.png").Output()

	image := image2base64("./matrix.png")
	return []byte(image)
}

func searchAVLM(year *ord.Year, y2s, m2s int) {
	if y2s < year.Year {
		searchAVLM(year.Left, y2s, m2s)
	} else if y2s > year.Year {
		searchAVLM(year.Right, y2s, m2s)
	} else {
		searchMonthM(year.Months, m2s)
	}
}

func searchMonthM(months *ord.Months, m2s int) {
	next := months.Start
	prev := months.Lastest
	monthSearched := next

	for i := 0; i < months.Size; i++ {
		if next.Month == m2s {
			monthSearched = next
			break
		} else if prev.Month == m2s {
			monthSearched = prev
			break
		}
		next = next.Next
		prev = prev.Previous
	}
	//xList := ord.AuxOrders
	var matrix []auxDepartments
	setMatrix(monthSearched.Orders, &matrix)
	contentM += setNodesContent(matrix)
	contentM += setRelations(matrix)
}

func setMatrix(orders *ord.Orders, matrix *[]auxDepartments) {
	headerY := orders.HeaderY.First
	xDays := setXRelations(orders)
	y := 1

	for headerY != nil {
		aux := mapXRelations(xDays, headerY.InnerList.First.ID, headerY.Department, y)
		*matrix = append(*matrix, aux)

		headersY = append(headersY, headerY.Department)
		
		realtionsY += getRelations(headerY.Department, strconv.Itoa(headerY.InnerList.First.ID)) 
		
		headerY = headerY.Next
		y++
	}
}

func setXRelations(orders *ord.Orders) []auxDays {
	days := []auxDays{}
	headerX := orders.HeaderX.First
	x := 1

	for headerX != nil {
		aux := auxDays{x, headerX.InnerList.First.ID}
		days = append(days, aux)

		headersX = append(headersX, strconv.Itoa(headerX.Order))
		relationsX += getRelations(strconv.Itoa(headerX.Order), strconv.Itoa(headerX.InnerList.First.ID))
		headerX = headerX.Next
		x++
	}
	return days
}

func mapXRelations(days []auxDays, id int, department string, y int) auxDepartments {
	var aux []auxDays
	for i := 0; i < len(days); i++ {
		if days[i].id == id {
			 aux = append(aux, days[i])
		}
	}
	return auxDepartments{y, aux, department}
}

//?????
//Escribir la definicion de los nodos en el .neato
func setNodesContent(matrix []auxDepartments) string {
	var response string
	for k := 0; k < len(headersX); k++ {
		response += getNodeContent(headersX[k], k+1, 0)
	}

	for z := 0; z < len(headersY); z++ {
		response += getNodeContent(headersY[z], 0, z+1)
	}

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row.yRow); j++ {
			cell := row.yRow[j]
			response += getNodeContent(strconv.Itoa(cell.id), cell.x, row.y)					
		}
	}
	return response
}

func getNodeContent(id string, x, y int) string {
	xS := strconv.Itoa(3*x)
	yS := strconv.Itoa(y)
	return "\nnode[label=\"" + id + "\", pos=\"" + xS + ",-" + yS + 
	"!\"]\"n" + id +"\";"
}

//?????
//Escribir las relaciones de los nodos
func setRelations(matrix []auxDepartments) string {
	response := "\n\n"

/// <<<< Aqui debo setear para mas nodos >>>>>


	// prevID := ""
	// for y := 0; y < len(matrix); y++ {
		
	// 	for x := 0; x < len(matrix[y].yRow); x++ {
	// 		id := strconv.Itoa(matrix[y].yRow[x].id)
			
	// 		if x == 0 {
	// 			title := headersY[y]
	// 			response += getRelations(title, id)
	// 		}

	// 	// 	if prevID

	// 	// 	// if x != 0 && y != 0 {
	// 	// 	// 	response += getRelations(prevID, id)
	// 	// 	// }

	// 	// 	// prevID = id
	// 	}
	// }

	// x, y := 0, 0
	// getColumRelations(matrix, &x, &y, &response)

	response += relationsX
	response += realtionsY

	return response
}

func getColumRelations(matrix []auxDepartments, x, y *int, r *string) {
	id := strconv.Itoa(matrix[*y].yRow[*x].id)
	// titleSeted := false

	// if matrix[*y]. {
	// 	title := headersX[*x]
	// 	*r += getRelations(title, id)
	// }
	
	fmt.Println(id)

	if *y < len(matrix) {
		*y++
	} else {
		*y = 0
		if *x < len(matrix[*y].yRow) {
			*x++
		} else {
			*x = 0
		}
	}

	if *y < len(matrix) && *x < len(matrix[*y].yRow) {
		getColumRelations(matrix, x, y, r)
	}	
}

func getRelations(father, child string) string {
	return "\n\"n" + father + "\"->\"n" + child + "\";" +
			"\n\"n" + child + "\"->\"n" + father + "\";"
}
