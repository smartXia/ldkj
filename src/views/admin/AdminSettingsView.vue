<script setup>
import { onMounted, reactive, ref, shallowRef } from 'vue'
import { adminRequest } from '../../services/adminApi'

const seoPages = ['home', 'case', 'message', 'about']
const site = reactive({ site_title: '', logo_url: '', contact: '', copyright: '', footer_links: '' })
const seoItems = ref([])
const error = shallowRef('')
const saved = shallowRef('')
const loading = shallowRef(false)

async function requestWithFallback(primary, fallback, options) {
  try {
    return await adminRequest(primary, options)
  } catch (err) {
    if (err.status === 404 && fallback) {
      return adminRequest(fallback, options)
    }
    throw err
  }
}

function normalizeSeoItems(payload) {
  if (Array.isArray(payload?.items)) return payload.items
  if (payload?.page_key || payload?.page) return [payload]
  return []
}

async function loadSeoSettings() {
  try {
    const payload = await adminRequest('/api/admin/settings/seo')
    const items = normalizeSeoItems(payload)
    if (items.length) return items
  } catch (err) {
    if (err.status !== 404) throw err
  }

  const items = []
  for (const page of seoPages) {
    try {
      items.push(await adminRequest(`/api/admin/seo?page_key=${page}`))
    } catch (err) {
      if (err.status !== 404) throw err
    }
  }
  return items
}

async function loadSettings() {
  loading.value = true
  error.value = ''
  saved.value = ''
  try {
    Object.assign(site, await requestWithFallback('/api/admin/settings/site', '/api/admin/settings'))
    seoItems.value = await loadSeoSettings()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

async function saveSite() {
  saved.value = ''
  error.value = ''
  try {
    await requestWithFallback('/api/admin/settings/site', '/api/admin/settings', {
      method: 'PUT',
      body: JSON.stringify(site),
    })
    saved.value = '站点设置已保存'
  } catch (err) {
    error.value = err.message
  }
}

async function saveSeo() {
  saved.value = ''
  error.value = ''
  try {
    try {
      await adminRequest('/api/admin/settings/seo', {
        method: 'PUT',
        body: JSON.stringify({ items: seoItems.value }),
      })
    } catch (err) {
      if (err.status !== 404) throw err
      for (const item of seoItems.value) {
        await adminRequest('/api/admin/seo', {
          method: 'PUT',
          body: JSON.stringify(item),
        })
      }
    }
    saved.value = 'SEO 设置已保存'
  } catch (err) {
    error.value = err.message
  }
}

onMounted(loadSettings)
</script>

<template>
  <section class="admin-page">
    <div class="page-title">
      <h1>系统设置</h1>
      <button type="button" @click="loadSettings">{{ loading ? '加载中...' : '刷新' }}</button>
    </div>
    <p v-if="error" class="admin-error">{{ error }}</p>
    <p v-if="saved" class="admin-success">{{ saved }}</p>
    <section class="admin-grid">
      <form class="admin-card editor" @submit.prevent="saveSite">
        <h2>站点配置</h2>
        <label>站点标题<input v-model="site.site_title" /></label>
        <label>Logo URL<input v-model="site.logo_url" /></label>
        <label>联系方式<textarea v-model="site.contact" rows="4" /></label>
        <label>版权信息<input v-model="site.copyright" /></label>
        <label>页脚链接<textarea v-model="site.footer_links" rows="4" /></label>
        <button type="submit">保存站点配置</button>
      </form>
      <section class="admin-card editor">
        <h2>SEO 配置</h2>
        <div v-for="item in seoItems" :key="item.page_key || item.page || item.id" class="seo-row">
          <strong>{{ item.page_key || item.page }}</strong>
          <label>标题<input v-model="item.title" /></label>
          <label>描述<textarea v-model="item.description" rows="3" /></label>
          <label>关键词<input v-model="item.keywords" /></label>
        </div>
        <p v-if="!seoItems.length">暂无 SEO 配置</p>
        <button type="button" @click="saveSeo">保存 SEO</button>
      </section>
    </section>
  </section>
</template>
