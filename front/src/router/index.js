import { createRouter, createWebHistory } from 'vue-router'
const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path:"/",
      redirect:"/home"
    },
    {
      path:"/home",
      name:"home",
      component:()=>import("../views/main.vue")
    },
    {
      path:"/admin",
      name:"admin",
      component:()=>import("../views/LoginView.vue")
    },
    {
      path:"/adminHome",
      name:"adminHome",
      component:()=>import("../views/home.vue"),
      redirect:"/adminHome/index",
      children:[
        {
          path:"/adminHome/index",
          name:"index",
          component:()=>import("../views/oneView.vue")
        },
        {
          path:"/adminHome/browse",
          name:"browse",
          component:()=>import("../views/Browse.vue")
        },
        {
          path: "/adminHome/order",
          name: "order",
          component: () => import("../components/template_table.vue")
        }
      ]
    }
  ]
})

export default router
