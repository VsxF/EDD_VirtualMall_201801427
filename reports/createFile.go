package reports

import (
	"fmt"
	"os"
)

type file struct {
	Name    string
	Content string
}

func (file *file) NewFile(name string) {
	file.Name = name
	file.Content = ""
}

func (file *file) AddText(cont string) {
	file.Content += cont
}

func CreateFile(file file) {
	f, err := os.Create(file.Name + ".dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(file.Content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
