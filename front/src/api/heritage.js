import request from "@/utils/request";

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