package account

type Bnodes struct {
	First *Bnode
	Last *Bnode
	Size int
}

func (bnodes *Bnodes) Insert(nuv *Bnode) bool {
	first := bnodes.First
	last := bnodes.Last

	if bnodes.First == nil {
		first = nuv
		last = nuv
		bnodes.aux()
	} else {
		if first == last {
			if nuv.Dpi < first.Dpi {
				nuv.Next = first
				first.Prev = nuv
				first.Left = nuv.Right
				first = nuv
				bnodes.aux()
			} else if nuv.Dpi > first.Dpi {
				last.Next = nuv
				nuv.Prev = last
				last.Right = nuv.Left
				last = nuv
				bnodes.aux()
			} else {
				return false
			}
		} else {
			if nuv.Dpi < first.Dpi {
				nuv.Next = first
				first.Prev = nuv
				first.Left = nuv.Right
				first = nuv
				bnodes.aux()
			} else if nuv.Dpi > last.Dpi {
				last.Next = nuv
				nuv.Prev = last
				last.Right = nuv.Left
				last = nuv
				bnodes.aux()
			} else {
				pivot := first
				for pivot != nil {
					if nuv.Dpi < pivot.Dpi {
						nuv.Next = pivot
						nuv.Prev = pivot.Prev

						pivot.Left = nuv.Right
						pivot.Prev = nuv.Left

						pivot.Prev.Next = nuv
						pivot.Prev = nuv
						bnodes.aux()
					} else if nuv.Dpi == pivot.Dpi {
						return false
					} else {
						pivot = pivot.Next
					}
				}
			}
		}
	}
	return false
}

func (bnodes *Bnodes) aux() bool {
	bnodes.Size++
	return true
}