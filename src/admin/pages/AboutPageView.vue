<script setup>
import { onMounted, reactive, ref } from 'vue'
import { create, list, update } from '../../services/adminApi'
import { useI18n } from '../../composables/useI18n'
import UploadField from '../components/UploadField.vue'

const { messages } = useI18n()
const loading = ref(false)
const saving = ref(false)
const error = ref('')
const success = ref('')
const pageId = ref(null)
const editingRow = ref(null)

const form = reactive({
  hero: { title: '', desc: '', image: '' },
  intro: { title: '', image: '', paragraphs: '' },
  meanings: '',
  valuesTitle: '',
  values: '',
  history: { title: '', video: '', poster: '' },
  honors: '',
  brands: '',
  teamTitle: '',
})

function parseJSON(value, fallback) {
  if (!value) return fallback
  if (typeof value === 'object') return value
  try {
    return JSON.parse(value)
  } catch {
    return fallback
  }
}

function lines(value) {
  return Array.isArray(value) ? value.join('\n') : value || ''
}

function lineList(value) {
  return String(value || '')
    .split('\n')
    .map((item) => item.trim())
    .filter(Boolean)
}

function pretty(value) {
  return JSON.stringify(value || [], null, 2)
}

function fillForm(about) {
  form.hero.title = about.hero?.title || ''
  form.hero.desc = lines(about.hero?.desc)
  form.hero.image = about.hero?.image || ''
  form.intro.title = about.intro?.title || ''
  form.intro.image = about.intro?.image || ''
  form.intro.paragraphs = lines(about.intro?.paragraphs)
  form.meanings = pretty(about.meanings)
  form.valuesTitle = about.valuesTitle || ''
  form.values = pretty(about.values)
  form.history.title = about.history?.title || ''
  form.history.video = about.history?.video || ''
  form.history.poster = about.history?.poster || ''
  form.honors = JSON.stringify(about.honors || { tabs: [], groups: {} }, null, 2)
  form.brands = pretty(about.brands)
  form.teamTitle = about.teamTitle || ''
}

function aboutPayload() {
  return {
    ...messages.value.about,
    hero: {
      title: form.hero.title,
      desc: lineList(form.hero.desc),
      image: form.hero.image,
    },
    intro: {
      title: form.intro.title,
      image: form.intro.image,
      paragraphs: lineList(form.intro.paragraphs),
    },
    meanings: parseJSON(form.meanings, []),
    valuesTitle: form.valuesTitle,
    values: parseJSON(form.values, []),
    history: {
      title: form.history.title,
      video: form.history.video,
      poster: form.history.poster,
    },
    honors: parseJSON(form.honors, { tabs: [], groups: {} }),
    brands: parseJSON(form.brands, []),
    teamTitle: form.teamTitle,
  }
}

async function fetchAboutPage() {
  loading.value = true
  error.value = ''
  success.value = ''

  try {
    const payload = await list('pages', { keyword: 'about', page: 1, pageSize: 20 })
    const page = payload.items.find((item) => item.page_key === 'about' || item.pageKey === 'about' || item.page === 'about')
    pageId.value = page?.id || null
    editingRow.value = page || null
    const extra = parseJSON(page?.extra_data || page?.extraData, {})
    fillForm(extra.about || messages.value.about)
  } catch (err) {
    fillForm(messages.value.about)
    error.value = err.message || '加载关于我们内容失败'
  } finally {
    loading.value = false
  }
}

async function handleSave() {
  saving.value = true
  error.value = ''
  success.value = ''

  try {
    const payload = {
      ...(editingRow.value || {}),
      page_key: 'about',
      pageKey: 'about',
      page: 'about',
      title: editingRow.value?.title || '关于我们',
      content: editingRow.value?.content || '',
      status: 'published',
      extra_data: JSON.stringify({ about: aboutPayload() }),
    }
    if (pageId.value) {
      await update('pages', pageId.value, payload)
    } else {
      const created = await create('pages', payload)
      pageId.value = created.id
      editingRow.value = created
    }
    success.value = '关于我们内容已保存'
  } catch (err) {
    error.value = err.message || '保存关于我们内容失败'
  } finally {
    saving.value = false
  }
}

onMounted(fetchAboutPage)
</script>

<template>
  <section class="admin-panel about-editor">
    <header class="admin-panel-header">
      <div>
        <h2>关于我们</h2>
        <span class="admin-muted">编辑前台 /about 页面展示内容</span>
      </div>
      <div class="admin-actions">
        <button class="admin-button secondary" type="button" @click="fetchAboutPage">刷新</button>
        <button class="admin-button" type="button" :disabled="saving" @click="handleSave">
          {{ saving ? '保存中' : '保存' }}
        </button>
      </div>
    </header>

    <div v-if="error" class="admin-status error">{{ error }}</div>
    <div v-else-if="success" class="admin-status">{{ success }}</div>
    <div v-if="loading" class="admin-status">正在加载内容</div>

    <form class="admin-form about-editor-form" @submit.prevent="handleSave">
      <fieldset class="about-section">
        <legend>首屏 Banner</legend>
        <label>标题<input v-model="form.hero.title" class="admin-input" /></label>
        <label class="wide">描述<textarea v-model="form.hero.desc" class="admin-textarea" /></label>
        <label class="wide">背景图<UploadField v-model="form.hero.image" type="page" /></label>
      </fieldset>

      <fieldset class="about-section">
        <legend>集团介绍</legend>
        <label>标题<input v-model="form.intro.title" class="admin-input" /></label>
        <label class="wide">段落<textarea v-model="form.intro.paragraphs" class="admin-textarea" /></label>
        <label class="wide">配图<UploadField v-model="form.intro.image" type="page" /></label>
      </fieldset>

      <fieldset class="about-section">
        <legend>价值观与释义</legend>
        <label>价值观标题<input v-model="form.valuesTitle" class="admin-input" /></label>
        <label class="wide">WSD 释义 JSON<textarea v-model="form.meanings" class="admin-textarea code-area" /></label>
        <label class="wide">价值观 JSON<textarea v-model="form.values" class="admin-textarea code-area" /></label>
      </fieldset>

      <fieldset class="about-section">
        <legend>发展历程</legend>
        <label>标题<input v-model="form.history.title" class="admin-input" /></label>
        <label class="wide">视频地址<UploadField v-model="form.history.video" type="page" /></label>
        <label class="wide">视频封面<UploadField v-model="form.history.poster" type="page" /></label>
      </fieldset>

      <fieldset class="about-section">
        <legend>荣誉、品牌与团队</legend>
        <label>团队标题<input v-model="form.teamTitle" class="admin-input" /></label>
        <label class="wide">企业荣誉 JSON<textarea v-model="form.honors" class="admin-textarea code-area" /></label>
        <label class="wide">子品牌 JSON<textarea v-model="form.brands" class="admin-textarea code-area" /></label>
      </fieldset>
    </form>
  </section>
</template>

<style scoped>
.about-editor-form {
  display: grid;
  gap: 18px;
}

.about-section {
  margin: 0;
  padding: 18px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  display: grid;
  grid-template-columns: minmax(220px, 1fr) minmax(220px, 1fr);
  gap: 16px;
}

.about-section legend {
  padding: 0 6px;
  color: #111827;
  font-weight: 700;
}

.about-section label {
  display: grid;
  gap: 8px;
  color: #4b5563;
  font-size: 13px;
}

.about-section .wide {
  grid-column: 1 / -1;
}

.code-area {
  min-height: 180px;
  font-family: Consolas, 'Courier New', monospace;
}

@media (max-width: 900px) {
  .about-section {
    grid-template-columns: 1fr;
  }
}
</style>
