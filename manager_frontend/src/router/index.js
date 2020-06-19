import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    redirect: "/login"
  },
  {
    path: "/admin",
    component: () => import("@/views/admin/index"),
    redirect: "/admin/management",
    children: [
      {
        path: "management",
        component: () => import("@/views/admin/management")
      },
      {
        path: "newArticle",
        component: () => import("@/views/admin/newArticle")
      },
      {
        path: "modifyArticle",
        component: () => import("@/views/admin/modifyArticle")
      },
      {
        path: "imageManagement",
        component: () => import("@/views/admin/imageManagement")
      },
      {
        path:"categoryMangement",
        component:()=>import("@/views/admin/categoryMangement")
      }
    ]
  },
  {
    path: "/register",
    name: "register",
    component: () => import("@/views/register/register")
  },
  {
    path: "/login",
    component: () => import("@/views/login/login")
  },
  {
    path: "*",
    component: () => import("@/views/page404")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
