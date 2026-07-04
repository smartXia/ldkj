<script setup>
import { onMounted, ref } from 'vue'
import { request } from '../../services/adminApi'

const pages = ref([
  { page: 'home', name: '首页', title: '', description: '', keywords: '' },
  { page: 'case', name: '案例', title: '', description: '', keywords: '' },
  { page: 'message', name: '资讯', title: '', description: '', keywords: '' },
  { page: 'about', name: '关于我们', title: '', description: '', keywords: '' },
])
const loading = ref(false)
const saving = ref(false)
const error = ref('')
const message = ref('')

async function fetchSeo() {
  loading.value = true
  error.value = ''

  try {
    const payload = await request('/settings/seo')
    const rows = payload.items || payload.data || payload
    if (Array.isArray(rows) && rows.length > 0) pages.value = rows
  } catch (err) {
    error.value = err.message || 'SEO 设置加载失败'
  } finally {
    loading.value = false
  }
}

async function saveSeo() {
  saving.value = true
  error.value = ''
  message.value = ''

  try {
    await request('/settings/seo', {
      method: 'PUT',
      body: { items: pages.value },
    })
    message.value = '已保存'
  } catch (err) {
    error.value = err.message || '保存失败'
  } finally {
    saving.value = false
  }
}

onMounted(fetchSeo)
</script>

<template>
  <section class="admin-panel">
    <header class="admin-panel-header">
      <div>
        <h2>SEO 设置</h2>
        <span class="admin-muted">按页面维护 TDK</span>
      </div>
      <button class="admin-button secondary" type="button" @click="fetchSeo">刷新</button>
    </header>

    <div v-if="loading" class="admin-status">正在加载设置</div>
    <div v-if="error" class="admin-status error">{{ error }}</div>
    <div v-if="message" class="admin-status">{{ message }}</div>

    <form @submit.prevent="saveSeo">
      <div class="admin-table-wrap">
        <table class="admin-table">
          <thead>
            <tr>
              <th style="width: 110px">页面</th>
              <th>Title</th>
              <th>Description</th>
              <th>Keywords</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in pages" :key="item.page">
              <td>{{ item.name || item.page }}</td>
              <td><input v-model="item.title" class="admin-input" /></td>
              <td><input v-model="item.description" class="admin-input" /></td>
              <td><input v-model="item.keywords" class="admin-input" /></td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="admin-form-actions" style="padding: 12px 14px">
        <button class="admin-button" type="submit" :disabled="saving">{{ saving ? '保存中' : '保存 SEO' }}</button>
      </div>
    </form>
  </section>
</template>
