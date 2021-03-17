import React from 'react'
import "../../styles/stores.css"
import Product from "./product"

export default class products extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            root: this.getValue(this.props.products.Root),
            size: this.getValue(this.props.products.Size)
        }
    }

    render() {
        console.log(this.state.root)
        return(
            <div className="content">
                <h1>Productos</h1>
                <div className="storesContainer">
                    <Product info={this.state.root}/>
                </div> 
            </div>
        )
    }

    getValue(node) {
        if (node == null) {
            return null
        } 
        if (node.kind != null) {
            node = this.getValue(node.value)
        }
        if (typeof(node) === 'number') {
            return null
        }
        return node
    }
}