import axios from 'axios'

export default(url='http://localhost:9090') =>{
    return axios.create({
        baseURL:url,
    })
}
