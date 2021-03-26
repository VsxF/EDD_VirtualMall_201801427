import React from 'react'
import "../styles/cargas.css"
import Dnd from "./Dnd"
import Api from "../api/api"

export default class Cargas extends React.Component {
  constructor(props) {
    super(props)
    this.handleGetFile = this.handleGetFile.bind(this)
    this.state = {
      message: "",
      color: "mSucces"
    }
  }

  render() {
    return (
      <div className="content">
        <center>
          <h2> Manejar Datos </h2>
          
          <label className="radio"><input type="radio" name="cargaType" value="2" id="stores" />Tiendas</label>
          <label className="radio"><input type="radio" name="cargaType" value="1" id="inventa" />Inventario</label>
          <label className="radio"><input type="radio" name="cargaType" value="0" id="pedidos" />Pedidos</label>
          
          <div className={ "message " + this.state.color}>
            { this.state.message }
          </div>
          <Dnd save={this.handleGetFile} />

          

          {/* <input type="button" value="Descargar" className="buttonS" onClick={this.download}/> */}
        </center>
      </div> 
    )
  }

  async handleGetFile(file) {
    if (document.querySelector('input[name="cargaType"]:checked') !== null) {
      if (file[0].name != null) {
        
        const reader = new FileReader()
        reader.onerror = () => console.log("file error")
        reader.onload = async () => {
          
          let response = JSON.parse(reader.result)
          switch (document.querySelector('input[name="cargaType"]:checked').value) {
            case "0":
              let res = await Api.setOrders(response) 
              this.showMessage(res, "Pedidos Agregados !")
              break;
            case "1":
              let res1 = await Api.setInvetorys(response) 
              this.showMessage(res1, "Invetarios Agregados !")
              break;
            case "2":
              let res2 = await Api.setStores(response) 
              this.showMessage(res2, "Tiendas Agregadas !")
              break;
          }
        }
        reader.readAsText(file[0])  
        
      }
    } else {
      this.showMessage("", "")
    }
    
 }

 async showMessage(res, state) {
  let color = "mSucces" 
  if (res === "") {
    state = "Error!"
    color = "mError"
   }
   this.setState({ message: state, color: color })
   await this.sleep()
   this.setState({ message: "" })
 }

  async sleep() {
    return new Promise(resolve => setTimeout(resolve, 3000))
  }

}