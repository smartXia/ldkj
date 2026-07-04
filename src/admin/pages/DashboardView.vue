<script setup>
import { onMounted, ref } from 'vue'
import { request } from '../../services/adminApi'

const loading = ref(false)
const error = ref('')
const stats = ref({
  cases: 0,
  news: 0,
  forms: 0,
  pendingForms: 0,
})
const activities = ref([])

async function fetchDashboard() {
  loading.value = true
  error.value = ''

  try {
    const payload = await request('/dashboard')
    stats.value = payload.stats || payload.data?.stats || stats.value
    activities.value = payload.activities || payload.data?.activities || []
  } catch (err) {
    error.value = err.message || 'Dashboard 数据加载失败'
  } finally {
    loading.value = false
  }
}

onMounted(fetchDashboard)
</script>

<template>
  <div class="admin-grid">
    <div class="admin-grid metrics">
      <article class="admin-metric">
        <span>案例总数</span>
        <strong>{{ stats.cases }}</strong>
      </article>
      <article class="admin-metric">
        <span>资讯总数</span>
        <strong>{{ stats.news }}</strong>
      </article>
      <article class="admin-metric">
        <span>表单线索</span>
        <strong>{{ stats.forms }}</strong>
      </article>
      <article class="admin-metric">
        <span>待处理</span>
        <strong>{{ stats.pendingForms }}</strong>
      </article>
    </div>

    <section class="admin-panel">
      <header class="admin-panel-header">
        <h2>最近操作</h2>
        <button class="admin-button secondary" type="button" @click="fetchDashboard">刷新</button>
      </header>
      <div v-if="error" class="admin-status error">{{ error }}</div>
      <div v-else-if="loading" class="admin-status">正在加载数据</div>
      <div v-else-if="activities.length === 0" class="admin-status">暂无操作记录</div>
      <div v-else class="admin-table-wrap">
        <table class="admin-table">
          <thead>
            <tr>
              <th>时间</th>
              <th>用户</th>
              <th>动作</th>
              <th>对象</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in activities" :key="item.id">
              <td>{{ item.createdAt || item.time || '-' }}</td>
              <td>{{ item.user || '-' }}</td>
              <td>{{ item.action || '-' }}</td>
              <td>{{ item.target || '-' }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </div>
</template>
