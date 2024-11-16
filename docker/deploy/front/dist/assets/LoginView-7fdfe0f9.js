import{c,o as i,a as p,r as w,n as f,u as V,g as h,b as v,d as S,w as x,v as $,e as m,t as B,_ as P,f as R,V as L,h as C,i as D,j as E,k as r,l as s,m as M,p as N,q as g,s as _,x as b,y as q,z as A,E as I,L as U,A as k}from "./index-3012a831.js";import{_ as y}from "./FormField-5aba7fd6.js";import{_ as T}from "./BaseButtons-a92ea42e.js";const j={__name:"SectionFullScreen",props:{bg:{type:String,required:!0,validator: t=>["purplePink","pinkRed"].includes(t)}},setup(t){const e=t,n=c(()=>{if(V().isEnabled)return h;switch(e.bg){case"purplePink":return S;case"pinkRed":return v}return""});return(u, o)=>(i(),p("div",{class:f(["flex min-h-screen items-center justify-center",n.value])},[w(u.$slots,"default",{cardClass:"w-11/12 md:w-7/12 lg:w-6/12 xl:w-4/12 shadow-2xl"})],2))}},z=["type","name","value"],F=m("span",{class:"check"},null,-1),G={class:"pl-2"},H={__name:"FormCheckRadio",props:{name:{type:String,required:!0},type:{type:String,default:"checkbox",validator: t=>["checkbox","radio","switch"].includes(t)},label:{type:String,default:null},modelValue:{type:[Array,String,Number,Boolean],default:null},inputValue:{type:[String,Number,Boolean],required:!0}},emits:["update:modelValue"],setup(t, {emit:e}){const n=t,u=c({get:()=>n.modelValue,set: a=>{e("update:modelValue",a)}}),o=c(()=>n.type==="radio"?"radio":"checkbox");return(a, d)=>(i(),p("label",{class:f(t.type)},[x(m("input",{"onUpdate:modelValue":d[0]||(d[0]= l=>u.value=l),type:o.value,name:t.name,value:t.inputValue},null,8,z),[[$,u.value]]),F,m("span",G,B(t.label),1)],2))}},J={},K={class:"bg-gray-50 dark:bg-slate-800 dark:text-slate-100"};function O(t, e){return i(),p("div",K,[w(t.$slots,"default")])}const Q=P(J,[["render",O]]),Z={__name:"LoginView",setup(t){R(()=>{L()});const e=C({address:"",password:""}),n=D(),u=()=>{if(e.address===""||e.password===""){I.warning("请填写完整表单");return}U(e).then(o=>{if(o.data.code===200){const a=o.data.data[0];localStorage.setItem("city",a[0]),localStorage.setItem("rank",a[1]),localStorage.setItem("token",a[2]),localStorage.setItem("name",e.address),k.success("欢迎回来🥳"),setTimeout(()=>{n.push("/dashboard")},1e3)}else k.error(o.data.msg)})};return(o, a)=>(i(),E(Q,null,{default:r(()=>[s(j,{bg:"pinkRed"},{default:r(({cardClass:d})=>[s(M,{class:f(d),"is-form":"",onSubmit:a[3]||(a[3]=N(l=>u(),["prevent"]))},{footer:r(()=>[s(T,null,{default:r(()=>[s(g,{type:"submit",color:"info",label:"Login"}),s(g,{to:"/dashboard",color:"info",outline:"",label:"Back"})]),_:1})]),default:r(()=>[s(y,{label:"Login",help:"Please enter your login"},{default:r(()=>[s(_,{modelValue:e.address,"onUpdate:modelValue":a[0]||(a[0]= l=>e.address=l),icon:b(q),name:"login",autocomplete:"username",placeholder:"请输入用户名"},null,8,["modelValue","icon"])]),_:1}),s(y,{label:"Password",help:"Please enter your password"},{default:r(()=>[s(_,{modelValue:e.password,"onUpdate:modelValue":a[1]||(a[1]= l=>e.password=l),icon:b(A),type:"password",name:"password",autocomplete:"current-password",placeholder:"请输入密码"},null,8,["modelValue","icon"])]),_:1}),s(H,{modelValue:e.remember,"onUpdate:modelValue":a[2]||(a[2]= l=>e.remember=l),name:"remember",label:"Remember","input-value":!0},null,8,["modelValue"])]),_:2},1032,["class"])]),_:1})]),_:1}))}};export{Z as default};