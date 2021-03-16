package products

import "fmt"

type Product struct {
	Name string
	Code int
	Descr string
	Price int
	Quant int
	Image string
	left *Product
	right *Product
	feq int
}

type Tree struct {
	Root *Product
	Size int
}

func NewTree() *Tree {
	return &Tree{nil, 0}
}

var already bool

func (tree *Tree) InsertNode(code, price, quant int, name, desc, image string) {
	new := &Product{name, code, desc, price, quant, image, nil, nil, 0}

	already = false

	if tree.Root == nil {
		tree.Root = new
		tree.Size++
	} else {
		insertNode(tree, tree.Root, new)
	}
	return already

}

func insertNode(tree *Tree, root, new *Product) {

}

