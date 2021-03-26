import React from 'react'
import ReactModal from 'react-modal'
import "../../styles/orders.css"
import Api from "../../api/api"
import Years from "./years"
import Months from "./months"
import ProductInfo from './productInfo'
import Buttons from "./buttons"
import GetValue from "../getValue"


export default class orders extends React.Component {

    constructor(props) {
        super(props)
        this.handdleMonthClick = this.handdleMonthClick.bind(this)
        this.handdleProductClick = this.handdleProductClick.bind(this)
        this.handdleCloseModel = this.handdleCloseModel.bind(this)
        this.handdleImgB64 = this.handdleImgB64.bind(this)
        this.handdleSetTextModal = this.handdleSetTextModal.bind(this)

        this.state = {
            info: [],
            month: [],
            products: [],
            date: "",
            showModal: false,
            textModal: "",
            imgb64: []
        }
    }

    async componentDidMount() {
        let res = await Api.getOrders()
        res = GetValue(res)
        if (res !== null) {
            this.setState({ info: res })
        }   
    }

    render() {
        return (
            <div className="ordersContainer">

                <h1>Ordenes</h1>
                <h2 className="monthTitle"> &nbsp;{GetValue(this.state.month.Name)} </h2>
                <div className="information">
                    <Years 
                        info={this.state.info} 
                        monthClick={this.handdleMonthClick} 
                    /> 
                    <Months 
                        month={this.state.month}
                        productClick={this.handdleProductClick} 
                    />
                </div>
                <div className="information">
                    <Buttons 
                        setImg={this.handdleImgB64}
                        setTxt={this.handdleSetTextModal}
                        pds={this.state.products}
                    />
                    <ProductInfo 
                        products={this.state.products}
                    />
                </div>
                

                <ReactModal 
                    isOpen={this.state.showModal}
                    contentLabel=""
                    ariaHideApp={false}
                    onClick={ () => this.handdleCloseModel()}
                >
                    {this.aux()}
                    <button onClick={ () => this.handdleCloseModel()}>Close [X]</button>
                </ReactModal>

            </div>
        )
    }

    aux() {
        if (this.state.imgb64.length > 0) {
            return <img src={"data:image/png;base64, " + this.state.imgb64} alt=":C"/>
        } else {
            return <h2> {this.state.textModal} </h2>
        }
    }

    handdleMonthClick(month) {
        this.setState({ month: month })
    }

    handdleProductClick(products) {
        this.setState({ products: products })
    }

    handdleCloseModel() {
        this.setState({ showModal: false })
    }

    handdleImgB64(img) {
        this.setState({ imgb64: img, showModal: true })
    }

    handdleSetTextModal(text) {
        this.setState({ textModal: text, showModal: true, imgb64: [] })
    }
}