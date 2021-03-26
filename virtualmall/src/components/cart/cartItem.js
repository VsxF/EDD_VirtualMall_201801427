import React from 'react'
import "../../styles/cart.css"

export default class cartItem extends React.Component {

    constructor(props){
        super(props) 
    }

    render() {
        return(
            <div className="container">
                <center>
                    <div className="cartItem">
                        <img src={this.props.info.Imagen} alt="producto" className="itemImage"/>
                        <div className="itemContent">
                            
                            <div>{this.props.info.Nombre} </div> 
                            <div>cant. {this.props.info.Cantidad} </div>   
                            <div>Q. {this.props.info.Precio/100} </div>   
                        </div>    
                        <div className="itemButton">
                            <input 
                                type="button" 
                                value="x" 
                                className="button" 
                                onClick={() => this.props.delete(this.props.info.Codigo)} 
                            />    
                        </div>    
                    </div> 
                </center>           
            </div>
        )
    }
}