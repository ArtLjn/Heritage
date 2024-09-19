import { createRouter, createWebHashHistory } from 'vue-router'
// import Style from '@/views/StyleView.vue'
import Home from '@/views/HomeView.vue'

const routes = [
  {
    meta: {
      title: 'Select style'
    },
    path: '/',
    name: 'style',
    redirect:"/dashboard"
    // component: Style
  },
  {
    // Document title tag
    // We combine it with defaultDocumentTitle set in `src/main.js` on router.afterEach hook
    meta: {
      title: 'Dashboard'
    },
    path: '/dashboard',
    name: 'dashboard',
    component: Home
  },
  {
    meta: {
      title: 'Login'
    },
    path: '/login',
    name: 'login',
    component: () => import('@/views/LoginView.vue')
  },
  {
    meta: {
      title:"Account",
    },
    path:"/account",
    name:"account",
    component: () => import('@/views/superAdmin/manageAccount.vue')
  },
  {
    meta: {
      title:"ApplyManage",
    },
    path:"/applyManage",
    name:"applyManage",
    component: () => import('@/views/admin/applyManage.vue')
  },
  {
    meta: {
      title: "WatchInheritorTask"
    },
    path: "/watchInheritor",
    name:"watchInheritor",
    component: () => import("@/views/admin/seeHeritageInheritorTask.vue")
  },
  {
    meta: {
      title: "WatchProjectTask"
    },
    path: "/watchProject",
    name:"watchProject",
    component: () => import("@/views/admin/seeHeritageProjectTask.vue")
  },
  {
    meta:{
      title: "ShowHeritage"
    },
    path:"/showHeritage",
    name: "showHeritage",
    component: () => import("@/views/show/showHeritage.vue")
  },
  {
    meta:{
      title: "seeInheritor"
    },
    path:"/seeInheritor",
    name: "seeInheritor",
    component: () => import("@/views/show/seeInheritor.vue")
  },
  {
    meta:{
      title: "seeProject"
    },
    path:"/seeProject",
    name: "seeProject",
    component: () => import("@/views/show/seeProject.vue")
  },
  {
    meta:{
      title: "Apply"
    },
    path: "/apply",
    name:"apply",
    component: () => import("@/views/user/apply.vue")
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { top: 0 }
  }
})

export default router
