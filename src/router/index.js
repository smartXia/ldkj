import { createRouter, createWebHistory } from 'vue-router'
import { adminRoutes } from '../admin'
import { getToken, hasPermission } from '../services/adminApi'

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

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

router.beforeEach((to) => {
  if (!to.path.startsWith('/admin') || to.path === '/admin/login') {
    return true
  }
  if (!getToken()) {
    return '/admin/login'
  }
  const permission = to.matched.find((record) => record.meta?.permission)?.meta.permission
  if (permission && !hasPermission(permission)) {
    return '/admin'
  }
  return true
})

export default router
