package account

type Btree struct {
	Root *Leaf
	Height int 
} 

func NewBtree() *Btree {
	return &Btree{nil, 0}
}

func (tree *Btree) Insert(nuv *Bnode) {
	if tree.Root == nil {
		nuv.Leaf = true
		tree.Root = nuv
	} else {
		if tree.Height == 0 {
			if !tree.Root.Insert(nuv) {
				tree.Height++
				//Insertar hoja
			} 
		} 
		// .....
	}
}
