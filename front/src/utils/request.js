import axios from 'axios'
import {ElMessage} from "element-plus";
const baseURL = '/api'

const instance = axios.create({
    baseURL,
    timeout: 10000
})

// 添加请求拦截器
instance.interceptors.request.use(config => {
    // 在发送请求之前做些什么
    config.headers['Content-Type'] = 'application/json'
    const token = localStorage.getItem('token')
    if(token) {
        config.headers.token = token
    }
    return config
}, error => {
    // 对请求错误做些什么
    return Promise.reject(error)
})

// 添加响应拦截器
instance.interceptors.response.use(response => {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    if (response.data.code === 200 || response.data.code === 201) {
        return response
    }
    if(response.data.code === 400) {
        ElMessage.error(response.data.info)
        return new Promise(() => {})
    }
    ElMessage.error(response.data.info)
    return Promise.reject(response.data)
}, error => {
    //特殊情况
    ElMessage.error(error.response?.data?.message || '服务异常')
    return Promise.reject(error) // 返回一个包含错误信息的Promise对象
})

export default instance
export { baseURL }