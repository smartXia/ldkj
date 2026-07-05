<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { logout } from '../../services/adminApi'
import '../styles.css'

const route = useRoute()
const router = useRouter()

const navItems = [
  { label: '仪表盘', path: '/admin', icon: 'dashboard' },
  { label: '服务', path: '/admin/services', icon: 'services' },
  { label: '案例', path: '/admin/cases', icon: 'cases' },
  { label: '资讯', path: '/admin/news', icon: 'news' },
  { label: 'Banner', path: '/admin/banners', icon: 'banner' },
  { label: '页面', path: '/admin/pages', icon: 'pages' },
  { label: '关于我们', path: '/admin/about', icon: 'pages' },
  { label: 'FAQ', path: '/admin/faqs', icon: 'faq' },
  { label: '表单', path: '/admin/forms', icon: 'forms' },
  { label: '站点', path: '/admin/site', icon: 'site' },
  { label: 'SEO', path: '/admin/seo', icon: 'seo' },
]

const iconPaths = {
  dashboard: 'M3 13h8V3H3v10Zm0 8h8v-6H3v6Zm10 0h8V11h-8v10Zm0-18v6h8V3h-8Z',
  services: 'M4 7h16M4 12h16M4 17h16M7 4v16M17 4v16',
  cases: 'M4 7h16v12H4V7Zm4 0V5h8v2',
  news: 'M5 4h11l3 3v13H5V4Zm10 0v4h4M8 11h8M8 15h8',
  banner: 'M4 5h16v14H4V5Zm3 10 3-4 3 3 2-2 2 3',
  pages: 'M7 3h10l4 4v14H7V3Zm9 1v4h4M3 7h2v14h12v2H3V7Zm7 6h7M10 17h5',
  faq: 'M12 3a8 8 0 1 0 0 16h1v2l3-3a8 8 0 0 0-4-15Zm0 11v-1c2-1 3-2 3-4a3 3 0 0 0-6 0h2a1 1 0 1 1 1 1c-2 1-2 2-2 4h2Zm0 3h.01',
  forms: 'M6 3h12v18H6V3Zm3 5h6M9 12h6M9 16h4',
  site: 'M12 3a9 9 0 1 0 0 18 9 9 0 0 0 0-18Zm0 0c2 2 3 5 3 9s-1 7-3 9M3 12h18',
  seo: 'M10 18a8 8 0 1 1 5.7-2.4L21 21l-2 2-5.2-5.3A8 8 0 0 1 10 18Zm-3-8h6M7 13h4',
}

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
          <svg class="admin-nav-icon" viewBox="0 0 24 24" aria-hidden="true">
            <path :d="iconPaths[item.icon]" />
          </svg>
          <span>{{ item.label }}</span>
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
