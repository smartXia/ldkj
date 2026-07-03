import { createRouter, createWebHistory } from 'vue-router'

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
