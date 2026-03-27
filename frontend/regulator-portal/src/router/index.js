import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginOld from "@/views/LoginOld.vue";
import LoginNew from "@/views/LoginNew.vue";
import Home from "@/views/Home.vue";
import Dashboard from "@/views/Dashboard.vue";
import DataAnalysis from "@/views/DataAnalysis.vue";
import DataExamine from "@/views/DataExamine.vue";
import CarbonEmissionReport from "@/views/CarbonEmissionReport.vue";
import CarbonCredit from "@/views/CarbonCredit.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'loginNew',
      component: LoginNew
    },
    {
      path: '/Home',
      name: 'home',
      component: Home,
      meta: {
        title: '首页'
      },
      children: [
        {
          path: '/Dashboard',
          name: 'dashboard',
          component: Dashboard,
          meta: {
            title: '控制台'
          }
        },
        {
          path: '/DataAnalysis',
          name: 'dataAnalysis',
          component: DataAnalysis
          ,
          meta: {
            title: '数据分析'
          }
        },
        {
          path: '/DataExamine',
          name: 'dataExamine',
          component: DataExamine,
          meta: {
            title: '数据审核'
          }
        },
        {
          path: '/CarbonEmissionReport',
          name: 'carbonEmissionReport',
          component: CarbonEmissionReport,
          meta: {
            title: '碳排放报告'
          }
        },
        {
          path: '/CarbonCredit',
          name: 'carbonCredit',
          component: CarbonCredit,
          meta: {
            title: '碳信用'
          }
        }
      ]
    }
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
