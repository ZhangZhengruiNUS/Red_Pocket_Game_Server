import axios from 'axios'

export default(url='http://aa1f1d3f1df22476cbf3887045f4b00a-595466472.ap-southeast-1.elb.amazonaws.com') =>{
    return axios.create({
        baseURL:url,
    })
}
