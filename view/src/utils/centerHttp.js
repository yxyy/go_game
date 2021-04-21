import axios from 'axios'

axios.defaults.headers.common['Authorization'] = window.localStorage.getItem("access_token")

const request = axios.create({
    baseURL:"http://127.0.0.1:8090"
})


export default  request


