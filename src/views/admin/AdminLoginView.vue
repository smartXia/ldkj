<script setup>
import { reactive, shallowRef } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { loginAdmin } from '../../services/adminApi'

const router = useRouter()
const route = useRoute()
const form = reactive({ username: 'admin', password: '' })
const loading = shallowRef(false)
const error = shallowRef('')

async function submit() {
  loading.value = true
  error.value = ''
  try {
    await loginAdmin(form)
    router.replace(route.query.redirect || '/admin')
  } catch (err) {
    error.value = err.message || '登录失败'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <main class="login-page">
    <form class="login-panel" @submit.prevent="submit">
      <p>WSD Admin</p>
      <h1>管理后台登录</h1>
      <label>
        账号
        <input v-model.trim="form.username" autocomplete="username" required />
      </label>
      <label>
        密码
        <input v-model="form.password" type="password" autocomplete="current-password" required />
      </label>
      <button type="submit" :disabled="loading">{{ loading ? '登录中...' : '登录' }}</button>
      <span v-if="error" class="error">{{ error }}</span>
    </form>
  </main>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 24px;
  background:
    linear-gradient(135deg, rgba(255, 61, 69, 0.12), transparent 32%),
    #f3f5f8;
}

.login-panel {
  width: min(420px, 100%);
  display: grid;
  gap: 16px;
  padding: 34px;
  border: 1px solid #e4e8f0;
  border-radius: 8px;
  background: #fff;
  box-shadow: 0 20px 60px rgba(20, 28, 44, 0.12);
}

.login-panel p {
  margin: 0;
  color: #ff3d45;
  font-weight: 700;
}

.login-panel h1 {
  margin: 0 0 8px;
  font-size: 28px;
}

.login-panel label {
  display: grid;
  gap: 8px;
  color: #4b5565;
  font-size: 14px;
}

.login-panel input {
  height: 42px;
  border: 1px solid #d7dee8;
  border-radius: 8px;
  padding: 0 12px;
}

.login-panel button {
  height: 44px;
  border: 0;
  border-radius: 8px;
  background: #ff3d45;
  color: #fff;
  font-weight: 700;
}

.login-panel button:disabled {
  opacity: 0.7;
}

.error {
  color: #e93039;
  font-size: 14px;
}
</style>
