<script setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { create, exportFile, list, remove, update } from '../../services/adminApi'
import UploadField from './UploadField.vue'

const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  resource: {
    type: String,
    required: true,
  },
  fields: {
    type: Array,
    required: true,
  },
  columns: {
    type: Array,
    required: true,
  },
  uploadType: {
    type: String,
    default: 'content',
  },
})

const rows = ref([])
const total = ref(0)
const loading = ref(false)
const saving = ref(false)
const error = ref('')
const drawerOpen = ref(false)
const editingId = ref(null)
const filters = reactive({ keyword: '', page: 1, pageSize: 10 })
const form = reactive({})

const isEditing = computed(() => editingId.value !== null)

function resetForm(row = {}) {
  props.fields.forEach((field) => {
    form[field.key] = row[field.key] ?? field.default ?? ''
  })
  editingId.value = row.id ?? null
}

async function fetchRows() {
  loading.value = true
  error.value = ''

  try {
    const payload = await list(props.resource, filters)
    rows.value = payload.items
    total.value = payload.total
  } catch (err) {
    rows.value = []
    total.value = 0
    error.value = err.message || '加载失败'
  } finally {
    loading.value = false
  }
}

function openCreate() {
  resetForm()
  drawerOpen.value = true
}

function openEdit(row) {
  resetForm(row)
  drawerOpen.value = true
}

async function handleSave() {
  saving.value = true
  error.value = ''

  try {
    if (isEditing.value) {
      await update(props.resource, editingId.value, { ...form })
    } else {
      await create(props.resource, { ...form })
    }
    drawerOpen.value = false
    await fetchRows()
  } catch (err) {
    error.value = err.message || '保存失败'
  } finally {
    saving.value = false
  }
}

async function handleDelete(row) {
  if (!window.confirm(`确认删除「${row.title || row.name || row.id}」？`)) return

  try {
    await remove(props.resource, row.id)
    await fetchRows()
  } catch (err) {
    error.value = err.message || '删除失败'
  }
}

function renderCell(row, column) {
  const value = row[column.key]
  if (column.type === 'date' && value) {
    return new Date(value).toLocaleString()
  }
  if (Array.isArray(value)) return value.join(' / ')
  return value || '-'
}

onMounted(fetchRows)
</script>

<template>
  <section class="admin-panel">
    <header class="admin-panel-header">
      <div>
        <h2>{{ title }}</h2>
        <span class="admin-muted">共 {{ total }} 条</span>
      </div>
      <div class="admin-actions">
        <button class="admin-button secondary" type="button" @click="exportFile(resource, filters, 'csv')">导出 CSV</button>
        <button class="admin-button" type="button" @click="openCreate">新建</button>
      </div>
    </header>

    <div class="admin-toolbar">
      <input v-model.trim="filters.keyword" class="admin-input" style="max-width: 260px" placeholder="搜索标题、分类或标签" />
      <button class="admin-button secondary" type="button" @click="fetchRows">查询</button>
    </div>

    <div v-if="error" class="admin-status error">{{ error }}</div>
    <div v-else-if="loading" class="admin-status">正在加载数据</div>
    <div v-else-if="rows.length === 0" class="admin-status">暂无数据，可以先新建内容或等待接口接入</div>
    <div v-else class="admin-table-wrap">
      <table class="admin-table">
        <thead>
          <tr>
            <th v-for="column in columns" :key="column.key">{{ column.label }}</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in rows" :key="row.id">
            <td v-for="column in columns" :key="column.key">
              <span v-if="column.type === 'status'" class="admin-chip">{{ renderCell(row, column) }}</span>
              <span v-else>{{ renderCell(row, column) }}</span>
            </td>
            <td>
              <span class="admin-actions">
                <button class="admin-button secondary" type="button" @click="openEdit(row)">编辑</button>
                <button class="admin-button danger" type="button" @click="handleDelete(row)">删除</button>
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <form v-if="drawerOpen" class="admin-drawer admin-form" @submit.prevent="handleSave">
      <div v-for="field in fields" :key="field.key" class="admin-field" :class="{ wide: field.wide || field.type === 'textarea' || field.type === 'upload' }">
        <label :for="field.key">{{ field.label }}</label>
        <textarea
          v-if="field.type === 'textarea'"
          :id="field.key"
          v-model="form[field.key]"
          class="admin-textarea"
          :placeholder="field.placeholder"
        />
        <select v-else-if="field.type === 'select'" :id="field.key" v-model="form[field.key]" class="admin-select">
          <option v-for="option in field.options" :key="option.value" :value="option.value">{{ option.label }}</option>
        </select>
        <UploadField v-else-if="field.type === 'upload'" v-model="form[field.key]" :type="uploadType" />
        <input
          v-else
          :id="field.key"
          v-model="form[field.key]"
          class="admin-input"
          :type="field.type || 'text'"
          :placeholder="field.placeholder"
        />
      </div>

      <div class="admin-form-actions">
        <button class="admin-button ghost" type="button" @click="drawerOpen = false">取消</button>
        <button class="admin-button" type="submit" :disabled="saving">{{ saving ? '保存中' : '保存' }}</button>
      </div>
    </form>
  </section>
</template>
