import React from 'react'
import "../../styles/stores.css"
import Api from "../../api/api"
import Store from "./store"


export default class Stores extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            stores: []
        }
    }

    async componentDidMount() {
        let stores = await Api.getStores()
        this.setState({ stores: stores })
    }

    render() {
        return (
            <div className="content">
                <h1>Stores</h1>
                <div className="storesContainer">
                    { this.mapDepartments() } 
                </div>
                            
            </div>
        )
    }

    mapDepartments() {
        let list = []
        if (this.state.stores.length != 0) {

            this.state.stores.Vector.map((stores, key) => { 
                if (stores.Stores.Size != 0 ) {

                    let aux = this.mapStores(stores.Stores, key)
                    list = [...list, ...aux]
                    //aux.map((item) => { list.push(item)})
                }   
             })
        }
        return list
    }

    mapStores(stores, key) {
        let list = []
        let store = stores.Start
        for (let i = 0; i < stores.Size; i++) {
            list.push(<Store store={store} key={key + i}/>)
            store = store.Next
        }
        return list
    }
}