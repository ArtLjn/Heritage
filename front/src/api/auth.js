import request from "@/utils/request";
import instance from "@/utils/request";
import router from "@/router";

export const Login = (data) => {
   return request.post('/api/account/login',data)
}

export const LogOut = () => {
   const city = localStorage.getItem('city')
   if (!city) return
   request.get(`/api/account/logout?city=${city}`).then((res) => {
      if (res.data.code === 200) {
         localStorage.clear()
         router.push('/admin')
      }
   }).catch((err) => {
      localStorage.clear()
      router.push('/admin')
   })
}

export const VerifyToken = () => {
   const token = localStorage.getItem('token')
   if (!token) return
   instance.interceptors.request.use(config => {
      config.headers.Authorization = token
      return config
   })
    request.get('/api/account/verifyToken').then((res) => {
       if (res.data.code === 200) {
          router.push('/adminHome')
       } else {
          router.push('/admin')
          localStorage.clear()
       }
    }).catch((err) => {
       router.push('/admin')
       localStorage.clear()
    })
}

export const GetAllAccount = () => {
    return request.get('/api/account/getAllAccount')
}
