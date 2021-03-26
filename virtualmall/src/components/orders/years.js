import React from 'react'
import Collapsible from 'react-collapsible';
import GetValue from '../getValue'

let yearsList = []
let monthList = []
export default class years extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
           list: [],
           monthCount: 0
        }
    }

    componentWillReceiveProps(props) {
        yearsList = []
        let root = GetValue(props.info.Root)
        this.setList(root)
    }

    render() {         
        return (
            <div className="yearsContainer">
                {this.state.list} 
            </div>
        )
    }

    //arbol avl
    inOrden(node) {
       if (node !== undefined && node !== null) {
            let left = node.Left
            let right = node.Right
            if (left !== null) {
                this.inOrden(left)
            }   
            yearsList = [...yearsList, node]
            if (right !== null) {
                this.inOrden(right)
            }
        }  
    }
    
    setList(root) {
        this.inOrden(root)
        let list = yearsList.map((year, key) => {
            year = GetValue(year)
            monthList = []
            let months = GetValue(year.Months)
            let yearINT = year.Year.value
            this.setMonths(months)
            let yearText = yearINT + " [" + monthList.length +"]"
            
            return (
                <Collapsible 
                    trigger = { this.getYearText("> " + yearText) } 
                    triggerWhenOpen = { this.getYearText("< " + yearText ) } 
                    key = {key}
                    onClick = {this.props.monthClick}
                >
                    {this.getMonths()}
                </Collapsible>
            )
        })
        this.setState({ list: list })
    }

    setMonths(months) {
        let next = GetValue(months.Start)
        for (let i = 0; i < 12; i++) {
            if (GetValue(next.HasOrders)) {
                monthList = [...monthList, next]
            }
            next = GetValue(next.Next)
        }
    }

    getMonths() {
        return monthList.map((month, key) => {
                let name = GetValue(month.Name)
                return <input 
                            type = "button"
                            key = {key} 
                            className = "yearOpen clickableText" 
                            onClick = {() => this.props.monthClick(month) }
                            value = {name}
                       />     
        })
    }

    getYearText(year) {
        return <input 
                    type = "button"
                    className = "year clickableText" 
                    value = {year}
                />
    }
}