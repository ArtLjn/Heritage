import{aj as a,H as o}from"./index-9711d56f.js";const g=()=>a.get("/api/heritage/getCategory"),n=()=>a.get("/api/heritage/getLevel"),p=(e,t,r)=>a.get(`/api/heritage/queryHeritageTask?page=${e}&size=${t}&raw=${r}`),c=e=>a.post("/api/heritage/CreateHeritageInheritor",e),u=e=>a.post("/api/heritage/CreateHeritageProject",e),H=e=>{const t=new FormData;return t.append("file",e),o.post("/api/upload",t,{headers:{"Content-Type":"multipart/form-data"}})},h=(e,t)=>a.get(`/api/heritage/auditHeritageTask?id=${e}&allow=${t}`),l="1",d="2",y=(e,t,r,i)=>a.get(`/api/heritage/queryHeritageByLocate?page=${e}&size=${t}&raw=${r}&locate=${i}`);export{h as A,c as C,g as G,l as H,p as Q,y as a,d as b,n as c,u as d,H as u};
