package reports

import (
	"fmt"
	"os/exec"
	data "../data/stores"
)

func GetComplete(vector *data.Vector) {
	fmt.Println("Complete report")
	content := "digraph G {\n rankdir=TD\n\tnode[shape=box]\n\tcompound=true\n\n"
	
	content += printVector(vector)
	
	content += "\n}"

	CreateFile(File{"pv", content, ".dot"})
	exec.Command("dot", "-Tpdf", "pv.dot", "-o", "pv.pdf").Output()
	//fmt.Println(vector)
}

func printVector(vector *data.Vector) string{
	content := ""
	prevID := ""

	for i := 0; i < len(vector.Vector); i++ {
		vectorr := vector.Vector[i]
		if prevID != "" {
			content += "\t\"" + prevID + "\"->\"" + vectorr.ID + "\" [constraint=false]\n"
		}
		prevID = vectorr.ID

		if vectorr.Stores.Start != nil {
			content += printStores(vectorr.ID, vectorr.Stores)	
		}	
	}
	return content
}

func printStores(id string, stores *data.Stores) string{ 
	content := "\n\tsubgraph \"cluster" + id + "\" {\n"
	content += "\t\t\"" + stores.Start.Name + "\""
	auxNode := stores.Start.Next

	auxContent := "\t\t\"" + stores.Lastest.Name + "\""
	auxNodeLastes := stores.Lastest.Previous

	for auxNode!= nil && auxNodeLastes != nil {
		content += "->\"" + auxNode.Name + "\" [arrowhead=rnormal]"
		auxContent += "->\"" + auxNodeLastes.Name + "\" [arrowhead=rnormal]"
		auxNode = auxNode.Next
		auxNodeLastes = auxNodeLastes.Previous
	}

	if stores.Start.Next != nil {
		content += "\n" + auxContent
	}
	
	content += "\n\t } \n\t\""
	content += id + "\"->\"" + stores.Start.Name + "\" [lhead=\"cluster" + id + "\" arrowhead=none]\n\n"
	return content
}