<script setup>
import { useRouter } from 'vue-router'
import { logoutAdmin } from '../../services/adminApi'

const router = useRouter()

const navItems = [
  { label: '仪表盘', to: '/admin' },
  { label: '案例管理', to: '/admin/cases' },
  { label: '资讯管理', to: '/admin/news' },
  { label: 'Banner', to: '/admin/banners' },
  { label: '服务管理', to: '/admin/services' },
  { label: '线索管理', to: '/admin/leads' },
  { label: '媒体上传', to: '/admin/media' },
  { label: '系统设置', to: '/admin/settings' },
]

function handleLogout() {
  logoutAdmin()
  router.replace('/admin/login')
}
</script>

<template>
  <div class="admin-shell">
    <aside class="admin-sidebar">
      <RouterLink class="admin-brand" to="/admin">
        <strong>WSD</strong>
        <span>内容管理后台</span>
      </RouterLink>
      <nav class="admin-nav" aria-label="后台导航">
        <RouterLink v-for="item in navItems" :key="item.to" :to="item.to">
          {{ item.label }}
        </RouterLink>
      </nav>
    </aside>
    <section class="admin-main">
      <header class="admin-topbar">
        <div>
          <p>Admin Console</p>
          <strong>官网内容运营中心</strong>
        </div>
        <div class="topbar-actions">
          <RouterLink to="/" target="_blank">查看前台</RouterLink>
          <button type="button" @click="handleLogout">退出</button>
        </div>
      </header>
      <main class="admin-content">
        <RouterView />
      </main>
    </section>
  </div>
</template>

<style scoped>
.admin-shell {
  min-height: 100vh;
  display: grid;
  grid-template-columns: 236px minmax(0, 1fr);
  background: #f4f6f9;
  color: #18202f;
}

.admin-sidebar {
  position: sticky;
  top: 0;
  height: 100vh;
  padding: 22px 16px;
  background: #151922;
  color: #fff;
}

.admin-brand {
  display: grid;
  gap: 2px;
  padding: 10px 12px 22px;
}

.admin-brand strong {
  font-size: 24px;
  letter-spacing: 0;
}

.admin-brand span {
  color: #aab2c3;
  font-size: 13px;
}

.admin-nav {
  display: grid;
  gap: 6px;
}

.admin-nav a {
  border-radius: 8px;
  padding: 11px 12px;
  color: #cbd3df;
  font-size: 14px;
}

.admin-nav a.router-link-exact-active,
.admin-nav a:hover {
  background: #ff3d45;
  color: #fff;
}

.admin-main {
  min-width: 0;
}

.admin-topbar {
  height: 72px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 28px;
  border-bottom: 1px solid #e7ebf2;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(14px);
}

.admin-topbar p {
  margin: 0;
  color: #7d8797;
  font-size: 12px;
  text-transform: uppercase;
}

.admin-topbar strong {
  display: block;
  margin-top: 2px;
  font-size: 18px;
}

.topbar-actions {
  display: flex;
  gap: 10px;
}

.topbar-actions a,
.topbar-actions button {
  border: 1px solid #dbe1ea;
  border-radius: 8px;
  padding: 8px 12px;
  background: #fff;
  color: #313a4a;
  font-size: 14px;
}

.topbar-actions button {
  border-color: #ffccd0;
  color: #e93039;
}

.admin-content {
  padding: 28px;
  overflow: visible;
}

@media (max-width: 860px) {
  .admin-shell {
    grid-template-columns: 1fr;
  }

  .admin-sidebar {
    position: static;
    height: auto;
  }

  .admin-nav {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}
</style>
