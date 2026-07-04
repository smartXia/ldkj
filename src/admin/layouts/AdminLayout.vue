<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { logout } from '../../services/adminApi'
import '../styles.css'

const route = useRoute()
const router = useRouter()

const navItems = [
  { label: '仪表盘', path: '/admin' },
  { label: '案例', path: '/admin/cases' },
  { label: '资讯', path: '/admin/news' },
  { label: 'Banner', path: '/admin/banners' },
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
        <strong>灵动 CMS</strong>
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
        <div>
          <p>南京灵动官网</p>
          <h1>{{ pageTitle }}</h1>
        </div>
        <button class="admin-button ghost" type="button" @click="handleLogout">退出</button>
      </header>
      <main class="admin-main">
        <RouterView />
      </main>
    </section>
  </div>
</template>
