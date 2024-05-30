import request from "@/utils/request";
import instance from "@/utils/request";
import axios from "axios";

export const GetCateGory = () => {
    return request.get('/api/heritage/getCategory')
}

export const QueryHeritageInheritor = () => {
    return request.get('/api/heritage/queryHeritageInheritor')
}

export const QueryHeritageProject = () => {
    return request.get('/api/heritage/queryHeritageProject')
}

export const GetLevel = () => {
    return request.get('/api/heritage/getLevel')
}

export const QueryHeritageTask = (page,size,raw) => {
    return request.get(`/api/heritage/queryHeritageTask?page=${page}&size=${size}&raw=${raw}`)
}

export const CreateHeritageInheritor = (data) => {
    return request.post('/api/heritage/CreateHeritageInheritor',data)
}

export const CreateHeritageProject = (data) => {
    return request.post('/api/heritage/CreateHeritageProject',data)
}

//上传文件
export const uploadFile = (file) => {
    const formData = new FormData();
    formData.append("file",file)
    return axios.post('/api/upload',formData,{
        headers:{
            'Content-Type':'multipart/form-data'
        }
    })
}