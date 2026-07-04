<script setup>
import { nextTick, onMounted, ref, watch } from 'vue'

const model = defineModel({ type: String, default: '' })

const editor = ref(null)
const sourceMode = ref(false)

const blockOptions = [
  { label: '正文', value: 'p' },
  { label: '标题 2', value: 'h2' },
  { label: '标题 3', value: 'h3' },
  { label: '引用', value: 'blockquote' },
]

function cleanHtml(html) {
  const template = document.createElement('template')
  template.innerHTML = html || ''
  template.content.querySelectorAll('script').forEach((node) => node.remove())
  return template.innerHTML
}

function syncEditorFromModel() {
  if (!editor.value || sourceMode.value) return
  const nextHtml = cleanHtml(model.value)
  if (editor.value.innerHTML !== nextHtml) editor.value.innerHTML = nextHtml
}

function syncModelFromEditor() {
  if (!editor.value) return
  model.value = cleanHtml(editor.value.innerHTML)
}

function focusEditor() {
  editor.value?.focus()
}

function runCommand(command, value = null) {
  sourceMode.value = false
  nextTick(() => {
    focusEditor()
    document.execCommand(command, false, value)
    syncModelFromEditor()
  })
}

function formatBlock(event) {
  runCommand('formatBlock', event.target.value)
}

function addLink() {
  const url = window.prompt('请输入链接地址')
  if (!url) return
  runCommand('createLink', url)
}

function addImage() {
  const url = window.prompt('请输入图片地址')
  if (!url) return
  runCommand('insertImage', url)
}

function toggleSourceMode() {
  sourceMode.value = !sourceMode.value
  if (!sourceMode.value) nextTick(syncEditorFromModel)
}

function updateSource(event) {
  model.value = cleanHtml(event.target.value)
}

watch(model, syncEditorFromModel)
onMounted(syncEditorFromModel)
</script>

<template>
  <div class="admin-richtext">
    <div class="admin-richtext-toolbar" aria-label="富文本工具栏">
      <select class="admin-richtext-select" aria-label="段落格式" @change="formatBlock">
        <option v-for="option in blockOptions" :key="option.value" :value="option.value">{{ option.label }}</option>
      </select>
      <button type="button" title="加粗" @click="runCommand('bold')"><strong>B</strong></button>
      <button type="button" title="斜体" @click="runCommand('italic')"><em>I</em></button>
      <button type="button" title="下划线" @click="runCommand('underline')"><u>U</u></button>
      <button type="button" title="无序列表" @click="runCommand('insertUnorderedList')">•</button>
      <button type="button" title="有序列表" @click="runCommand('insertOrderedList')">1.</button>
      <button type="button" title="链接" @click="addLink">↗</button>
      <button type="button" title="图片" @click="addImage">▧</button>
      <button type="button" title="清除格式" @click="runCommand('removeFormat')">Tx</button>
      <button type="button" :class="{ active: sourceMode }" title="源码" @click="toggleSourceMode">源码</button>
    </div>

    <textarea
      v-if="sourceMode"
      class="admin-richtext-source"
      :value="model"
      aria-label="HTML 源码"
      @input="updateSource"
    />
    <div
      v-else
      ref="editor"
      class="admin-richtext-editor"
      contenteditable="true"
      role="textbox"
      aria-multiline="true"
      data-placeholder="请输入资讯正文"
      @input="syncModelFromEditor"
      @blur="syncModelFromEditor"
    />
  </div>
</template>
