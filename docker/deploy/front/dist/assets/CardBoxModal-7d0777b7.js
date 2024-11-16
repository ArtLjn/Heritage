import{o as c,a as k,e as h,t as p,r as C,c as b,w as m,a9 as f,j as i,k as a,l as s,m as y,q as r,T as _,x as v,a6 as x,p as g,aa as B}from "./index-3012a831.js";import{_ as S}from "./BaseButtons-a92ea42e.js";const $={class:"flex items-center justify-between mb-3"},V={class:"text-2xl"},D={__name:"CardBoxComponentTitle",props:{title:{type:String,required:!0}},setup(e){return(l, u)=>(c(),k("div",$,[h("h1",V,p(e.title),1),C(l.$slots,"default")]))}},N={class:"space-y-3"},j={__name:"CardBoxModal",props:{title:{type:String,required:!0},button:{type:String,default:"info"},buttonLabel:{type:String,default:"Done"},hasCancel:Boolean,modelValue:{type:[String,Number,Boolean],default:null}},emits:["update:modelValue","cancel","confirm"],setup(e, {emit:l}){const u=e,o=b({get:()=>u.modelValue,set: t=>l("update:modelValue",t)}),d= t=>{o.value=!1,l(t)},w=()=>d("confirm"),n=()=>d("cancel");return window.addEventListener("keydown", t=>{t.key==="Escape"&&o.value&&n()}),(t, q)=>m((c(),i(B,{onOverlayClick:n},{default:a(()=>[m(s(y,{class:"shadow-lg max-h-modal w-11/12 md:w-3/5 lg:w-2/5 xl:w-4/12 z-50","is-modal":""},{footer:a(()=>[s(S,null,{default:a(()=>[s(r,{label:e.buttonLabel,color:e.button,onClick:w},null,8,["label","color"]),e.hasCancel?(c(),i(r,{key:0,label:"Cancel",color:e.button,outline:"",onClick:n},null,8,["color"])):_("",!0)]),_:1})]),default:a(()=>[s(D,{title:e.title},{default:a(()=>[e.hasCancel?(c(),i(r,{key:0,icon:v(x),color:"whiteDark",small:"","rounded-full":"",onClick:g(n,["prevent"])},null,8,["icon","onClick"])):_("",!0)]),_:1},8,["title"]),h("div",N,[C(t.$slots,"default")])]),_:3},512),[[f,o.value]])]),_:3},512)),[[f,o.value]])}};export{j as _};
