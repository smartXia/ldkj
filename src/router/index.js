import { createRouter, createWebHistory } from 'vue-router'
import { adminRoutes } from '../admin'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/HomeView.vue'),
  },
  {
    path: '/service',
    name: 'service',
    component: () => import('../views/ServiceView.vue'),
  },
  {
    path: '/service/:id',
    name: 'service-detail',
    component: () => import('../views/ServiceDetailView.vue'),
  },
  {
    path: '/cooperate',
    name: 'cooperate',
    component: () => import('../views/PartnerView.vue'),
  },
  {
    path: '/partner',
    redirect: '/cooperate',
  },
  {
    path: '/case',
    name: 'case',
    component: () => import('../views/CaseView.vue'),
  },
  {
    path: '/case/:id',
    name: 'case-detail',
    component: () => import('../views/CaseDetailView.vue'),
  },
  {
    path: '/message',
    name: 'message',
    component: () => import('../views/MessageView.vue'),
  },
  {
    path: '/message/:id',
    name: 'message-detail',
    component: () => import('../views/MessageDetailView.vue'),
  },
  {
    path: '/about',
    redirect: '/',
  },
  ...adminRoutes,
  {
    path: '/zh-cn',
    redirect: '/',
  },
  {
    path: '/zh-cn/service',
    redirect: '/service',
  },
  {
    path: '/zh-cn/service/:id',
    redirect: (to) => `/service/${to.params.id}`,
  },
  {
    path: '/zh-cn/cooperate',
    redirect: '/cooperate',
  },
  {
    path: '/zh-cn/partner',
    redirect: '/cooperate',
  },
  {
    path: '/zh-cn/case',
    redirect: '/case',
  },
  {
    path: '/zh-cn/case/:id',
    redirect: (to) => `/case/${to.params.id}`,
  },
  {
    path: '/zh-cn/message',
    redirect: '/message',
  },
  {
    path: '/zh-cn/message/:id',
    redirect: (to) => `/message/${to.params.id}`,
  },
  {
    path: '/zh-cn/about',
    redirect: '/',
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

export default createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})
