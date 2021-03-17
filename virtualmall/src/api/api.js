import axios from 'axios'

class api {

    async setStores(info) {
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