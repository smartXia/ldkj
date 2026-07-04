<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { logout } from '../../services/adminApi'
import '../styles.css'

const route = useRoute()
const router = useRouter()

const navItems = [
  { label: '仪表盘', path: '/admin' },
  { label: '服务', path: '/admin/services' },
  { label: '案例', path: '/admin/cases' },
  { label: '资讯', path: '/admin/news' },
  { label: 'Banner', path: '/admin/banners' },
  { label: '页面', path: '/admin/pages' },
  { label: 'FAQ', path: '/admin/faqs' },
  { label: '表单', path: '/admin/forms' },
  { label: '站点', path: '/admin/site' },
  { label: 'SEO', path: '/admin/seo' },
]

const pageTitle = computed(() => navItems.find((item) => item.path === route.path)?.label || '管理端')

function handleLogout() {
  logout()
  router.push('/admin/login')
}
</script>

<template>
  <div class="admin-shell">
    <aside class="admin-sidebar">
      <div class="admin-brand">
        <strong>灵动信息</strong>
        <span>内容管理</span>
      </div>
      <nav class="admin-nav" aria-label="后台导航">
        <RouterLink
          v-for="item in navItems"
          :key="item.path"
          :to="item.path"
          :class="{ active: route.path === item.path }"
        >
          {{ item.label }}
        </RouterLink>
      </nav>
    </aside>

    <section class="admin-workspace">
      <header class="admin-topbar">
        <div class="admin-topbar-title">
          <p class="admin-topbar-meta">灵动信息后台管理</p>
          <h1>{{ pageTitle }}</h1>
        </div>
        <div class="admin-topbar-actions">
          <RouterLink class="admin-button secondary" to="/">返回前台</RouterLink>
          <button class="admin-button ghost" type="button" @click="handleLogout">退出登录</button>
        </div>
      </header>
      <main class="admin-main">
        <RouterView />
      </main>
    </section>
  </div>
</template>
