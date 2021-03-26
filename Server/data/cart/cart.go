package cart

import (
	p "../products"
)

type Cart struct {
	Products []*p.Product
}

func NewCart() *Cart {
	return &Cart{ []*p.Product{} }
}

//name string, code int, desc string, price, quant int, image string
func (prods *Cart) InsertProduct(product *p.Product) {
	prods.Products = append(prods.Products, product)
}

func (prods *Cart) DeleteProduct(code int) {
	size := len(prods.Products)
	for i := 0; i < size; i++ {
		if prods.Products[i].Code == code {
			if size-1 == i {
				prods.Products = prods.Products[:i]
			} else if size > 1 {
				prods.Products = append(prods.Products[:i], prods.Products[i+1:]...)
			} else {
				prods.Products = prods.Products[:0]
				size = 0
			}
			break
		}
	}
}