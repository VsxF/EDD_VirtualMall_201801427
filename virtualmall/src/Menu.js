import ReactDOM from 'react-dom';
import "./styles/Menu.css"
import Index from './indexView'
import Cargas from "./components/Cargas" 
import Stores from "./components/stores/Stores"

function Menu() {
    redirect()
    return (
        <div className="bar">
            <div className="item" onClick={() => redirect()} >Virtual Mall</div>
            <div className="item" onClick={() => redirect(1)} >Tiendas</div>
            <div className="item" onClick={() => redirect(4)} >Datos</div>
        </div>
    )
}

function redirect(opt) {
    switch (opt) {
        case 1:
            ReactDOM.render(<Stores />, document.getElementById("content"));
            break;
        case 4:
            ReactDOM.render(<Cargas />, document.getElementById("content"));
            break;
        default:
            ReactDOM.render(<Index />, document.getElementById("content"));
            break;
    }
}

export default Menu;