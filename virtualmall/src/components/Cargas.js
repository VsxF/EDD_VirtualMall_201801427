import React from 'react'
import "../styles/cargas.css"
import Dnd from "./Dnd"
import Api from "../api/api"

export default class Cargas extends React.Component {
  constructor(props) {
    super(props)
    this.handleSave = this.handleGetFile.bind(this)
    //this.download = this.download.bind(this)
  }

  render() {
    return (
      <div className="content">
        <center>
          <h2> Manejar Datos </h2>
          
          <label className="radio"><input type="radio" name="cargaType" value="0" id="pedidos" />Pedidos</label>
          <label className="radio"><input type="radio" name="cargaType" value="1" id="inventa" />Inventario</label>

          <Dnd save={this.handleGetFile} />

          <input type="button" value="Descargar" className="buttonS" onClick={this.download}/>
        </center>
      </div> 
    )
  }

  handleGetFile(file) {
    if (file[0].name != null) {
      const reader = new FileReader()
      reader.onerror = () => console.log("file error")
      reader.onload = () => {
        let response = JSON.parse(reader.result)
        //if (document.getElementById("inventa").checked) {
                  
          Api.setStores(response)

        //}
        
      }
      reader.readAsText(file[0])  
      console.log(reader)
    }
 }

 download() {
   console.log("down")
  Api.getStores()


  // let a = { "Datos": [
  //       {
  //           "Indice": "A",
  //           "Departamentos": [
  //               {
  //                   "Nombre": "Amazon Coins",
  //                   "Tiendas": [
  //                       {
  //                           "Nombre": "Amaya, Rolón y Chavarría Asociados",
  //                           "Descripcion": "Illum accusantium voluptate voluptatem in corrupti dolorem velit et.",
  //                           "Contacto": "976191834",
  //                           "Calificacion": 5,
  //                           "Logo":"https://economipedia.com/wp-content/uploads/2015/10/apple-300x300.png"
  //                       },
  //                       {
  //                           "Nombre": "Aguayo Cepeda S.A.",
  //                           "Descripcion": "Numquam ea est error inventore et porro veritatis.",
  //                           "Contacto": "949 586 354",
  //                           "Calificacion": 4,
  //                           "Logo":"https://i.pinimg.com/originals/3d/0f/0e/3d0f0e8f600627fde858f6c6e668e999.gif"
  //                       }
  //                   ]
  //               },
  //               {
  //                   "Nombre": "Sensores",
  //                   "Tiendas": [
  //                       {
  //                           "Nombre": "Amaya, Rolón y Chavarría Asociados",
  //                           "Descripcion": "Enim incidunt beatae enim quisquam harum fuga molestiae at.",
  //                           "Contacto": "915 133 743",
  //                           "Calificacion": 3,
  //                           "Logo":"https://as.com/meristation/imagenes/2021/01/18/mexico/1610944753_187605_1610981923_noticia_normal.jpg"
  //                       },
  //                       {
  //                           "Nombre": "Aguayo Cepeda S.A.",
  //                           "Descripcion": "Animi similique quas quam consectetur dolorem.",
  //                           "Contacto": "961.244.645",
  //                           "Calificacion": 4,
  //                           "Logo":"https://graffica.info/wp-content/uploads/2017/01/Kentucky-Fried-Chicken.jpg"
  //                       }
  //                   ]
  //               }
  //           ]
  //       },
  //       {
  //           "Indice": "B",
  //           "Departamentos": [
  //               {
  //                   "Nombre": "Amazon Coins",
  //                   "Tiendas": [
                        
  //                   ]
  //               },
  //               {
  //                   "Nombre": "Sensores",
  //                   "Tiendas": [
  //                       {
  //                           "Nombre": "Briones Pineda S.A.",
  //                           "Descripcion": "Minus esse nemo eveniet sapiente iste sapiente repudiandae sapiente.",
  //                           "Contacto": "918-764-088",
  //                           "Calificacion": 1,
  //                           "Logo":"https://static-cse.canva.com/blob/211898/17-50-logotipos-que-te-inspiraran.jpg"
  //                       }
  //                   ]
  //               }
  //           ]
  //       }
  //   ]
//}
  // console.log(a)
  // Api.setStores(a)
 }

 setFileInfo(content) {
  
 }

}