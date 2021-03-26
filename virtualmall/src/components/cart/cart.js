import React from 'react'
import "../../styles/stores.css"
import Api from "../../api/api"
import CartItem from "./cartItem"

export default class cart extends React.Component {

    constructor(props) {
        super(props)
        this.buyProducts = this.buyProducts.bind(this)
        this.handdleDeleteProduct = this.handdleDeleteProduct.bind(this)
        this.state = {
            list: []
        }
    }

    async componentDidMount() {
        let products = await Api.getCart()
        this.mapPoducts(products)        
    }

    render() {
        return (
            <div className="content">
                <h1>Carro de Compras</h1>
                <div className="storesContainer">
                    { this.state.list }
                </div>
                <div>
                    <input type="button" value="Comprar" className="buttonS" onClick={() => this.buyProducts()} />
                </div>

            </div>
        )
    }

    mapPoducts(products) {
        let list = []
        products = products.Products
        if (products.lengh !== 0) {
            list = products.map((item, key) => {
                return <CartItem info={item} key={key} delete={this.handdleDeleteProduct} />
            })
            this.setState({ list: list })
        }  
    }

    async buyProducts() {
        let a = await Api.buyCart()
        if (a != null) {
            this.setState({ list: [] })
            //Mostrar estado de vendido 
        }
    }

    async handdleDeleteProduct(code) {
        let a = await Api.delCart(code)
        this.mapPoducts(a)
    }
}