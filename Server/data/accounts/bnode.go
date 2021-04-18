package account

type Account struct {
	Email string
	Pw string
	Name string
	Status bool // 0 admin - 1 client
}

type Bnode struct {
	Dpi int
	Account *Account

	Next *Bnode
	Prev *Bnode
	Left *Leaf
	Right *Leaf

	x, y int
}

func NewNode(dpi int, email, pw, name string, status bool) *Bnode {
	acc := &Account{email, pw, name, status}
	return &Bnode{dpi, acc, nil, nil, nil, nil, 0, 0}
}