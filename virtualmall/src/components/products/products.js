import React from 'react'
import "../../styles/stores.css"
import Product from "./product"

let productList = [];

export default class products extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            list: []
        }
    }

    componentDidMount() {
        this.mapProducts()
    }

    render() {
        console.log(this.props.products.Root)
        return(
            <div className="content">
                <h1>Productos</h1>
                <div className="storesContainer">
                    {this.state.list}
                    
                </div> 
            </div>
        )
    }

    //Falta mostrar todos los items
    mapProducts() {
        let root = this.getValuee(this.props.products.Root)
        productList = []
        if (root !== null) {
            this.inOrden(root)

            let list = productList.map((item, key) => {
                return <Product info={item} key={key} storeInfo={this.props.storeInfo} />
            })     
            this.setState({ list: list })
        }
    }

    //Los productos vienen en un arbol avl
    inOrden(node) {
        node = this.getValuee(node)
        if (node !== null) {
            
            let left = this.getValuee(node.Left)
            let right = this.getValuee(node.Right)
            productList = [...productList, node]
            
            if (left !== null) {
                this.inOrden(left)
            }   

            if (right !== null) {
                this.inOrden(right)
            }
        }  
    }

    getValuee(node) {
        if (node == null) {
            return null
        } 
        if (node.kind != null) {
            node = this.getValuee(node.value)
        }
        if (typeof(node) === 'number') {
            return null
        }
        return node
    }
}