import ReactDOM from 'react-dom';
import "./styles/Menu.css"
import CartImage from "./media/cart.svg"
import Index from './indexView'
import Cargas from "./components/Cargas" 
import Stores from "./components/stores/Stores"
import Cart from "./components/cart/cart"
import Orders from "./components/orders/orders"


function Menu() {
    redirect()
    return (
        <div className="bar">
            <div className="item" onClick={() => redirect()} >Virtual Mall</div>
            <div className="item" onClick={() => redirect(1)} >Tiendas</div>
            <div className="item" onClick={() => redirect(2)} >Ordenes</div>
            <div className="item" onClick={() => redirect(3)} >Datos</div>
            <div className="item cart" onClick={() => redirect(4) } >
                <img src={CartImage} alt="carro" className="cartImage" />
            </div>
        </div>
    )
}

function redirect(opt) {
    switch (opt) {
        case 1:
            ReactDOM.render(<Stores />, document.getElementById("content"));
            break;
        case 2:
            ReactDOM.render(<Orders />, document.getElementById("content"));
            break;
        case 3:
            ReactDOM.render(<Cargas />, document.getElementById("content"));
            break;
        case 4:
            ReactDOM.render(<Cart />, document.getElementById("content"));
            break;
        default:
            ReactDOM.render(<Index />, document.getElementById("content"));
            break;
    }
}

export default Menu;