package orders
//import "fmt"

type Year struct {
	Left *Year
	Right *Year
	balance int
	Year int
	Months *Months
}

type Calendar struct {
	Root *Year
	Size int
}

func NewCalendar() *Calendar {
	return &Calendar{nil, 0}
}

var already bool

func (tree *Calendar) InsertYear(year int, months *Months) bool {
	new := &Year{nil, nil, 0, year, months}
	already = false
	if tree.Root == nil {
		tree.Root = new
		tree.Size++
	} else {
		insertYear(tree, tree.Root, new)
		return already
	}
	return already
}

func insertYear(tree *Calendar, root, new *Year) {
	if new.Year > root.Year {
		if root.Right == nil {
			root.Right = new
			tree.Size++
		} else {
			insertYear(tree, root.Right, new)
		}
	} else if new.Year < root.Year {
		if root.Left == nil {
			root.Left = new
			tree.Size++
		} else {
			insertYear(tree, root.Left, new)
		}
	} else if new.Year == root.Year {
		aux := new.Months.Start
		
		for !aux.HasOrders {
			aux = aux.Next
		} 

		root.Months.setMonth(aux.Month, aux.Orders)
		already = true
	}
	balance(tree, root)
}

func (tree *Calendar) GetYear(year int) *Year {
	var response = getYear(tree.Root, year)
	return response 
}

func getYear(root *Year, year int) *Year {
	if root == nil {
		return nil
	} else if root.Year == year {
		return root
	} else {
		var value1 *Year
		if year > root.Year {
			value1 = getYear(root.Right, year)
		} else if year < root.Year {
			value1 = getYear(root.Left, year)
		}
		return value1
	}
}

func (tree *Calendar) GetYearQuantity() int {
	return tree.Size
}

//Retornar profundidad
func (tree *Calendar) GetDepth() int {
	depth := getDepth(tree.Root)
	return depth
}

func getDepth(root *Year) int {
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

func getParent(root *Year, year int) *Year {
	if year > root.Year {
		if year == root.Right.Year {
			return root
		} else {
			return getParent(root.Right, year)
		}
	} else if year < root.Year {
		if year == root.Left.Year {
			return root
		} else {
			return getParent(root.Left, year)
		} 
	} else {
		return nil
	}
}

//Rotaciones
//Rotacion left-left
func rotLL(tree *Calendar, node1, node2 *Year) {
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
		temp := getParent(tree.Root, node1.Year)
		if temp.Left == node1 {
			temp.Left = node2
		} else {
			temp.Right = node2
		}
	}
}

//Rotacion Right-Right
func rotRR(tree *Calendar, node1, node2 *Year) {
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
		temp := getParent(tree.Root, node1.Year)
		if temp.Left == node1 {
			temp.Left = node2
		} else {
			temp.Right = node2
		}
	}
}

//Rotacion left-Right
func rotLR(tree *Calendar, node1, node2, node3 *Year) {
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
		temp := getParent(tree.Root, node1.Year) 
		if temp.Left == node1 {
			temp.Left = node3
		} else {
			temp.Right = node3
		}
	}
}

//Rotacion Right-Left
func rotRL(tree *Calendar, node1, node2, node3 *Year) {
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
		temp := getParent(tree.Root, node1.Year) 
		if temp.Left == node1 {
			temp.Left = node3
		} else {
			temp.Right = node3
		}
	}
}

//Equilibra el arbol
func balance(tree *Calendar, root *Year) {
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