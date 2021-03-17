import React from 'react'
import ReactDOM from 'react-dom';
import "../../styles/stores.css"
import Products from "../products/products"

export default class Store extends React.Component {

    constructor(props) {
        super(props)
        this.goToStores = this.goToStores.bind(this)
        this.state = {
            products: this.getValue(this.props.store.Productos)
        }
    }

    render() {
        let logo = this.getValue(this.props.store.Logo)
        let name = this.getValue(this.props.store.Nombre)
        
        return (
            <div className="store" onClick={this.goToStores }>
                <img src={logo} alt="Logo" className="image" />
                <br />
                <label> { name } </label>
            </div>
        )
    }

    goToStores() {
        ReactDOM.render(<Products products={this.state.products} />, document.getElementById("content"));
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