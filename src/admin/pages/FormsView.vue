<script setup>
import { onMounted, reactive, ref } from 'vue'
import { exportFile, list, update } from '../../services/adminApi'

const rows = ref([])
const total = ref(0)
const loading = ref(false)
const error = ref('')
const filters = reactive({
  keyword: '',
  status: '',
  page: 1,
  pageSize: 20,
})

async function fetchRows() {
  loading.value = true
  error.value = ''

  try {
    const payload = await list('forms', filters)
    rows.value = payload.items
    total.value = payload.total
  } catch (err) {
    rows.value = []
    total.value = 0
    error.value = err.message || '表单数据加载失败'
  } finally {
    loading.value = false
  }
}

async function markStatus(row, status) {
  try {
    await update('forms', row.id, { status })
    await fetchRows()
  } catch (err) {
    error.value = err.message || '状态更新失败'
  }
}

onMounted(fetchRows)
</script>

<template>
  <section class="admin-panel">
    <header class="admin-panel-header">
      <div>
        <h2>表单管理</h2>
        <span class="admin-muted">共 {{ total }} 条</span>
      </div>
      <div class="admin-actions">
        <button class="admin-button secondary" type="button" @click="exportFile('forms', filters, 'csv')">导出 CSV</button>
        <button class="admin-button secondary" type="button" @click="exportFile('forms', filters, 'excel')">导出 Excel</button>
      </div>
    </header>

    <div class="admin-toolbar">
      <input v-model.trim="filters.keyword" class="admin-input" style="max-width: 260px" placeholder="搜索姓名、电话、公司" />
      <select v-model="filters.status" class="admin-select" style="max-width: 140px">
        <option value="">全部状态</option>
        <option value="pending">未处理</option>
        <option value="processed">已处理</option>
      </select>
      <button class="admin-button secondary" type="button" @click="fetchRows">查询</button>
    </div>

    <div v-if="error" class="admin-status error">{{ error }}</div>
    <div v-else-if="loading" class="admin-status">正在加载数据</div>
    <div v-else-if="rows.length === 0" class="admin-status">暂无表单线索</div>
    <div v-else class="admin-table-wrap">
      <table class="admin-table">
        <thead>
          <tr>
            <th>姓名</th>
            <th>电话</th>
            <th>公司</th>
            <th>职位</th>
            <th>邮箱</th>
            <th>来源</th>
            <th>提交时间</th>
            <th>状态</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in rows" :key="row.id">
            <td>{{ row.name || '-' }}</td>
            <td>{{ row.phone || '-' }}</td>
            <td>{{ row.company || '-' }}</td>
            <td>{{ row.position || '-' }}</td>
            <td>{{ row.email || '-' }}</td>
            <td>{{ row.source || '-' }}</td>
            <td>{{ row.createdAt || '-' }}</td>
            <td><span class="admin-chip">{{ row.status === 'processed' ? '已处理' : '未处理' }}</span></td>
            <td>
              <button
                class="admin-button secondary"
                type="button"
                @click="markStatus(row, row.status === 'processed' ? 'pending' : 'processed')"
              >
                {{ row.status === 'processed' ? '标为未处理' : '标为已处理' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
