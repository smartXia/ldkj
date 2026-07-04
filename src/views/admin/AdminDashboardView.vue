<script setup>
import { onMounted, ref, shallowRef } from 'vue'
import { adminRequest } from '../../services/adminApi'

const loading = shallowRef(true)
const error = shallowRef('')
const stats = ref({ cases: 0, news: 0, forms: 0, pendingForms: 0, services: 0 })
const activities = ref([])

async function loadDashboard() {
  loading.value = true
  error.value = ''
  try {
    const data = await adminRequest('/api/admin/dashboard')
    stats.value = data.stats || stats.value
    activities.value = data.activities || []
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(loadDashboard)
</script>

<template>
  <section class="admin-page">
    <div class="page-title">
      <h1>仪表盘</h1>
      <button type="button" @click="loadDashboard">刷新</button>
    </div>
    <p v-if="error" class="admin-error">{{ error }}</p>
    <div class="metric-grid" :class="{ loading }">
      <article>
        <span>待处理线索</span>
        <strong>{{ stats.pendingForms }}</strong>
      </article>
      <article>
        <span>案例数量</span>
        <strong>{{ stats.cases }}</strong>
      </article>
      <article>
        <span>资讯数量</span>
        <strong>{{ stats.news }}</strong>
      </article>
      <article>
        <span>媒体素材</span>
        <strong>{{ stats.services }}</strong>
      </article>
    </div>
    <section class="admin-card">
      <h2>最近操作</h2>
      <table>
        <thead>
          <tr><th>用户</th><th>动作</th><th>模块</th><th>时间</th></tr>
        </thead>
        <tbody>
          <tr v-for="item in activities" :key="item.id">
            <td>{{ item.username || item.user || '-' }}</td>
            <td>{{ item.action }}</td>
            <td>{{ item.module }}</td>
            <td>{{ item.created_at || item.createdAt || '-' }}</td>
          </tr>
          <tr v-if="!activities.length">
            <td colspan="4">暂无操作记录</td>
          </tr>
        </tbody>
      </table>
    </section>
  </section>
</template>
