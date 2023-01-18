import axios from 'axios'
import qs from 'qs'
import { ElMessage } from 'element-plus'
// import store from "../store";
// import CryptoJs from 'crypto-js'
// 使用element-ui Message做消息提醒


// var msk=document.getElementsByClassName('axios-mask')[0]
const service = axios.create({

  baseURL: 'http://localhost:8181',
  // 超时时间 单位是ms，这里设置了3s的超时时间
  timeout: 300 * 1000
})
// 2.请求拦截器
service.interceptors.request.use(config => {
  const token=localStorage.token
  if(token){

    config.headers['Authorization']=token
  }
  return config
}, error => {
  Promise.reject(error)
})

// response interceptor
service.interceptors.response.use((config) => {
  return config
}, (error) => {

  if (error.response) {
    const errorMessage = error.response.data === null ? '系统内部异常，请联系网站管理员' : error.response.data.message
    switch (error.response.status) {
      case 404:
        ElMessage('not found')

        break
      case 403:
        ElMessage('暂无操作权限')


        break
      case 401:
        ElMessage('认证已失效，请重新登录')
        // let aid=getQueryVariable('corpId')
        // localStorage.removeItem(aid)
        // alert(aid)

        break
    }
  }
  return Promise.reject(error)
})

export default service
