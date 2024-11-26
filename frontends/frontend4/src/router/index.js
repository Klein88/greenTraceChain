import { createRouter, createWebHistory } from 'vue-router'
import Login from "@/views/Login.vue";
import Home from "@/views/Home.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
      children: [
        {
          path: '/transactionData',
          name: 'transactionData',
          component: () => import('../views/TransactionData.vue')
        },
        {
          path: '/carbonAccountingData',
          name: 'carbonAccountingData',
          component: () => import('../views/CarbonAccountingData/CarbonAccountingData.vue')
        },
        {
          path: '/traceSource1/:Id',
          name: 'traceSource1',
          component: () => import('../views/CarbonAccountingData/TraceSource1.vue')
        },
        {
          path: '/traceSource2/:Id',
          name: 'traceSource2',
          component: () => import('../views/CarbonAccountingData/TraceSource2.vue')
        },
        {
          path: '/traceSource3/:Id',
          name: 'traceSource3',
          component: () => import('../views/CarbonAccountingData/TraceSource3.vue')
        },
        {
          path: '/traceSource4/:Id',
          name: 'traceSource4',
          component: () => import('../views/CarbonAccountingData/TraceSource4.vue')
        },
        {
          path: '/userView',
          name: 'userView',
          component: () => import('../views/UserView.vue')
        },
        {
          path: '/userRequest',
          name: 'userRequest',
          component: () => import('../views/UserRequest.vue')
        },
        {
          path: '/contractManagement',
          name: 'contractManagement',
          component: () => import('../views/ContractManagement.vue')
        },
        {
          path: '/contractWriting/:url',
          name: 'contractWriting',
          component: () => import('../views/ContractWriting.vue')
        },
        {
          path: '/chainManagement/:url',
          name: 'chainManagement',
          component: () => import('../views/ChainManagement.vue')
        }
      ]
    }
  ]
})

export default router
