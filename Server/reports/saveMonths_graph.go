package reports

import (
	"fmt"
	"os/exec"
	"os"
	"bufio"
  	"encoding/base64"
)

func GetMonthsGraph() []byte {
	content = "digraph G {\nrankdir=LR\nnode[shape=box]"
	printMonths("Enero", "Febrero")
	printMonths("Febrero", "Marzo")
	printMonths("Marzo", "Abril")
	printMonths("Abril", "Mayo")
	printMonths("Mayo", "Junio")
	printMonths("Junio", "Julio")
	printMonths("Julio", "Agosto")
	printMonths("Agosto", "Septiembre")
	printMonths("Septiembre", "Octubre")
	printMonths("Octubre", "Noviembre")
	printMonths("Noviembre", "Diciembre")
	content += "\n}"
	
	CreateFile(File{"months", content, ".dot"})
	exec.Command("dot", "-Tpng", "months.dot", "-o", "months.png").Output()

	image := image2base64("months.png")
	return []byte(image)
}

func printMonths(actual, next string) {
	content += "\n\"" + actual + "\"->\"" + next + "\" [arrowhead=rnormal]"
	content += "\n\"" + next + "\"->\"" + actual + "\" [arrowhead=rnormal]"
}
//Se usa en varios reports
func image2base64(dir string) string {
	imgFile, err := os.Open(dir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	return base64.StdEncoding.EncodeToString(buf)
}