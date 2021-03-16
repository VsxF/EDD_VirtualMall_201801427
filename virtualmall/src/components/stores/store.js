import React from 'react'
import "../../styles/stores.css"

export default class Store extends React.Component {

    constructor(props) {
        super(props)
        this.goToStores = this.goToStores.bind(this)
    }

    render() {
        return (
            <div className="store" onClick={this.goToStores }>
                <img src={this.props.store.Logo} alt="Logo" className="image" />
                <br />
                <label> { this.props.store.Nombre } </label>
            </div>
        )
    }

    goToStores() {
        alert('so')
    }
    
}