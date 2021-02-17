package reports

import (
	"fmt"
	"os"
)

type File struct {
	Name    string
	Content string
	Ext string
}

func NewFile(name string, ext string) *File {
	var file File
	file.Name = name
	file.Content = ""
	file.Ext = ext
	return &file
}

func (file *File) AddText(cont string) {
	file.Content += cont
}

func CreateFile(file File) {
	f, err := os.Create(file.Name + file.Ext)
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
	fmt.Println(l, "Archivo creado")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
