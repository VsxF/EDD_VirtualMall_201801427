import React from 'react'
import GetValue from '../getValue'

let tableWith = "99%"
let tdWith = "20%"
let departments = []
let days = []
let data  = []

export default class months extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            list: []
        }
    }

    componentWillReceiveProps(props) {
        if (props.month !== null && props.month.length !== 0) {
            this.getDataInfo(props)
        }
    }

    render() {
        
        return (
            <div className="monthContainer">
                <table className="monthTable" width={tableWith}>
                    {this.state.list}
                </table>
            </div>
        )
    }

    getDataInfo(props) {
            let orders = GetValue(GetValue(props.month).Orders)
           
            orders = GetValue(GetValue(orders.HeaderX).First)

            departments = []
            days = []
            let a = 1
            let list = []

            for (let i = 0; i < a; i++) {
                if (orders !== null) {
                    list = [...list,  this.getMonthInfo(orders)]
                    orders = GetValue(orders.Next)
                    a++
                }
            }
            list = [this.getDays(), ...list]
            console.log(departments)
            console.log(days)
            this.setState({ list: list })
    }

    getMonthInfo(month) {
        
        month = GetValue(GetValue(month.InnerList).First)
        
        let department = <td className="monthNode monthGreenHeader" width={tdWith}>{GetValue(month.Departamento)}</td>
        
        let transparent = []
        for (let i = 0; i < days.length; i++) {
            const day = days[i];
            if (day === month.X.value) {
                break
            }
            transparent = [...transparent, <td className="monthNode" width={tdWith}></td>]
        }
        let product = <td 
                        className="monthNode monthGreenNode" 
                        width={tdWith}
                        key={GetValue(month.Fecha)}
                        onClick={() =>this.props.productClick(month)}>
                            {GetValue(month.Fecha)}
                    </td>

        departments = [...departments, GetValue(month.Departamento)]
        days = [...days, month.X.value] 
        data = [...data, month]
        
        return <tr >
                {department}
                {transparent}
                {product}
            </tr>
    }

    getDays() {
        return (
            <tr>
                <td className="monthNode" width={tdWith}>
                    <div className="tableRigthText">Dias</div>
                    <div className="tableButtonText">Departamentos</div>
                </td>

                {days.map(day => { return <td className="monthNode monthGreenHeader" width={tdWith}>{day}</td> } )}
            </tr>
        )
    }
}