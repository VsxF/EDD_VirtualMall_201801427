import React from 'react'
import "../../styles/product.css"

export default class product extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            name: this.getValue(this.props.info.Nombre),
            code: this.getValue(this.props.info.Codigo),
            desc: this.getValue(this.props.info.Descripcion),
            price: this.getValue(this.props.info.Precio),
            quant: this.getValue(this.props.info.Cantidad),
            image: this.getValue(this.props.info.Imagen)
        }
    }

    render() {
        console.log(this.state.quant)
        return (
            <div className="product"> 
               <img src={this.state.image} alt="Product" className="image" />
               <div> {this.state.Name} </div>
               <div> {this.state.Price} </div>
               <div> {this.state.desc} </div>
               <div> Existentes: {this.state.quant} </div>   
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
        return node
    }
}