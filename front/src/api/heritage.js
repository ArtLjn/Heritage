import request from "@/utils/request";

export const GetCateGory = () => {
    return request.get('/api/heritage/getCategory')
}