<script setup>
import { onMounted, reactive, ref } from 'vue'
import { request } from '../../services/adminApi'
import UploadField from '../components/UploadField.vue'

const loading = ref(false)
const saving = ref(false)
const message = ref('')
const error = ref('')
const form = reactive({
  siteTitle: '',
  logo: '',
  phone: '',
  email: '',
  address: '',
  copyright: '',
  icp: '',
  analyticsCode: '',
})

async function fetchSettings() {
  loading.value = true
  error.value = ''

  try {
    const payload = await request('/settings/site')
    Object.assign(form, payload.data || payload)
  } catch (err) {
    error.value = err.message || '站点设置加载失败'
  } finally {
    loading.value = false
  }
}

async function saveSettings() {
  saving.value = true
  message.value = ''
  error.value = ''

  try {
    await request('/settings/site', {
      method: 'PUT',
      body: { ...form },
    })
    message.value = '已保存'
  } catch (err) {
    error.value = err.message || '保存失败'
  } finally {
    saving.value = false
  }
}

onMounted(fetchSettings)
</script>

<template>
  <section class="admin-panel">
    <header class="admin-panel-header">
      <div>
        <h2>站点设置</h2>
        <span class="admin-muted">基础信息、联系方式、版权与统计代码</span>
      </div>
      <button class="admin-button secondary" type="button" @click="fetchSettings">刷新</button>
    </header>

    <div v-if="loading" class="admin-status">正在加载设置</div>
    <div v-if="error" class="admin-status error">{{ error }}</div>
    <div v-if="message" class="admin-status">{{ message }}</div>

    <form class="admin-form" @submit.prevent="saveSettings">
      <div class="admin-field">
        <label for="siteTitle">网站标题</label>
        <input id="siteTitle" v-model="form.siteTitle" class="admin-input" />
      </div>
      <div class="admin-field">
        <label for="phone">联系电话</label>
        <input id="phone" v-model="form.phone" class="admin-input" />
      </div>
      <div class="admin-field wide">
        <label>Logo</label>
        <UploadField v-model="form.logo" type="site" />
      </div>
      <div class="admin-field">
        <label for="email">邮箱</label>
        <input id="email" v-model="form.email" class="admin-input" type="email" />
      </div>
      <div class="admin-field">
        <label for="icp">备案号</label>
        <input id="icp" v-model="form.icp" class="admin-input" />
      </div>
      <div class="admin-field wide">
        <label for="address">地址</label>
        <input id="address" v-model="form.address" class="admin-input" />
      </div>
      <div class="admin-field wide">
        <label for="copyright">版权信息</label>
        <input id="copyright" v-model="form.copyright" class="admin-input" />
      </div>
      <div class="admin-field wide">
        <label for="analyticsCode">统计代码</label>
        <textarea id="analyticsCode" v-model="form.analyticsCode" class="admin-textarea" />
      </div>
      <div class="admin-form-actions">
        <button class="admin-button" type="submit" :disabled="saving">{{ saving ? '保存中' : '保存设置' }}</button>
      </div>
    </form>
  </section>
</template>
