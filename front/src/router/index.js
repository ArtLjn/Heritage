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
          path:"/adminHome/apply",
          name:"apply",
          component:()=>import("../views/apply.vue")
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

router.beforeEach((to, from, next) => {
  // 检查是否访问的是 adminHome 及其子目录
  if (to.path.startsWith("/adminHome")) {
    // 如果用户已经认证，即存在 token
    if (localStorage.getItem("token")) {
      next(); // 允许访问
    } else {
      next("/admin"); // 如果未认证，重定向到登录页面
    }
  } else {
    next(); // 对于其他路由，不做认证检查
  }
});

export default router
