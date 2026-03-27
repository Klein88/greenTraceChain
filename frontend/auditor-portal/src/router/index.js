import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from "@/views/Login.vue";
import Home from "@/views/Home.vue";
import Dashboard from "@/views/Dashboard.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // {
    //   path: '/',
    //   name: 'register',
    //   component:()=>import('../views/LoginNew.vue')
    // },
    {
      path: '/',
      name: 'Regisiter',
      component: ()=>import('../views/Regisiter.vue')
    },
    {
      path: '/Home',
      name: 'home',
      component: Home,
      children: [
        {
          path: '/Dashboard',
          name: 'dashboard',
          component: Dashboard
        },
        {
          path:'/Shop',
          name:'shop',
          component:()=>import('../views/Shop.vue')
        },
        {
          path:'/CarbonAccounting',
          name:'CarbonAccounting',
          component:()=>import('../views/CarbonAccounting.vue')
        },
        {
          path:'/CarbonAccountingData',
          name:'CarbonAccountingData',
          component:()=>import('../views/CarbonAccountingData.vue')
        },
        {
          path:'/carbonAccountingLast',
          name:'carbonAccountingLast',
          component:()=>import('../views/CarbonAccoutingLast.vue')
        },
        {
          path:'/SuccessUpload',
          name:'SuccessUpload',
          component:()=>import('../views/SuccessUpload.vue')
        },
        {
          path:'/Teach',
          name:'Teach',
          component:()=>import('../views/Teach.vue')
        },
        {
          path:'/TeachMain',
          name:'TeachMain',
          component:()=>import('../views/TeachMain.vue')
        },
        {
          path:'/TeachMain2',
          name:'TeachMain2',
          component:()=>import('../views/TeachMain2.vue')
        },
        {
          path:'/Personal',
          name:'Personal',
          component:()=>import('../views/Personal.vue')
        },
        {
          path:'/Traceability',
          name:'Traceability',
          component:()=>import('../views/Traceability.vue')
        },
        {
          path:'/Traceabilitydetail',
          name:'Traceabilitydetail',
          component:()=>import('../views/Traceabilitydetail.vue')
        },
        {
          path:'/CarbonAccoutingLastChemicallndustry',
          name:'CarbonAccoutingLastChemicallndustry',
          component:()=>import('../views/CarbonAccoutingLast_Chemicallndustry.vue')
        },
        {
          path:'/AddNewModel',
          name:'AddNewModel',
          component:()=>import('../views/AddNewModel.vue')
        },
        {
          path:'/NewModel',
          name:'NewModel',
          component:()=>import('../views/NewModel.vue')
        },
      ]
    },
    {
      path: '/TransactionDetail',
      name : 'TransactionDetail',
      component:() => import('../views/TransactionDetailView.vue'),
      props: (route) => ({ dataFromPage1: route.query.data })
    },
/*
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
*/
  ]
})

export default router
