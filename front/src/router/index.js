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
      component:()=>import("../views/show/home.vue"),
    },
    {
      path:"/edit",
      name:"edit",
      component:()=>import("../components/show/edit.vue"),
    },
    {
      path:'/look',
      name:"look",
      component:()=>import("../components/show/look.vue")
    },
    {
      path:"/admin",
      name:"admin",
      component:()=>import("../views/admin/LoginView.vue")
    },
    {
      path:"/adminHome",
      name:"adminHome",
      component:()=>import("../views/admin/adminHome.vue"),
      children:[
        {
          path:"index",
          name:"index",
          component:()=>import("../views/admin/oneView.vue")
        },
        {
          path:"apply",
          name:"apply",
          component:()=>import("../views/admin/apply.vue"),
        },
        {
          path:"individual",
          name:"individual",
          component:() => import("../views/admin/individual.vue")
        },
        {
          path:'view',
          name:"view",
          component:()=>import("../views/admin/apply_view.vue")
        },
        {
          path:"heritage",
          name:"heritage",
          component:()=>import("../views/admin/heritage.vue")
        },
        {
          path:"manager",
          name:"manager",
          component:()=>import("../components/admin/ManagerAccount.vue")
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
