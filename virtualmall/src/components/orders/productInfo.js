import React from 'react'
import GetValue from "../getValue"

export default class productInfo extends React.Component {

    state = {
        list: []
    }

    componentWillReceiveProps(props) {
        if (props.products.length !== 0 && props.products !== undefined) {
            this.setProduct(props.products)
        }       
    }

    render() {
        return (
            <div className="yearsContainer infoContainer">
                <div className="tableBorder">
                <table width="98%" height="92%" className="tableInfo">
                    <thead className="theadInfo">
                        <tr>
                            <th width="17%">Numero de Pedido</th>
                            <th width="20%">Cliente</th>
                            <th width="63%">Direccion (codigo productos)</th>
                        </tr>
                   </thead>
                    <tbody className="tbodyInfo">
                        {this.state.list}
                    </tbody>
                </table>
                </div>
            </div>
        )
    }

    setProduct(product) {
        let list = <tr>
                        <td>{ product.ID.value }</td>
                        <td>No hay cliente Dx</td>
                        <td>{ this.getProducts(product.Productos) }</td>
                    </tr>

    this.setState({ list: list })
    }

    getProducts(products) {
        products = GetValue(products)  
        return products.map(item => {
            item = GetValue(item)
            return item.Codigo.value + ","
        })
    }
}