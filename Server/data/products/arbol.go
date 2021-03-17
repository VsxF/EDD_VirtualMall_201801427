package products

type Product struct {
	Left *Product
	Right *Product
	balance int 
	Name string `json:"Nombre"`
	Code int `json:"Codigo"`
	Descr string `json:"Descripcion"`
	Price int `json:"Precio"`
	Quant int `json:"Cantidad"`
	Image string `json:"Imagen"`
	
}

type Tree struct {
	Root *Product 
	Size int 
}

func NewTree() *Tree {
	return &Tree{nil, 0}
}

var already bool

func (tree *Tree) InsertProduct(name string, code int, desc string, price, quant int, image string) bool {
	new := &Product{nil, nil, 0, name, code, desc, price, quant, image}
	
	already = false
	
	if tree.Root == nil {
		tree.Root = new
		tree.Size++
	} else {
		insertProduct(tree, tree.Root, new)
		return already
	}
	return already
}

func insertProduct(tree *Tree, root, new *Product) {
	if new.Code > root.Code {
		if root.Right == nil {
			root.Right = new
		} else {
			insertProduct(tree, root.Right, new)
		}
	} else if new.Code < root.Code {
		if root.Left == nil {
			root.Left = new
		} else {
			insertProduct(tree, root.Left, new)
		}
	} else if new.Code == root.Code {
		already = true
	}
	balance(tree, root)
}

func (tree *Tree) GetProduct(code int) *Product {
	var response = getProduct(tree.Root, code)
	return response 
}

func getProduct(root *Product, code int) *Product {
	if root == nil {
		return nil
	} else if root.Code == code {
		return root
	} else {
		var value1 *Product
		if code > root.Code {
			value1 = getProduct(root.Right, code)
		} else if code < root.Code {
			value1 = getProduct(root.Left, code)
		}
		return value1
	}
}

func (tree *Tree) GetProductQuantity() int {
	return tree.Size
}

//Retornar profundidad
func (tree *Tree) GetDepth() int {
	depth := getDepth(tree.Root)
	return depth
}

func getDepth(root *Product) int {
	if root == nil {
		return 0
	} else {
		var left = getDepth(root.Left)
		var right = getDepth(root.Right)

		if left > right {
			return left +1
		} else {
			return right +1
		}
	}
}

func getParent(root *Product, code int) *Product {
	if code > root.Code {
		if code == root.Right.Code {
			return root
		} else {
			return getParent(root.Right, code)
		}
	} else if code < root.Code {
		if code == root.Left.Code {
			return root
		} else {
			return getParent(root.Left, code)
		} 
	} else {
		return nil
	}
}

//Rotaciones
//Rotacion left-left
func rotLL(tree *Tree, node1, node2 *Product) {
	node1.Left = node2.Right
	node2.Right = node1

	if node2.balance == -1 {
		node1.balance = 0
		node2.balance = 0
	} else {
		node1.balance = -1
		node2.balance = 0
	}
	if tree.Root == node1 {
		node1 = node2
		tree.Root = node2
	} else {
		temp := getParent(tree.Root, node1.Code)
		if temp.Left == node1 {
			temp.Left = node2
		} else {
			temp.Right = node2
		}
	}
}

//Rotacion Right-Right
func rotRR(tree *Tree, node1, node2 *Product) {
	node1.Right = node2.Left
	node1.Left = node1

	if node2.balance == 1 {
		node1.balance = 0
		node2.balance = 0
	} else {
		node1.balance = 1
		node2.balance = 0
	}
	if tree.Root == node1 {
		node1 = node2
		tree.Root = node2
	} else {
		temp := getParent(tree.Root, node1.Code)
		if temp.Left == node1 {
			temp.Left = node2
		} else {
			temp.Right = node2
		}
	}
}

//Rotacion left-Right
func rotLR(tree *Tree, node1, node2, node3 *Product) {
	node1.Left = node3.Right
	node3.Right = node1
	node2.Right = node3.Left
	node3.Left = node2
	
	if node3.balance == 1 {
		node2.balance = -1
	} else {
		node2.balance = 0
	}

	if node3.balance == -1 {
		node1.balance = 1
	} else {
		node1.balance = 0
	}
	node3.balance = 0

	if tree.Root == node1 {
		node1 = node3
		tree.Root = node3
	} else {
		temp := getParent(tree.Root, node1.Code) 
		if temp.Left == node1 {
			temp.Left = node3
		} else {
			temp.Right = node3
		}
	}
}

//Rotacion Right-Left
func rotRL(tree *Tree, node1, node2, node3 *Product) {
	node1.Right = node3.Left
	node3.Left = node1
	node2.Left = node3.Right
	node3.Right = node2
	
	if node3.balance == 1 {
		node1.balance = -1
	} else {
		node1.balance = 0
	}

	if node3.balance == -1 {
		node2.balance = 1
	} else {
		node2.balance = 0
	}
	node3.balance = 0

	if tree.Root == node1 {
		node1 = node3
		tree.Root = node3
	} else {
		temp := getParent(tree.Root, node1.Code) 
		if temp.Left == node1 {
			temp.Left = node3
		} else {
			temp.Right = node3
		}
	}
}

//Equilibra el arbol
func balance(tree *Tree, root *Product) {
	left := getDepth(root.Left)
	right := getDepth(root.Right)

	root.balance = right - left
	if root.balance == -2 {
		if root.Left.balance > 0 {
			rotLR(tree, root, root.Left, root.Left.Right)
		} else {
			rotLL(tree, root, root.Left)
		}
	} else if root.balance == 2 {
		if root.Right.balance < 0 {
			rotRL(tree, root, root.Right, root.Right.Left)
		} else {
			rotRR(tree, root, root.Right)
		}
	}
	
}