<script setup>
import { onMounted, ref, shallowRef } from 'vue'
import { listResource, updateFormStatus } from '../../services/adminApi'

const rows = ref([])
const error = shallowRef('')

async function loadRows() {
  try {
    const data = await listResource('forms', { page: 1, page_size: 50 })
    rows.value = data.items || []
  } catch (err) {
    error.value = err.message
  }
}

async function markProcessed(row) {
  await updateFormStatus(row.id, 'processed')
  await loadRows()
}

onMounted(loadRows)
</script>

<template>
  <section class="admin-page">
    <div class="page-title">
      <h1>线索管理</h1>
      <button type="button" @click="loadRows">刷新</button>
    </div>
    <p v-if="error" class="admin-error">{{ error }}</p>
    <section class="admin-card">
      <table>
        <thead>
          <tr><th>姓名</th><th>电话</th><th>公司</th><th>需求</th><th>状态</th><th>操作</th></tr>
        </thead>
        <tbody>
          <tr v-for="row in rows" :key="row.id">
            <td>{{ row.name || '-' }}</td>
            <td>{{ row.phone || '-' }}</td>
            <td>{{ row.company || '-' }}</td>
            <td>{{ row.requirement || row.interest || '-' }}</td>
            <td>{{ row.status }}</td>
            <td><button type="button" @click="markProcessed(row)">标记已处理</button></td>
          </tr>
          <tr v-if="!rows.length"><td colspan="6">暂无线索</td></tr>
        </tbody>
      </table>
    </section>
  </section>
</template>
