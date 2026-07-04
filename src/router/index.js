import { createRouter, createWebHistory } from 'vue-router'
import { isAdminAuthenticated } from '../services/adminApi'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/HomeView.vue'),
  },
  {
    path: '/case',
    name: 'case',
    component: () => import('../views/CaseView.vue'),
  },
  {
    path: '/message',
    name: 'message',
    component: () => import('../views/MessageView.vue'),
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue'),
  },
  {
    path: '/zh-cn',
    redirect: '/',
  },
  {
    path: '/zh-cn/case',
    redirect: '/case',
  },
  {
    path: '/zh-cn/message',
    redirect: '/message',
  },
  {
    path: '/zh-cn/about',
    redirect: '/about',
  },
  {
    path: '/admin/login',
    name: 'admin-login',
    component: () => import('../views/admin/AdminLoginView.vue'),
  },
  {
    path: '/admin',
    component: () => import('../components/admin/AdminShell.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'admin-dashboard',
        component: () => import('../views/admin/AdminDashboardView.vue'),
      },
      {
        path: 'cases',
        name: 'admin-cases',
        component: () => import('../views/admin/AdminResourceView.vue'),
        props: { type: 'cases' },
      },
      {
        path: 'news',
        name: 'admin-news',
        component: () => import('../views/admin/AdminResourceView.vue'),
        props: { type: 'news' },
      },
      {
        path: 'banners',
        name: 'admin-banners',
        component: () => import('../views/admin/AdminResourceView.vue'),
        props: { type: 'banners' },
      },
      {
        path: 'services',
        name: 'admin-services',
        component: () => import('../views/admin/AdminResourceView.vue'),
        props: { type: 'services' },
      },
      {
        path: 'leads',
        name: 'admin-leads',
        component: () => import('../views/admin/AdminLeadsView.vue'),
      },
      {
        path: 'media',
        name: 'admin-media',
        component: () => import('../views/admin/AdminMediaView.vue'),
      },
      {
        path: 'settings',
        name: 'admin-settings',
        component: () => import('../views/admin/AdminSettingsView.vue'),
      },
    ],
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
  if (to.meta.requiresAuth && !isAdminAuthenticated()) {
    return { path: '/admin/login', query: { redirect: to.fullPath } }
  }
  if (to.path === '/admin/login' && isAdminAuthenticated()) {
    return '/admin'
  }
})

export default router
