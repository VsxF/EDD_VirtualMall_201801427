import axios from 'axios'

const base = "localhost:3000/"
axios.baseURL = 'http://localhost:3000/'
class api {

    constructor() {
        
    }

    async setStores(info) {
        // axios.post(base + "cargartienda", { 
        //     "headers": {
        //         "content-type": "application/json",
        //         },

        // })
        // .then(res => {
        //     console.log(res)
        //     return "a"
        // })
        console.log(info)
        axios({
            url: "cargartienda",
            method: 'post',
            baseURL: 'http://localhost:3000/',
            headers: {
                'X-Requested-With': 'XMLHttpRequest'},
            data: info
        })

    }

    async getStores() {
        let res = await axios({
            url: "getTiendas",
            method: 'get',
            baseURL: 'http://localhost:3000/',
            headers: {
                'X-Requested-With': 'XMLHttpRequest',
                'Access-Control-Allow-Origin': '*'
            }
        })
        .then(function (res) {
            return res.data
        });

        return res
    }
}

let aux = new api()
export default aux