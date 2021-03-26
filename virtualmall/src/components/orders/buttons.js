import React from 'react'
import Api from '../../api/api'

export default class buttons extends React.Component {

    constructor(props) {
        super(props)
        this.years = this.years.bind(this)
        this.months = this.months.bind(this)
        this.matrix = this.matrix.bind(this)
        this.days = this.days.bind(this)
        this.checkProps = this.checkProps.bind(this)
    }

    render() {
        return (
            <div className="yearsContainer buttonsContainer">
                <input 
                    type="button" 
                    value="Guardar Matriz" 
                    className="buttonOrders" 
                    onClick={() => this.matrix()}
                />

                <input 
                    type="button" 
                    value="Guardar Estructura Years" 
                    className="buttonOrders" 
                    onClick={() => this.years()}
                />

                <input 
                    type="button" 
                    value="Guardar Estructura Meses" 
                    className="buttonOrders" 
                    onClick={() => this.months()}
                />

                <input 
                    type="button" 
                    value="Guardar Pedidos por Dia" 
                    className="buttonOrders" 
                    onClick={() => this.days()}
                />

            </div>
        )
    }

    async matrix() {
        if (this.checkProps()) {
            let img = await Api.getMatrixGraph(this.props.pds.Fecha.value)
            this.props.setImg(img)
        } 
    }

    async years() {
        let img = await Api.getYearsGraph()
        this.props.setImg(img)
    }

    async months() {
        let img = await Api.getMonthsGraph()
        this.props.setImg(img)
    }

    async days() {
        if (this.checkProps()) {
            let res = await Api.getDays(this.props.pds.Fecha.value)
            let text = res.Pedidos.map((item, key) => {
                return <p key={key}>
                            -- Orden: {item.ID} <br />
                            -- Tienda: {item.Tienda} <br />
                            -- Calificacion: {item.Calificacion} <br />
                            -- Departamento: {item.Departamento} <br />
                            -- Fecha: {item.Fecha} <br />
                            -- Productos: {item.Productos.lenght} <br />
                            {item.Productos.map(item2 => {
                                return <div>---- Codigo: {item2.Codigo} <br /> </div>
                            })}
                            <br />
                            <br />
                        </p>
            })
            this.props.setTxt(text)
        }
    }

    checkProps() {
        if (this.props.pds.Fecha == null) {
            this.props.setTxt("Seleccione un mes/dia")
            return false
        }
        return true
    }
}