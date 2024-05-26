import request from "@/utils/request";

export const Login = (data) => {
   return request.post('/api/account/login',data)
}

