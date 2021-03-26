import React from 'react'
import "../../styles/product.css"
import Add from "../../media/addtocart.svg"
import Added from "../../media/added.svg"
import Api from "../../api/api"

export default class product extends React.Component {

    constructor(props) {
        super(props)
        this.iconPress = this.iconPress.bind(this)

        this.state = {
            name: this.getValue(this.props.info.Nombre),
            code: this.getValue(this.props.info.Codigo),
            desc: this.getValue(this.props.info.Descripcion),
            price: this.getValue(this.props.info.Precio),
            quant: this.getValue(this.props.info.Cantidad),
            image: this.getValue(this.props.info.Imagen),
            icon: Add
        }
    }

    render() {
        return (
            <div className="product"> 
                <div className="imageContainer">
                    <img src={this.state.image} alt="Product" className="image" />

                    <div className="circle" onClick={ () => this.iconPress() }>
                        <img src={this.state.icon} alt="add" className="icon"/>
                    </div>
                </div>
        
               <div className="spects" onClick={ () => this.delete() }>
                  <div className="name"> {this.state.name} </div>
                  <div className="price"> Q. {this.state.price/100} </div>
                  <div className="quant"> Existentes: {this.state.quant} </div> 
               </div>
               {/* <div className="onHover">
                    <div > {this.state.desc} </div>
               </div> */}
               
                 
            </div>  
        )
    }

    iconPress() {
        if (this.state.icon === Add ) {
            this.add()
            this.setState({ icon: Added })
        } else {
            this.delete()
            this.setState({ icon: Add })
        }
    }
    
    async add() {
        let json = {
            Nombre: this.state.name,
            Codigo: this.state.code,
            Descripcion: this.state.desc,
            Precio: this.state.price,
            Cantidad: 1,
            Imagen: this.state.image,
            Left: this.props.storeInfo
        }

        let a = await Api.add2Cart(json)
        //CONTADOR DE CARRITO
    }

    async delete() {
     await Api.delCart(this.state.code) 
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