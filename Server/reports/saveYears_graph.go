package reports

import (
	"strconv"
	"os/exec"
	ord "../data/orders"
)
var content string
func GetYearsGraph(orders *ord.Calendar) []byte {
	content = "digraph G {"
	printYears(orders.Root)
	content += "\n}"
	
	CreateFile(File{"years", content, ".dot"})
	exec.Command("dot", "-Tpng", "years.dot", "-o", "years.png").Output()

	image := image2base64("years.png")

	return []byte(image)
}

func printYears(year *ord.Year) {
	if year.Left != nil {
		content += "\n\"" + strconv.Itoa(year.Year) + "\"->\"" + strconv.Itoa(year.Left.Year) + "\""
		printYears(year.Left)
	}
	if year.Right != nil {
		content += "\n\"" + strconv.Itoa(year.Year) + "\"->\"" + strconv.Itoa(year.Right.Year)+ "\""
		printYears(year.Right)
	}
}