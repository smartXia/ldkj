<script setup>
import { onMounted, reactive, ref } from 'vue'
import { create, list, remove, update } from '../../services/adminApi'

const roleOptions = [
  { label: '超级管理员', value: 'super_admin' },
  { label: '编辑', value: 'editor' },
  { label: '运营', value: 'operator' },
]

const rows = ref([])
const total = ref(0)
const loading = ref(false)
const saving = ref(false)
const error = ref('')
const message = ref('')
const drawerOpen = ref(false)
const editingId = ref(null)
const filters = reactive({
  keyword: '',
  status: '',
  page: 1,
  pageSize: 20,
})
const form = reactive({
  username: '',
  password: '',
  real_name: '',
  email: '',
  phone: '',
  status: 'active',
  roles: [],
})

function resetForm() {
  editingId.value = null
  Object.assign(form, {
    username: '',
    password: '',
    real_name: '',
    email: '',
    phone: '',
    status: 'active',
    roles: [],
  })
}

async function fetchRows() {
  loading.value = true
  error.value = ''
  try {
    const payload = await list('users', filters)
    rows.value = payload.items
    total.value = payload.total
  } catch (err) {
    rows.value = []
    total.value = 0
    error.value = err.message || '用户数据加载失败'
  } finally {
    loading.value = false
  }
}

function openCreate() {
  resetForm()
  drawerOpen.value = true
}

function openEdit(row) {
  editingId.value = row.id
  Object.assign(form, {
    username: row.username || '',
    password: '',
    real_name: row.real_name || '',
    email: row.email || '',
    phone: row.phone || '',
    status: row.status || 'active',
    roles: [...(row.roles || [])],
  })
  drawerOpen.value = true
}

function roleLabel(role) {
  return roleOptions.find((item) => item.value === role)?.label || role
}

async function saveUser() {
  saving.value = true
  error.value = ''
  message.value = ''
  const payload = { ...form, roles: [...form.roles] }
  if (!payload.password) delete payload.password

  try {
    if (editingId.value) {
      await update('users', editingId.value, payload)
      message.value = '用户已更新'
    } else {
      await create('users', payload)
      message.value = '用户已创建'
    }
    drawerOpen.value = false
    await fetchRows()
  } catch (err) {
    error.value = err.message || '保存失败'
  } finally {
    saving.value = false
  }
}

async function deleteUser(row) {
  if (!window.confirm(`确认删除用户 ${row.username}？`)) return
  try {
    await remove('users', row.id)
    await fetchRows()
  } catch (err) {
    error.value = err.message || '删除失败'
  }
}

onMounted(fetchRows)
</script>

<template>
  <section class="admin-panel">
    <header class="admin-panel-header">
      <div>
        <h2>用户管理</h2>
        <span class="admin-muted">共 {{ total }} 个后台账号</span>
      </div>
      <button class="admin-button" type="button" @click="openCreate">新建用户</button>
    </header>

    <div class="admin-toolbar">
      <input v-model.trim="filters.keyword" class="admin-input" style="max-width: 260px" placeholder="搜索账号、姓名或邮箱" />
      <select v-model="filters.status" class="admin-select" style="max-width: 140px">
        <option value="">全部状态</option>
        <option value="active">启用</option>
        <option value="disabled">停用</option>
      </select>
      <button class="admin-button secondary" type="button" @click="fetchRows">查询</button>
    </div>

    <div v-if="error" class="admin-status error">{{ error }}</div>
    <div v-if="message" class="admin-status">{{ message }}</div>
    <div v-if="loading" class="admin-status">正在加载用户</div>
    <div v-else-if="rows.length === 0" class="admin-status">暂无后台用户</div>
    <div v-else class="admin-table-wrap">
      <table class="admin-table">
        <thead>
          <tr>
            <th>账号</th>
            <th>姓名</th>
            <th>联系方式</th>
            <th>角色</th>
            <th>权限</th>
            <th>状态</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in rows" :key="row.id">
            <td>{{ row.username }}</td>
            <td>{{ row.real_name || '-' }}</td>
            <td>{{ row.email || row.phone || '-' }}</td>
            <td>{{ (row.roles || []).map(roleLabel).join('、') || '-' }}</td>
            <td>{{ (row.permissions || []).join('、') || '-' }}</td>
            <td><span class="admin-chip">{{ row.status === 'disabled' ? '停用' : '启用' }}</span></td>
            <td>
              <span class="admin-actions">
                <button class="admin-button secondary" type="button" @click="openEdit(row)">编辑</button>
                <button class="admin-button danger" type="button" @click="deleteUser(row)">删除</button>
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <form v-if="drawerOpen" class="admin-drawer admin-form" @submit.prevent="saveUser">
      <div class="admin-field">
        <label for="username">账号</label>
        <input id="username" v-model.trim="form.username" class="admin-input" required />
      </div>
      <div class="admin-field">
        <label for="password">密码</label>
        <input id="password" v-model="form.password" class="admin-input" type="password" :required="!editingId" placeholder="编辑时留空表示不修改" />
      </div>
      <div class="admin-field">
        <label for="realName">姓名</label>
        <input id="realName" v-model.trim="form.real_name" class="admin-input" />
      </div>
      <div class="admin-field">
        <label for="email">邮箱</label>
        <input id="email" v-model.trim="form.email" class="admin-input" type="email" />
      </div>
      <div class="admin-field">
        <label for="phone">手机</label>
        <input id="phone" v-model.trim="form.phone" class="admin-input" />
      </div>
      <div class="admin-field">
        <label for="status">状态</label>
        <select id="status" v-model="form.status" class="admin-select">
          <option value="active">启用</option>
          <option value="disabled">停用</option>
        </select>
      </div>
      <div class="admin-field wide">
        <label>角色</label>
        <label v-for="role in roleOptions" :key="role.value" class="admin-chip">
          <input v-model="form.roles" type="checkbox" :value="role.value" />
          {{ role.label }}
        </label>
      </div>
      <div class="admin-form-actions">
        <button class="admin-button ghost" type="button" @click="drawerOpen = false">取消</button>
        <button class="admin-button" type="submit" :disabled="saving">{{ saving ? '保存中' : '保存' }}</button>
      </div>
    </form>
  </section>
</template>
