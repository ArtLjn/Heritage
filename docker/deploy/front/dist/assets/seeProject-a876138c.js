import{_ as B,q as p,C as h,P as b,m as y,N as k,O as x,Z as L,$ as C,X as F,I as t,o as g,j as M,k as n,l as o,e,M as s,t as a}from"./index-9711d56f.js";import{_ as S}from"./BaseButtons-a0f83186.js";const v={components:{BaseButtons:S,BaseButton:p,BaseLevel:h,SectionTitleLineWithButton:b,CardBox:y,SectionMain:k,LayoutAuthenticated:x},data(){return{fieldForm:{number:"",name:"",category:"",locate:"",rank:"",details:""}}},mounted(){this.fieldForm=this.$route.query},methods:{mdiKeyboardReturn(){return L},mdiBallotOutline(){return C},back(){F.back()}}},N={class:"space-y-3 text-center md:text-left lg:mx-12"},T={class:"flex justify-center md:block"},j=e("b",null,"Number: ",-1),O=e("b",null,"Name: ",-1),R=e("b",null,"CateGory: ",-1),$=e("b",null,"Rank: ",-1),A=e("b",null,"Locate: ",-1),H=e("b",null,"Details:",-1),K=["innerHTML"];function V(W,q,D,I,l,i){const r=t("BaseButton"),u=t("BaseButtons"),_=t("SectionTitleLineWithButton"),d=t("BaseLevel"),c=t("CardBox"),m=t("SectionMain"),f=t("LayoutAuthenticated");return g(),M(f,null,{default:n(()=>[o(m,null,{default:n(()=>[o(_,{icon:i.mdiBallotOutline(),title:"Heritage Inheritor Message",main:""},{default:n(()=>[o(u,null,{default:n(()=>[o(r,{icon:i.mdiKeyboardReturn(),label:"返回",color:"info","rounded-full":"",small:"",onClick:i.back},null,8,["icon","onClick"])]),_:1})]),_:1},8,["icon"]),o(c,{form:""},{default:n(()=>[o(c,null,{default:n(()=>[o(d,{type:"justify-around lg:justify-center"},{default:n(()=>[e("div",N,[e("div",T,[e("p",null,[j,s(a(l.fieldForm.number),1)]),e("p",null,[O,s(a(l.fieldForm.name),1)]),e("p",null,[R,s(a(l.fieldForm.category),1)]),e("p",null,[$,s(a(l.fieldForm.rank),1)]),e("p",null,[A,s(a(l.fieldForm.locate),1)]),e("p",null,[H,e("span",{innerHTML:l.fieldForm.details},null,8,K)])])])]),_:1})]),_:1})]),_:1})]),_:1})]),_:1})}const G=B(v,[["render",V]]);export{G as default};
