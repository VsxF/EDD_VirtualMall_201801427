import axios from 'axios'

const baseUrl = "http://localhost:3000/"
const headers = { "x-Requested-With": "XMLHttpRequest" }

class api {

    //??????????????????????
    // Stores
    async setStores(info) {
        let res = await axios({
            url: "cargartienda",
            method: 'post',
            baseURL: baseUrl,
            headers: headers,
            data: info
        })
        .then( function (res) {
            return res.data
        })
        return res
    }

    async getStores() {
        let res = await axios({
            url: "getTiendas",
            method: 'get',
            baseURL: baseUrl,
            headers: headers
        })
        .then(function (res) {
            return res.data
        });
        return res
    }

    async setInvetorys(info) {
        let res = await axios({
            url: "/setInventorys",
            method: "post",
            baseURL: baseUrl,
            headers: headers,
            data: info
        })
        .then(function (res) {
            return res.data
        })
        return res
    }
    

    //??????????????????????????????
    // Cart - Carrito de compras
    async add2Cart(info) {
        let res = await axios({
            url: "add2Cart",
            method: "post",
            baseURL: baseUrl,
            headers: headers,
            data: info     
        })
    }

    async delCart(code) {
        let res = await axios({
            url: "delCart/" + code,
            method: "delete",
            baseURL: baseUrl,
            headers: headers
        })
        .then( function (res) {
            return res.data
        })
        return res
    }
    
    async getCart() {
        let res = await axios({
            url: "getCart",
            method: "get",
            baseURL: baseUrl,
            headers: headers
        })
        .then( function (res) {
            return res.data
        })
        return res
    }

    async buyCart() {
        let res = await axios({
            url: "buyCart",
            method: "post",
            baseURL: baseUrl,
            headers: headers
        })
        return "idk"
    }

    //?????????????????????????????????????
    //Orders
    async getOrders() {
        let res = await axios({
            url: "getOrders",
            method: "get",
            baseURL: baseUrl,
            headers: headers
        })
        .then(function (res) {
            return res.data
        })
        return res
    }

    async setOrders(info) {
        let res = await axios ({
            url: "/setOrders",
            method: "post",
            baseURL: baseUrl,
            headers: headers,
            data: info
        })
        .then(function (res) {
            return res.data
        })
        return res
    }

    //?????????????????????????????????????
    //Reports
    async getMatrixGraph(date) {
        return await axios({
            url: "/getMatrixGraph/" + date,
            method: "get",
            baseURL: baseUrl,
            headers: headers,
        })
        .then((res) => {
            return res.data
        })
    }

    async getYearsGraph() {
        return await axios({
            url: "/getYearsGraph",
            method: "get",
            baseURL: baseUrl,
            headers: headers,
        })
        .then((res) => {
            return res.data
        })
    }

    async getMonthsGraph() {
        return await axios({
            url: "/getMonthsGraph",
            method: "get",
            baseURL: baseUrl,
            headers: headers,
        })
        .then((res) => {
            return res.data
        })
    }

    async getDays(date) {
        return await axios({
            url: "/getDays/" + date,
            method: "get",
            baseURL: baseUrl,
            headers: headers,
        })
        .then((res) => {
            return res.data
        })
    }


}

let aux = new api()
export default aux