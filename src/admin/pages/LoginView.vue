<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../../services/adminApi'
import '../styles.css'

const router = useRouter()
const form = reactive({
  username: '',
  password: '',
})
const loading = ref(false)
const error = ref('')

async function handleLogin() {
  loading.value = true
  error.value = ''

  try {
    await login(form)
    router.push('/admin')
  } catch (err) {
    error.value = err.message || '登录失败，请检查账号密码'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <main class="admin-login">
    <section class="admin-login-card">
      <h1>管理端登录</h1>
      <p class="admin-muted">灵动信息内容后台</p>
      <form @submit.prevent="handleLogin">
        <div class="admin-field">
          <label for="username">账号</label>
          <input id="username" v-model.trim="form.username" class="admin-input" autocomplete="username" required />
        </div>
        <div class="admin-field">
          <label for="password">密码</label>
          <input id="password" v-model="form.password" class="admin-input" type="password" autocomplete="current-password" required />
        </div>
        <div v-if="error" class="admin-status error">{{ error }}</div>
        <button class="admin-button" type="submit" :disabled="loading">{{ loading ? '登录中' : '登录' }}</button>
      </form>
    </section>
  </main>
</template>
