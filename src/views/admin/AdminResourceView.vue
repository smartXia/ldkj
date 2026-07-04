<script setup>
import { computed, onMounted, reactive, ref, shallowRef, watch } from 'vue'
import { createResource, deleteResource, listResource, updateFormStatus, updateResource } from '../../services/adminApi'

const configs = {
  cases: {
    title: '案例管理',
    resource: 'cases',
    columns: ['title', 'industry', 'platform', 'status'],
    fields: [
      ['title', '标题'], ['slug', '标识'], ['industry', '行业'], ['platform', '平台'],
      ['strategy', '策略'], ['cover_url', '封面图URL'], ['summary', '摘要'], ['content', '正文', 'textarea'], ['status', '状态'],
    ],
    defaults: { status: 'published' },
  },
  news: {
    title: '资讯管理',
    resource: 'news',
    columns: ['title', 'category', 'status', 'published_at'],
    fields: [
      ['title', '标题'], ['slug', '标识'], ['category', '分类'], ['cover_url', '封面图URL'],
      ['summary', '摘要'], ['content', '正文', 'textarea'], ['status', '状态'], ['published_at', '发布时间'],
    ],
    defaults: { status: 'published' },
  },
  banners: {
    title: 'Banner 管理',
    resource: 'banners',
    columns: ['title', 'page', 'status', 'updated_at'],
    fields: [
      ['title', '标题'], ['subtitle', '副标题'], ['image_url', '图片URL'], ['link_url', '链接'],
      ['button_text', '按钮文案'], ['page', '页面'], ['status', '状态'],
    ],
    defaults: { status: 'published', page: 'home' },
  },
  services: {
    title: '服务管理',
    resource: 'services',
    columns: ['title', 'slug', 'status', 'sort_order'],
    fields: [
      ['title', '标题'], ['slug', '标识'], ['subtitle', '副标题'], ['summary', '摘要'],
      ['cover_url', '封面图URL'], ['icon_url', '图标URL'], ['content', '正文', 'textarea'],
      ['highlight_text', '亮点'], ['process_text', '流程'], ['sort_order', '排序'], ['status', '状态'],
    ],
    defaults: { status: 'published', sort_order: 0 },
  },
}

const props = defineProps({
  type: {
    type: String,
    required: true,
  },
})

const config = computed(() => configs[props.type])
const rows = ref([])
const total = shallowRef(0)
const loading = shallowRef(false)
const saving = shallowRef(false)
const error = shallowRef('')
const editing = shallowRef(null)
const form = reactive({})
const keyword = shallowRef('')

function resetForm(row = null) {
  editing.value = row
  const source = row || config.value.defaults || {}
  for (const field of config.value.fields) {
    form[field[0]] = source[field[0]] ?? ''
  }
}

async function loadRows() {
  loading.value = true
  error.value = ''
  try {
    const data = await listResource(config.value.resource, { page: 1, page_size: 30, keyword: keyword.value })
    rows.value = data.items || []
    total.value = data.total || rows.value.length
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

async function save() {
  saving.value = true
  error.value = ''
  try {
    if (editing.value?.id) {
      await updateResource(config.value.resource, editing.value.id, { ...editing.value, ...form })
    } else {
      await createResource(config.value.resource, { ...form })
    }
    resetForm()
    await loadRows()
  } catch (err) {
    error.value = err.message
  } finally {
    saving.value = false
  }
}

async function remove(row) {
  if (!window.confirm(`确认删除「${row.title || row.name || row.id}」？`)) return
  await deleteResource(config.value.resource, row.id)
  await loadRows()
}

watch(() => props.type, () => {
  resetForm()
  loadRows()
})

onMounted(() => {
  resetForm()
  loadRows()
})
</script>

<template>
  <section class="admin-page">
    <div class="page-title">
      <div>
        <h1>{{ config.title }}</h1>
        <p>共 {{ total }} 条，数据来自后端 {{ `/api/admin/${config.resource}` }}</p>
      </div>
      <button type="button" @click="resetForm()">新建</button>
    </div>
    <p v-if="error" class="admin-error">{{ error }}</p>
    <section class="admin-grid">
      <form class="admin-card editor" @submit.prevent="save">
        <h2>{{ editing ? '编辑内容' : '新建内容' }}</h2>
        <label v-for="field in config.fields" :key="field[0]">
          {{ field[1] }}
          <textarea v-if="field[2] === 'textarea'" v-model="form[field[0]]" rows="5" />
          <input v-else v-model="form[field[0]]" />
        </label>
        <button type="submit" :disabled="saving">{{ saving ? '保存中...' : '保存' }}</button>
      </form>
      <section class="admin-card list-card">
        <div class="toolbar">
          <input v-model.trim="keyword" placeholder="搜索关键词" @keyup.enter="loadRows" />
          <button type="button" @click="loadRows">查询</button>
        </div>
        <table>
          <thead>
            <tr>
              <th v-for="col in config.columns" :key="col">{{ col }}</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="row in rows" :key="row.id">
              <td v-for="col in config.columns" :key="col">{{ row[col] || '-' }}</td>
              <td class="actions">
                <button type="button" @click="resetForm(row)">编辑</button>
                <button type="button" @click="remove(row)">删除</button>
              </td>
            </tr>
            <tr v-if="!rows.length">
              <td :colspan="config.columns.length + 1">{{ loading ? '加载中...' : '暂无数据' }}</td>
            </tr>
          </tbody>
        </table>
      </section>
    </section>
  </section>
</template>
