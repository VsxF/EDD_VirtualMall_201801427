import React from 'react'
import "../../styles/stores.css"
import Api from "../../api/api"
import Store from "./store"

export default class Stores extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            list: []
        }
    }

    async componentDidMount() {
        let stores = await Api.getStores()
        this.mapDepartments(stores)
        //console.log(stores)
    }

    render() {
        return (
            <div className="content">
                <h1>Tiendas</h1>
                <div className="storesContainer">
                    { this.state.list } 
                </div>                         
            </div>
        )
    }

    mapDepartments(stores) {
        let list = []
        if (stores.length !== 0) {
            let aux = stores
            aux = this.getValue(aux)
            aux = this.getValue(aux.Vector)
            
            aux.map((stores, key) => {
                stores = this.getValue(stores)
                if (stores.Stores.Size !== 0 ) {
                    let aux2 = this.mapStores(stores.Stores, key)
                    list = [...list, ...aux2]
                    //aux.map((item) => { list.push(item)})
                }   
            })
        }
        this.setState({ list: list })   
    }

    mapStores(stores, key) {
        stores = this.getValue(stores)
        let list = []
        let store = stores.Start
        store = this.getValue(store)

        for (let i = 0; i < stores.Size.value; i++) {
            if (store !== null) {
                let aux = this.getValue(store.Nombre) + store.Calificacion.value
                list.push(<Store store={store} key={aux}/>)
            }
            store = this.getValue(store.Next)
        }
        return list
    }

    //El vector de GO tiene cyclos "recursivos"
    //solo se puede convertir a json con una estructura mas extensa y diferente.
    //Eso en el backend, pero la peticion retorna un vector con un orden relativamente distinto
    //para ir obteniendo la informacion "util" es esta funcion
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