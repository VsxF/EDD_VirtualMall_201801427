package account

const max = 5
const min = 1

type Leaf struct {
	Keys *Bnodes
	Root bool
	Size int
}

//Manejar leaf desde memoria
func Insert(nuv *Bnode, leaf *Leaf) {
	if leaf.Keys.Insert(nuv) {
		if leaf.Keys.Size < 5 {
			return 
		}
	}

	return true
}