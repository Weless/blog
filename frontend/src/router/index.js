import Vue from 'vue'
import VueRouter from 'vue-router'
import '../assets/css/pace.css'
import '../assets/css/custom.css'

Vue.use(VueRouter)

  const routes = [
  {
      path: "/",
      redirect: "/home"
  },
  {
    path: '/home',
    name: 'home',
    component:() => import("../views/Home.vue"),
    redirect:"/home/articles",
    children:[
      {
        path: "articles",
        component:()=>import("../views/Articles.vue")
      },
      {
        path:'/categories/:cid',
        name: 'categories',
        component: ()=>import("../views/Categories.vue")
      },
      {
        path:'/archive/:year/:month',
        name: 'archive',
        component:()=>import("../views/Archives.vue")
      },
      {
        path:'/article/:articleId',
        name:'articleDetail',
        component:()=>import("../views/Detail.vue")
      },
    ]
  },
  {
    path:"/404",
    name:"notFound",
    component:()=>import("../views/404.vue")
  },
  {
    path:"*",
    redirect:"/404"
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
